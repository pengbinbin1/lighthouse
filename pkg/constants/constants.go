/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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

package constants

import (
	"sigs.k8s.io/mcs-api/pkg/apis/v1alpha1"
)

const (
	LabelSourceNamespace  = "lighthouse.submariner.io/sourceNamespace"
	LabelValueManagedBy   = "lighthouse-agent.submariner.io"
	MCSLabelSourceCluster = "multicluster.kubernetes.io/source-cluster"
	LabelIsHeadless       = "lighthouse.submariner.io/is-headless"
)

const ServiceExportSynced v1alpha1.ServiceExportConditionType = "Synced"
