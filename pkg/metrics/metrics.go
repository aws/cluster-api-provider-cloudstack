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

// Package metrics implements custom metrics for CAPC
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	cserrors "sigs.k8s.io/cluster-api-provider-cloudstack/pkg/errors"
	crtlmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

// AcsCustomMetrics encapsulates all CloudStack custom metrics defined for the controller.
type ACSCustomMetrics struct {
	acsReconciliationErrorCount *prometheus.CounterVec
}

// NewCustomMetrics constructs an ACSCustomMetrics with all desired CloudStack custom metrics and any supporting resources.
func NewCustomMetrics() ACSCustomMetrics {
	customMetrics := ACSCustomMetrics{}
	customMetrics.acsReconciliationErrorCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "acs_reconciliation_errors",
			Help: "Count of reconciliation errors caused by ACS issues, bucketed by error code",
		},
		[]string{"acs_error_code"},
	)
	if err := crtlmetrics.Registry.Register(customMetrics.acsReconciliationErrorCount); err != nil {
		if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
			customMetrics.acsReconciliationErrorCount = are.ExistingCollector.(*prometheus.CounterVec)
		} else {
			// Something else went wrong!
			panic(err)
		}
	}

	return customMetrics
}

// EvaluateErrorAndIncrementAcsReconciliationErrorCounter accepts a CloudStack error message and increments
// the custom acs_reconciliation_errors counter, labeled with the error code if present in the error message.
func (m *ACSCustomMetrics) EvaluateErrorAndIncrementAcsReconciliationErrorCounter(acsError error) string {
	errorCode := cserrors.GetACSErrorCode(acsError)
	m.acsReconciliationErrorCount.WithLabelValues(errorCode).Inc()

	return errorCode
}
