/*
Copyright 2022 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"reflect"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"

	infrav1 "github.com/aws/cluster-api-provider-cloudstack/api/v1beta1"
	csCtrlrUtils "github.com/aws/cluster-api-provider-cloudstack/controllers/utils"
	"github.com/aws/cluster-api-provider-cloudstack/pkg/cloud"
	capiv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/util"
)

// RBAC permissions used in all reconcilers. Events and Secrets.
// "" empty string as the api group indicates core kubernetes objects. "*" indicates all objects.
// +kubebuilder:rbac:groups="",resources=secrets;,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=events,verbs=get;list;watch;create;update;patch

// RBAC permissions for CloudStackCluster.
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=cloudstackclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=cloudstackclusters/status,verbs=create;get;update;patch
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=cloudstackclusters/finalizers,verbs=update
// +kubebuilder:rbac:groups=cluster.x-k8s.io,resources=clusters;clusters/status,verbs=get;list;watch

// CloudStackClusterReconciliationRunner is a ReconciliationRunner with extensions specific to CloudStackClusters.
// The runner does the actual reconciliation.
type CloudStackClusterReconciliationRunner struct {
	csCtrlrUtils.ReconciliationRunner
	Zones                 *infrav1.CloudStackZoneList
	ReconciliationSubject *infrav1.CloudStackCluster
	CSUser                cloud.Client
}

// CloudStackClusterReconciler is the k8s controller manager's interface to reconcile a CloudStackCluster.
// This is primarily to adapt to k8s.
type CloudStackClusterReconciler struct {
	csCtrlrUtils.ReconcilerBase
}

// Initialize a new CloudStackCluster reconciliation runner with concrete types and initialized member fields.
func NewCSClusterReconciliationRunner() *CloudStackClusterReconciliationRunner {
	// Set concrete type and init pointers.
	runner := &CloudStackClusterReconciliationRunner{}
	runner.ReconciliationSubject = &infrav1.CloudStackCluster{}
	runner.ReconciliationRunner = csCtrlrUtils.NewRunner(runner.ReconciliationSubject) // Initializes base pointers.
	runner.Zones = &infrav1.CloudStackZoneList{}

	// For the CloudStackCluster, the ReconciliationSubject is the CSCluster
	runner.CSCluster = runner.ReconciliationSubject

	return runner
}

// Reconcile is the method k8s will call upon a reconciliation request.
func (reconciler *CloudStackClusterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (retRes ctrl.Result, retErr error) {
	r := NewCSClusterReconciliationRunner()
	return r.
		UsingBaseReconciler(reconciler.ReconcilerBase).
		ForRequest(req).
		WithRequestCtx(ctx).
		RunReconciliationStages(
			r.SetupLogger,
			r.GetReconciliationSubject,
			r.GetCAPICluster,
			r.SetupPatcher,
			r.IfDeletionTimestampIsZero(r.Reconcile),
			r.Else(r.ReconcileDelete),
			r.PatchChangesBackToAPI,
		)
}

// Reconcile actually reconciles the CloudStackCluster.
func (r *CloudStackClusterReconciliationRunner) Reconcile() (res ctrl.Result, reterr error) {
	return r.RunReconciliationStages(
		r.CreateZones(r.CSCluster.Spec.Zones),
		r.CheckOwnedCRDsForReadiness(infrav1.GroupVersion.WithKind("CloudStackZone")),
		r.GetZones(r.Zones),
		r.VerifyZoneCRDs,
		r.SetFailureDomains,
		r.ResolveClusterDetails)
}

// ResolveClusterDetails fetches cluster specific details like domain and account IDs.
func (r *CloudStackClusterReconciliationRunner) ResolveClusterDetails() (ctrl.Result, error) {
	// Ensure that CAPI won't prematurely delete this CloudStackCluster.
	controllerutil.AddFinalizer(r.CSCluster, infrav1.ClusterFinalizer)

	// Create and or fetch cluster components.
	err := r.CS.GetOrCreateCluster(r.CSCluster)
	if err == nil {
		r.Log.Info("Fetched cluster info successfully.")
		r.Log.V(1).Info("Post fetch cluster status.", "clusterStatus", r.ReconciliationSubject.Status)

		// Set cluster to ready to indicate readiness to CAPI.
		r.ReconciliationSubject.Status.Ready = true
	}
	return ctrl.Result{}, err
}

// CheckZoneDetails verifies the Zone CRDs found match against those requested.
func (r *CloudStackClusterReconciliationRunner) VerifyZoneCRDs() (ctrl.Result, error) {
	expected := len(r.CSCluster.Spec.Zones)
	actual := len(r.Zones.Items)
	if expected != actual {
		return r.RequeueWithMessage(fmt.Sprintf("Expected %d Zones, but found %d", expected, actual))
	}
	return ctrl.Result{}, nil
}

// SetFailureDomains sets failure domains to be used for CAPI machine placement.
func (r *CloudStackClusterReconciliationRunner) SetFailureDomains() (ctrl.Result, error) {
	r.CSCluster.Status.FailureDomains = capiv1.FailureDomains{}
	for _, zone := range r.Zones.Items {
		r.CSCluster.Status.FailureDomains[zone.Spec.ID] = capiv1.FailureDomainSpec{ControlPlane: true}
	}
	return ctrl.Result{}, nil
}

// ReconcileDelete cleans up resources used by the cluster and finaly removes the CloudStackCluster's finalizers.
func (r *CloudStackClusterReconciliationRunner) ReconcileDelete() (ctrl.Result, error) {
	r.Log.V(1).Info("Deleting CloudStackCluster.")
	if err := r.CS.DisposeClusterResources(r.ReconciliationSubject); err != nil {
		return ctrl.Result{}, err
	}
	controllerutil.RemoveFinalizer(r.ReconciliationSubject, infrav1.ClusterFinalizer)
	return ctrl.Result{}, nil
}

// Called in main, this registers the cluster reconciler to the CAPI controller manager.
func (r *CloudStackClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	controller, err := ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.CloudStackCluster{}).
		WithEventFilter(
			predicate.Funcs{
				UpdateFunc: func(e event.UpdateEvent) bool {
					oldCluster := e.ObjectOld.(*infrav1.CloudStackCluster).DeepCopy()
					newCluster := e.ObjectNew.(*infrav1.CloudStackCluster).DeepCopy()
					// Ignore resource version because they are unique
					oldCluster.ObjectMeta.ResourceVersion = ""
					newCluster.ObjectMeta.ResourceVersion = ""
					// Ignore finalizers updates
					oldCluster.ObjectMeta.Finalizers = nil
					newCluster.ObjectMeta.Finalizers = nil
					// Ignore ManagedFields because they are mirror of ObjectMeta
					oldCluster.ManagedFields = nil
					newCluster.ManagedFields = nil
					// Ignore incremental status updates
					oldCluster.Status = infrav1.CloudStackClusterStatus{}
					newCluster.Status = infrav1.CloudStackClusterStatus{}

					return !reflect.DeepEqual(oldCluster, newCluster)
				},
			},
		).Build(r)
	if err != nil {
		return err
	}
	return controller.Watch( // Add a watch on CAPI Cluster objects for unpause and ready events.
		&source.Kind{Type: &capiv1.Cluster{}},
		handler.EnqueueRequestsFromMapFunc(
			util.ClusterToInfrastructureMapFunc(infrav1.GroupVersion.WithKind("CloudStackCluster"))),
		predicate.Funcs{
			UpdateFunc: func(e event.UpdateEvent) bool {
				r.BaseLogger.Info("Reoncile Update Event triggered.")
				oldCluster := e.ObjectOld.(*capiv1.Cluster)
				newCluster := e.ObjectNew.(*capiv1.Cluster)
				return oldCluster.Spec.Paused && !newCluster.Spec.Paused
			},
			DeleteFunc: func(e event.DeleteEvent) bool {
				r.BaseLogger.Info("Reoncile Delete Event triggered.")
				_, ok := e.Object.GetAnnotations()[capiv1.PausedAnnotation]
				return ok
			},
			CreateFunc: func(e event.CreateEvent) bool {
				//r.BaseLogger.V(1).Info("Reoncile Create Event triggered.")
				r.BaseLogger.Info("Reoncile Create Event triggered.")
				_, ok := e.Object.GetAnnotations()[capiv1.PausedAnnotation]
				return ok
			}},
	)
}
