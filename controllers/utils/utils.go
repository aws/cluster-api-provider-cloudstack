/*
Copyright 2022.

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

package utils

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	capiv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	controlplanev1 "sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"
	"sigs.k8s.io/cluster-api/util"
	clientPkg "sigs.k8s.io/controller-runtime/pkg/client"
)

// GetMachineSet attempts to fetch a MachineSet from CAPI machine owner reference.
func GetMachineSetFromCAPIMachine(
	ctx context.Context,
	client clientPkg.Client,
	capiMachine *capiv1.Machine,
) (*capiv1.MachineSet, error) {

	ref := GetManagementOwnerRef(capiMachine)
	gv, err := schema.ParseGroupVersion(ref.APIVersion)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if gv.Group == capiv1.GroupVersion.Group {
		key := clientPkg.ObjectKey{
			Namespace: capiMachine.Namespace,
			Name:      ref.Name,
		}

		machineSet := &capiv1.MachineSet{}
		if err := client.Get(ctx, key, machineSet); err != nil {
			return nil, err
		}

		return machineSet, nil
	}
	return nil, nil
}

// GetKubeadmControlPlaneFromCAPIMachine attempts to fetch a KubeadmControlPlane from a CAPI machine owner reference.
func GetKubeadmControlPlaneFromCAPIMachine(
	ctx context.Context,
	client clientPkg.Client,
	capiMachine *capiv1.Machine,
) (*controlplanev1.KubeadmControlPlane, error) {

	ref := GetManagementOwnerRef(capiMachine)
	gv, err := schema.ParseGroupVersion(ref.APIVersion)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if gv.Group == controlplanev1.GroupVersion.Group {
		key := clientPkg.ObjectKey{
			Namespace: capiMachine.Namespace,
			Name:      ref.Name,
		}

		controlPlane := &controlplanev1.KubeadmControlPlane{}
		if err := client.Get(ctx, key, controlPlane); err != nil {
			return nil, err
		}

		return controlPlane, nil
	}
	return nil, nil
}

// IsOwnerDeleted returns a boolean if the owner of the CAPI machine has been deleted.
func IsOwnerDeleted(ctx context.Context, client clientPkg.Client, capiMachine *capiv1.Machine) (bool, error) {
	if util.IsControlPlaneMachine(capiMachine) {
		// The controlplane sticks around after deletion pending the deletion of its machiens.
		// As such, need to check the deletion timestamp thereof.
		if cp, err := GetKubeadmControlPlaneFromCAPIMachine(ctx, client, capiMachine); cp != nil && cp.DeletionTimestamp == nil {
			return false, nil
		} else if err != nil && !strings.Contains(err.Error(), "not found") {
			return false, err
		}
	} else {
		// The machineset is deleted immediately, regardless of machine ownership.
		// It is sufficient to check for its existence.
		if ms, err := GetMachineSetFromCAPIMachine(ctx, client, capiMachine); ms != nil {
			return false, nil
		} else if err != nil && !strings.Contains(err.Error(), "not found") {
			return false, err
		}
	}
	return true, nil
}

// fetchOwnerRef simply searches a list of OwnerReference objects for a given kind.
func fetchOwnerRef(refList []meta.OwnerReference, kind string) *meta.OwnerReference {
	for _, ref := range refList {
		if ref.Kind == kind {
			return &ref
		}
	}
	return nil
}

// GetManagementOwnerRef returns the reference object pointing to the CAPI machine's manager.
func GetManagementOwnerRef(capiMachine *capiv1.Machine) *meta.OwnerReference {
	if util.IsControlPlaneMachine(capiMachine) {
		return fetchOwnerRef(capiMachine.OwnerReferences, "KubeadmControlPlane")
	} else {
		return fetchOwnerRef(capiMachine.OwnerReferences, "MachineSet")
	}
}
