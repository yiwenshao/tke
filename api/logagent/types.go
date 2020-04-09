/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package logagent

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AddonPhase defines the phase of helm constructor.
type AddonPhase string

const (
	// AddonPhaseInitializing means is wait initializing.
	AddonPhaseInitializing AddonPhase = "Initializing"
	// AddonPhaseReinitializing means is reinitializing.
	AddonPhaseReinitializing AddonPhase = "Reinitializing"
	// AddonPhaseChecking means is wait checking.
	AddonPhaseChecking AddonPhase = "Checking"
	// AddonPhaseRunning means is running.
	AddonPhaseRunning AddonPhase = "Running"
	// AddonPhaseUpgrading means is upgrading.
	AddonPhaseUpgrading AddonPhase = "Upgrading"
	// AddonPhaseFailed means has been failed.
	AddonPhaseFailed AddonPhase = "Failed"
	// AddonPhasePending means the controller is proceeding deploying
	AddonPhasePending AddonPhase = "Pending"
	// AddonPhaseUnhealthy means some pods of GPUManager is partial running
	AddonPhaseUnhealthy AddonPhase = "Unhealthy"
	// AddonPhaseTerminating means addon terminating
	AddonPhaseTerminating AddonPhase = "Terminating"
	// AddonPhaseUnknown means addon unknown
	AddonPhaseUnknown AddonPhase = "Unknown"
)


// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogAgent is a manager to collect logs of workload.
type LogAgent struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// Spec defines the desired identities of LogCollector.
	// +optional
	Spec LogAgentSpec
	// +optional
	Status LogAgentStatus
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogAgentList is the whole list of all LogCollector which owned by a tenant.
type LogAgentList struct {
	metav1.TypeMeta
	// +optional
	metav1.ListMeta

	// List of volume decorators.
	Items []LogAgent
}


// LogCollectorSpec describes the attributes of a LogCollector.
type LogAgentSpec struct {
	TenantID    string
	ClusterName string
	Version     string
}

// LogCollectorStatus is information about the current status of a LogCollector.
type LogAgentStatus struct {
	// +optional
	Version string
	// Phase is the current lifecycle phase of the LogCollector of cluster.
	// +optional
	Phase AddonPhase
	// Reason is a brief CamelCase string that describes any failure.
	// +optional
	Reason string
	// RetryCount is a int between 0 and 5 that describes the time of retrying initializing.
	// +optional
	RetryCount int32
	// LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.
	// +optional
	LastReInitializingTimestamp metav1.Time
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogAgentProxyOptions is the query options to a kube-apiserver proxy call for LogAgent crd object.
type LogAgentProxyOptions struct {
	metav1.TypeMeta `json:",inline"`

	Namespace string `json:"namespace,omitempty"`
	Name      string `json:"name,omitempty"`
}



// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogFileTree
type LogFileTree struct {
	metav1.TypeMeta `json:",inline"`
	Spec LogFileTreeSpec `json:"spec"`
}

type LogFileTreeSpec struct {
	ClusterId string `json:"clusterId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Container string `json:"container,omitempty"`
	Pod       string `json:"pod,omitempty"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogFileContent
type LogFileContent struct {
	metav1.TypeMeta `json:",inline"`
	Spec            LogFileContentSpec `json:"spec"`
}

type LogFileContentSpec struct {
	ClusterId string `json:"clusterId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Container string `json:"container,omitempty"`
	Pod       string `json:"pod,omitempty"`
	Start     int32  `json:"start,omitempty"`
	Length    int32  `json:"length,omitempty"`
	Filepath  string `json:"filepath,omitempty"`
}



// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogFileContent
type LogFileDownload struct {
	metav1.TypeMeta `json:",inline"`
	Spec            LogFileDownloadSpec `json:"spec"`
}

type LogFileDownloadSpec struct {
	ClusterId string `json:"clusterId,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Container string `json:"container,omitempty"`
	Pod       string `json:"pod,omitempty"`
	Start     int32  `json:"start,omitempty"`
	Length    int32  `json:"length,omitempty"`
	Filepath  string `json:"filepath,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=deleteCollection
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMap holds configuration data for tke to consume.
type ConfigMap struct {
	metav1.TypeMeta
	// +optional
	metav1.ObjectMeta

	// Data contains the configuration data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// Values with non-UTF-8 byte sequences must use the BinaryData field.
	// The keys stored in Data must not overlap with the keys in
	// the BinaryData field, this is enforced during validation process.
	// +optional
	Data map[string]string

	// BinaryData contains the binary data.
	// Each key must consist of alphanumeric characters, '-', '_' or '.'.
	// BinaryData can contain byte sequences that are not in the UTF-8 range.
	// The keys stored in BinaryData must not overlap with the ones in
	// the Data field, this is enforced during validation process.
	// +optional
	BinaryData map[string][]byte
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigMapList is a resource containing a list of ConfigMap objects.
type ConfigMapList struct {
	metav1.TypeMeta

	// +optional
	metav1.ListMeta

	// Items is the list of ConfigMaps.
	Items []ConfigMap
}


//// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
////  LogFileContent
//type LogFileContent struct {
//	metav1.TypeMeta
//
//	ClusterId string
//	Namespace string
//	Name      string
//	Container string
//	Pod       string
//	Start     int32
//	Size      int32
//	Filepath  string
//}
//
//// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
//// LogFileDownload
//type LogFileDownload struct {
//	metav1.TypeMeta
//
//	ClusterId string
//	Namespace string
//	Name      string
//	Container string
//	Pod       string
//	Start     int32
//	Size      int32
//	Filepath  string
//}


//
//// +genclient
//// +genclient:nonNamespaced
//// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
//type LogCollector struct {
//	metav1.TypeMeta
//	metav1.ObjectMeta
//	Spec              LogCollectorSpec
//	Status            LogCollectorStatus
//}
//
//// +genclient:nonNamespaced
//// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
//type LogCollectorList struct {
//	metav1.TypeMeta
//	metav1.ListMeta
//	Items           []LogCollector
//}
//
//// LogCollectorSpec
//type LogCollectorSpec struct {
//	TenantID    string
//	ClusterName string
//	Description string
//}
//
//// TODO: Add useful information, such as conditions, etc.
//type LogCollectorStatus struct {
//	// +optional
//	Input  string
//	// +optional
//	Output string
//}
//
//
//// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
//// APIKeyReq contains expiration time used to apply the api key.
//type APIKeyReq struct {
//	metav1.TypeMeta
//
//	// Expire is required, holds the duration of the api key become invalid. By default, 168h(= seven days)
//	Expire metav1.Duration `json:"expire,omitempty"`
//
//	// Description describes api keys usage.
//	Description string `json:"description"`
//
//	// +optional
//	Spec LogFileTreeSpec `json:"spec"`
//}
