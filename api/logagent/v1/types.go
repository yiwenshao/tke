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

package v1

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
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the desired identities of LogCollector.
	// +optional
	Spec LogAgentSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	// +optional
	Status LogAgentStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogAgentList is the whole list of all LogCollector which owned by a tenant.
type LogAgentList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// List of volume decorators.
	Items []LogAgent `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// LogCollectorSpec describes the attributes of a LogCollector.
type LogAgentSpec struct {
	TenantID    string `json:"tenantID" protobuf:"bytes,1,opt,name=tenantID"`
	ClusterName string `json:"clusterName" protobuf:"bytes,2,opt,name=clusterName"`
	Version     string `json:"version,omitempty" protobuf:"bytes,3,opt,name=version"`
}

// LogCollectorStatus is information about the current status of a LogCollector.
type LogAgentStatus struct {
	// +optional
	Version string `json:"version,omitempty" protobuf:"bytes,1,opt,name=version"`
	// Phase is the current lifecycle phase of the LogCollector of cluster.
	// +optional
	Phase AddonPhase `json:"phase,omitempty" protobuf:"bytes,2,opt,name=phase"`
	// Reason is a brief CamelCase string that describes any failure.
	// +optional
	Reason string `json:"reason,omitempty" protobuf:"bytes,3,opt,name=reason"`
	// RetryCount is a int between 0 and 5 that describes the time of retrying initializing.
	// +optional
	RetryCount int32 `json:"retryCount" protobuf:"varint,4,name=retryCount"`
	// LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.
	// +optional
	LastReInitializingTimestamp metav1.Time `json:"lastReInitializingTimestamp" protobuf:"bytes,5,name=lastReInitializingTimestamp"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LogFileTree
type LogFileTree struct {
	metav1.TypeMeta `json:",inline"`

	ClusterId string `json:"clusterId,omitempty" protobuf:"bytes,1,opt,name=clusterId"`
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
	Name      string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
	Container string `json:"container,omitempty" protobuf:"bytes,4,opt,name=container"`
	Pod       string `json:"pod,omitempty" protobuf:"bytes,5,opt,name=pod"`
}

//// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
////  LogFileContent
//type LogFileContent struct {
//	metav1.TypeMeta `json:",inline"`
//
//	ClusterId string `json:"clusterId,omitempty" protobuf:"bytes,1,opt,name=clusterId"`
//	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
//	Name      string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
//	Container string `json:"container,omitempty" protobuf:"bytes,4,opt,name=container"`
//	Pod       string `json:"pod,omitempty" protobuf:"bytes,5,opt,name=pod"`
//	Start     int32  `json:"start,omitempty" protobuf:"varint,6,opt,name=start"`
//	Length    int32  `json:"length,omitempty" protobuf:"varint,7,opt,name=length"`
//	Filepath  string `json:"filepath,omitempty" protobuf:"bytes,8,opt,name=filepath"`
//}
//
//// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
//
//// LogFileDownload
//type LogFileDownload struct {
//	metav1.TypeMeta `json:",inline"`
//
//	ClusterId string `json:"clusterId,omitempty" protobuf:"bytes,1,opt,name=clusterId"`
//	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`
//	Name      string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
//	Container string `json:"container,omitempty" protobuf:"bytes,4,opt,name=container"`
//	Pod       string `json:"pod,omitempty" protobuf:"bytes,5,opt,name=pod"`
//	Start     int32  `json:"start,omitempty" protobuf:"varint,6,opt,name=start"`
//	Length    int32  `json:"length,omitempty" protobuf:"varint,7,opt,name=length"`
//	Filepath  string `json:"filepath,omitempty" protobuf:"bytes,8,opt,name=filepath"`
//}
