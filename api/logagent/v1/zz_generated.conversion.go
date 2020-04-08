// +build !ignore_autogenerated

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	logagent "tkestack.io/tke/api/logagent"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ConfigMap)(nil), (*logagent.ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMap_To_logagent_ConfigMap(a.(*ConfigMap), b.(*logagent.ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.ConfigMap)(nil), (*ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_ConfigMap_To_v1_ConfigMap(a.(*logagent.ConfigMap), b.(*ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMapList)(nil), (*logagent.ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMapList_To_logagent_ConfigMapList(a.(*ConfigMapList), b.(*logagent.ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.ConfigMapList)(nil), (*ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_ConfigMapList_To_v1_ConfigMapList(a.(*logagent.ConfigMapList), b.(*ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogAgent)(nil), (*logagent.LogAgent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogAgent_To_logagent_LogAgent(a.(*LogAgent), b.(*logagent.LogAgent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogAgent)(nil), (*LogAgent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogAgent_To_v1_LogAgent(a.(*logagent.LogAgent), b.(*LogAgent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogAgentList)(nil), (*logagent.LogAgentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogAgentList_To_logagent_LogAgentList(a.(*LogAgentList), b.(*logagent.LogAgentList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogAgentList)(nil), (*LogAgentList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogAgentList_To_v1_LogAgentList(a.(*logagent.LogAgentList), b.(*LogAgentList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogAgentProxyOptions)(nil), (*logagent.LogAgentProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogAgentProxyOptions_To_logagent_LogAgentProxyOptions(a.(*LogAgentProxyOptions), b.(*logagent.LogAgentProxyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogAgentProxyOptions)(nil), (*LogAgentProxyOptions)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogAgentProxyOptions_To_v1_LogAgentProxyOptions(a.(*logagent.LogAgentProxyOptions), b.(*LogAgentProxyOptions), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogAgentSpec)(nil), (*logagent.LogAgentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogAgentSpec_To_logagent_LogAgentSpec(a.(*LogAgentSpec), b.(*logagent.LogAgentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogAgentSpec)(nil), (*LogAgentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogAgentSpec_To_v1_LogAgentSpec(a.(*logagent.LogAgentSpec), b.(*LogAgentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogAgentStatus)(nil), (*logagent.LogAgentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogAgentStatus_To_logagent_LogAgentStatus(a.(*LogAgentStatus), b.(*logagent.LogAgentStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogAgentStatus)(nil), (*LogAgentStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogAgentStatus_To_v1_LogAgentStatus(a.(*logagent.LogAgentStatus), b.(*LogAgentStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogFileContent)(nil), (*logagent.LogFileContent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogFileContent_To_logagent_LogFileContent(a.(*LogFileContent), b.(*logagent.LogFileContent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogFileContent)(nil), (*LogFileContent)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogFileContent_To_v1_LogFileContent(a.(*logagent.LogFileContent), b.(*LogFileContent), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogFileContentSpec)(nil), (*logagent.LogFileContentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogFileContentSpec_To_logagent_LogFileContentSpec(a.(*LogFileContentSpec), b.(*logagent.LogFileContentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogFileContentSpec)(nil), (*LogFileContentSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogFileContentSpec_To_v1_LogFileContentSpec(a.(*logagent.LogFileContentSpec), b.(*LogFileContentSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogFileDownload)(nil), (*logagent.LogFileDownload)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogFileDownload_To_logagent_LogFileDownload(a.(*LogFileDownload), b.(*logagent.LogFileDownload), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogFileDownload)(nil), (*LogFileDownload)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogFileDownload_To_v1_LogFileDownload(a.(*logagent.LogFileDownload), b.(*LogFileDownload), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogFileDownloadSpec)(nil), (*logagent.LogFileDownloadSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogFileDownloadSpec_To_logagent_LogFileDownloadSpec(a.(*LogFileDownloadSpec), b.(*logagent.LogFileDownloadSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogFileDownloadSpec)(nil), (*LogFileDownloadSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogFileDownloadSpec_To_v1_LogFileDownloadSpec(a.(*logagent.LogFileDownloadSpec), b.(*LogFileDownloadSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogFileTree)(nil), (*logagent.LogFileTree)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogFileTree_To_logagent_LogFileTree(a.(*LogFileTree), b.(*logagent.LogFileTree), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogFileTree)(nil), (*LogFileTree)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogFileTree_To_v1_LogFileTree(a.(*logagent.LogFileTree), b.(*LogFileTree), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LogFileTreeSpec)(nil), (*logagent.LogFileTreeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LogFileTreeSpec_To_logagent_LogFileTreeSpec(a.(*LogFileTreeSpec), b.(*logagent.LogFileTreeSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*logagent.LogFileTreeSpec)(nil), (*LogFileTreeSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_logagent_LogFileTreeSpec_To_v1_LogFileTreeSpec(a.(*logagent.LogFileTreeSpec), b.(*LogFileTreeSpec), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_ConfigMap_To_logagent_ConfigMap(in *ConfigMap, out *logagent.ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_v1_ConfigMap_To_logagent_ConfigMap is an autogenerated conversion function.
func Convert_v1_ConfigMap_To_logagent_ConfigMap(in *ConfigMap, out *logagent.ConfigMap, s conversion.Scope) error {
	return autoConvert_v1_ConfigMap_To_logagent_ConfigMap(in, out, s)
}

func autoConvert_logagent_ConfigMap_To_v1_ConfigMap(in *logagent.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_logagent_ConfigMap_To_v1_ConfigMap is an autogenerated conversion function.
func Convert_logagent_ConfigMap_To_v1_ConfigMap(in *logagent.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	return autoConvert_logagent_ConfigMap_To_v1_ConfigMap(in, out, s)
}

func autoConvert_v1_ConfigMapList_To_logagent_ConfigMapList(in *ConfigMapList, out *logagent.ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]logagent.ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ConfigMapList_To_logagent_ConfigMapList is an autogenerated conversion function.
func Convert_v1_ConfigMapList_To_logagent_ConfigMapList(in *ConfigMapList, out *logagent.ConfigMapList, s conversion.Scope) error {
	return autoConvert_v1_ConfigMapList_To_logagent_ConfigMapList(in, out, s)
}

func autoConvert_logagent_ConfigMapList_To_v1_ConfigMapList(in *logagent.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_logagent_ConfigMapList_To_v1_ConfigMapList is an autogenerated conversion function.
func Convert_logagent_ConfigMapList_To_v1_ConfigMapList(in *logagent.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	return autoConvert_logagent_ConfigMapList_To_v1_ConfigMapList(in, out, s)
}

func autoConvert_v1_LogAgent_To_logagent_LogAgent(in *LogAgent, out *logagent.LogAgent, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_LogAgentSpec_To_logagent_LogAgentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_LogAgentStatus_To_logagent_LogAgentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_LogAgent_To_logagent_LogAgent is an autogenerated conversion function.
func Convert_v1_LogAgent_To_logagent_LogAgent(in *LogAgent, out *logagent.LogAgent, s conversion.Scope) error {
	return autoConvert_v1_LogAgent_To_logagent_LogAgent(in, out, s)
}

func autoConvert_logagent_LogAgent_To_v1_LogAgent(in *logagent.LogAgent, out *LogAgent, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_logagent_LogAgentSpec_To_v1_LogAgentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_logagent_LogAgentStatus_To_v1_LogAgentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_logagent_LogAgent_To_v1_LogAgent is an autogenerated conversion function.
func Convert_logagent_LogAgent_To_v1_LogAgent(in *logagent.LogAgent, out *LogAgent, s conversion.Scope) error {
	return autoConvert_logagent_LogAgent_To_v1_LogAgent(in, out, s)
}

func autoConvert_v1_LogAgentList_To_logagent_LogAgentList(in *LogAgentList, out *logagent.LogAgentList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]logagent.LogAgent)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_LogAgentList_To_logagent_LogAgentList is an autogenerated conversion function.
func Convert_v1_LogAgentList_To_logagent_LogAgentList(in *LogAgentList, out *logagent.LogAgentList, s conversion.Scope) error {
	return autoConvert_v1_LogAgentList_To_logagent_LogAgentList(in, out, s)
}

func autoConvert_logagent_LogAgentList_To_v1_LogAgentList(in *logagent.LogAgentList, out *LogAgentList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]LogAgent)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_logagent_LogAgentList_To_v1_LogAgentList is an autogenerated conversion function.
func Convert_logagent_LogAgentList_To_v1_LogAgentList(in *logagent.LogAgentList, out *LogAgentList, s conversion.Scope) error {
	return autoConvert_logagent_LogAgentList_To_v1_LogAgentList(in, out, s)
}

func autoConvert_v1_LogAgentProxyOptions_To_logagent_LogAgentProxyOptions(in *LogAgentProxyOptions, out *logagent.LogAgentProxyOptions, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1_LogAgentProxyOptions_To_logagent_LogAgentProxyOptions is an autogenerated conversion function.
func Convert_v1_LogAgentProxyOptions_To_logagent_LogAgentProxyOptions(in *LogAgentProxyOptions, out *logagent.LogAgentProxyOptions, s conversion.Scope) error {
	return autoConvert_v1_LogAgentProxyOptions_To_logagent_LogAgentProxyOptions(in, out, s)
}

func autoConvert_logagent_LogAgentProxyOptions_To_v1_LogAgentProxyOptions(in *logagent.LogAgentProxyOptions, out *LogAgentProxyOptions, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_logagent_LogAgentProxyOptions_To_v1_LogAgentProxyOptions is an autogenerated conversion function.
func Convert_logagent_LogAgentProxyOptions_To_v1_LogAgentProxyOptions(in *logagent.LogAgentProxyOptions, out *LogAgentProxyOptions, s conversion.Scope) error {
	return autoConvert_logagent_LogAgentProxyOptions_To_v1_LogAgentProxyOptions(in, out, s)
}

func autoConvert_v1_LogAgentSpec_To_logagent_LogAgentSpec(in *LogAgentSpec, out *logagent.LogAgentSpec, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.ClusterName = in.ClusterName
	out.Version = in.Version
	return nil
}

// Convert_v1_LogAgentSpec_To_logagent_LogAgentSpec is an autogenerated conversion function.
func Convert_v1_LogAgentSpec_To_logagent_LogAgentSpec(in *LogAgentSpec, out *logagent.LogAgentSpec, s conversion.Scope) error {
	return autoConvert_v1_LogAgentSpec_To_logagent_LogAgentSpec(in, out, s)
}

func autoConvert_logagent_LogAgentSpec_To_v1_LogAgentSpec(in *logagent.LogAgentSpec, out *LogAgentSpec, s conversion.Scope) error {
	out.TenantID = in.TenantID
	out.ClusterName = in.ClusterName
	out.Version = in.Version
	return nil
}

// Convert_logagent_LogAgentSpec_To_v1_LogAgentSpec is an autogenerated conversion function.
func Convert_logagent_LogAgentSpec_To_v1_LogAgentSpec(in *logagent.LogAgentSpec, out *LogAgentSpec, s conversion.Scope) error {
	return autoConvert_logagent_LogAgentSpec_To_v1_LogAgentSpec(in, out, s)
}

func autoConvert_v1_LogAgentStatus_To_logagent_LogAgentStatus(in *LogAgentStatus, out *logagent.LogAgentStatus, s conversion.Scope) error {
	out.Version = in.Version
	out.Phase = logagent.AddonPhase(in.Phase)
	out.Reason = in.Reason
	out.RetryCount = in.RetryCount
	out.LastReInitializingTimestamp = in.LastReInitializingTimestamp
	return nil
}

// Convert_v1_LogAgentStatus_To_logagent_LogAgentStatus is an autogenerated conversion function.
func Convert_v1_LogAgentStatus_To_logagent_LogAgentStatus(in *LogAgentStatus, out *logagent.LogAgentStatus, s conversion.Scope) error {
	return autoConvert_v1_LogAgentStatus_To_logagent_LogAgentStatus(in, out, s)
}

func autoConvert_logagent_LogAgentStatus_To_v1_LogAgentStatus(in *logagent.LogAgentStatus, out *LogAgentStatus, s conversion.Scope) error {
	out.Version = in.Version
	out.Phase = AddonPhase(in.Phase)
	out.Reason = in.Reason
	out.RetryCount = in.RetryCount
	out.LastReInitializingTimestamp = in.LastReInitializingTimestamp
	return nil
}

// Convert_logagent_LogAgentStatus_To_v1_LogAgentStatus is an autogenerated conversion function.
func Convert_logagent_LogAgentStatus_To_v1_LogAgentStatus(in *logagent.LogAgentStatus, out *LogAgentStatus, s conversion.Scope) error {
	return autoConvert_logagent_LogAgentStatus_To_v1_LogAgentStatus(in, out, s)
}

func autoConvert_v1_LogFileContent_To_logagent_LogFileContent(in *LogFileContent, out *logagent.LogFileContent, s conversion.Scope) error {
	if err := Convert_v1_LogFileTreeSpec_To_logagent_LogFileTreeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_LogFileContent_To_logagent_LogFileContent is an autogenerated conversion function.
func Convert_v1_LogFileContent_To_logagent_LogFileContent(in *LogFileContent, out *logagent.LogFileContent, s conversion.Scope) error {
	return autoConvert_v1_LogFileContent_To_logagent_LogFileContent(in, out, s)
}

func autoConvert_logagent_LogFileContent_To_v1_LogFileContent(in *logagent.LogFileContent, out *LogFileContent, s conversion.Scope) error {
	if err := Convert_logagent_LogFileTreeSpec_To_v1_LogFileTreeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_logagent_LogFileContent_To_v1_LogFileContent is an autogenerated conversion function.
func Convert_logagent_LogFileContent_To_v1_LogFileContent(in *logagent.LogFileContent, out *LogFileContent, s conversion.Scope) error {
	return autoConvert_logagent_LogFileContent_To_v1_LogFileContent(in, out, s)
}

func autoConvert_v1_LogFileContentSpec_To_logagent_LogFileContentSpec(in *LogFileContentSpec, out *logagent.LogFileContentSpec, s conversion.Scope) error {
	out.ClusterId = in.ClusterId
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Container = in.Container
	out.Pod = in.Pod
	out.Start = in.Start
	out.Length = in.Length
	out.Filepath = in.Filepath
	return nil
}

// Convert_v1_LogFileContentSpec_To_logagent_LogFileContentSpec is an autogenerated conversion function.
func Convert_v1_LogFileContentSpec_To_logagent_LogFileContentSpec(in *LogFileContentSpec, out *logagent.LogFileContentSpec, s conversion.Scope) error {
	return autoConvert_v1_LogFileContentSpec_To_logagent_LogFileContentSpec(in, out, s)
}

func autoConvert_logagent_LogFileContentSpec_To_v1_LogFileContentSpec(in *logagent.LogFileContentSpec, out *LogFileContentSpec, s conversion.Scope) error {
	out.ClusterId = in.ClusterId
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Container = in.Container
	out.Pod = in.Pod
	out.Start = in.Start
	out.Length = in.Length
	out.Filepath = in.Filepath
	return nil
}

// Convert_logagent_LogFileContentSpec_To_v1_LogFileContentSpec is an autogenerated conversion function.
func Convert_logagent_LogFileContentSpec_To_v1_LogFileContentSpec(in *logagent.LogFileContentSpec, out *LogFileContentSpec, s conversion.Scope) error {
	return autoConvert_logagent_LogFileContentSpec_To_v1_LogFileContentSpec(in, out, s)
}

func autoConvert_v1_LogFileDownload_To_logagent_LogFileDownload(in *LogFileDownload, out *logagent.LogFileDownload, s conversion.Scope) error {
	if err := Convert_v1_LogFileTreeSpec_To_logagent_LogFileTreeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_LogFileDownload_To_logagent_LogFileDownload is an autogenerated conversion function.
func Convert_v1_LogFileDownload_To_logagent_LogFileDownload(in *LogFileDownload, out *logagent.LogFileDownload, s conversion.Scope) error {
	return autoConvert_v1_LogFileDownload_To_logagent_LogFileDownload(in, out, s)
}

func autoConvert_logagent_LogFileDownload_To_v1_LogFileDownload(in *logagent.LogFileDownload, out *LogFileDownload, s conversion.Scope) error {
	if err := Convert_logagent_LogFileTreeSpec_To_v1_LogFileTreeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_logagent_LogFileDownload_To_v1_LogFileDownload is an autogenerated conversion function.
func Convert_logagent_LogFileDownload_To_v1_LogFileDownload(in *logagent.LogFileDownload, out *LogFileDownload, s conversion.Scope) error {
	return autoConvert_logagent_LogFileDownload_To_v1_LogFileDownload(in, out, s)
}

func autoConvert_v1_LogFileDownloadSpec_To_logagent_LogFileDownloadSpec(in *LogFileDownloadSpec, out *logagent.LogFileDownloadSpec, s conversion.Scope) error {
	out.ClusterId = in.ClusterId
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Container = in.Container
	out.Pod = in.Pod
	out.Start = in.Start
	out.Length = in.Length
	out.Filepath = in.Filepath
	return nil
}

// Convert_v1_LogFileDownloadSpec_To_logagent_LogFileDownloadSpec is an autogenerated conversion function.
func Convert_v1_LogFileDownloadSpec_To_logagent_LogFileDownloadSpec(in *LogFileDownloadSpec, out *logagent.LogFileDownloadSpec, s conversion.Scope) error {
	return autoConvert_v1_LogFileDownloadSpec_To_logagent_LogFileDownloadSpec(in, out, s)
}

func autoConvert_logagent_LogFileDownloadSpec_To_v1_LogFileDownloadSpec(in *logagent.LogFileDownloadSpec, out *LogFileDownloadSpec, s conversion.Scope) error {
	out.ClusterId = in.ClusterId
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Container = in.Container
	out.Pod = in.Pod
	out.Start = in.Start
	out.Length = in.Length
	out.Filepath = in.Filepath
	return nil
}

// Convert_logagent_LogFileDownloadSpec_To_v1_LogFileDownloadSpec is an autogenerated conversion function.
func Convert_logagent_LogFileDownloadSpec_To_v1_LogFileDownloadSpec(in *logagent.LogFileDownloadSpec, out *LogFileDownloadSpec, s conversion.Scope) error {
	return autoConvert_logagent_LogFileDownloadSpec_To_v1_LogFileDownloadSpec(in, out, s)
}

func autoConvert_v1_LogFileTree_To_logagent_LogFileTree(in *LogFileTree, out *logagent.LogFileTree, s conversion.Scope) error {
	if err := Convert_v1_LogFileTreeSpec_To_logagent_LogFileTreeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_LogFileTree_To_logagent_LogFileTree is an autogenerated conversion function.
func Convert_v1_LogFileTree_To_logagent_LogFileTree(in *LogFileTree, out *logagent.LogFileTree, s conversion.Scope) error {
	return autoConvert_v1_LogFileTree_To_logagent_LogFileTree(in, out, s)
}

func autoConvert_logagent_LogFileTree_To_v1_LogFileTree(in *logagent.LogFileTree, out *LogFileTree, s conversion.Scope) error {
	if err := Convert_logagent_LogFileTreeSpec_To_v1_LogFileTreeSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_logagent_LogFileTree_To_v1_LogFileTree is an autogenerated conversion function.
func Convert_logagent_LogFileTree_To_v1_LogFileTree(in *logagent.LogFileTree, out *LogFileTree, s conversion.Scope) error {
	return autoConvert_logagent_LogFileTree_To_v1_LogFileTree(in, out, s)
}

func autoConvert_v1_LogFileTreeSpec_To_logagent_LogFileTreeSpec(in *LogFileTreeSpec, out *logagent.LogFileTreeSpec, s conversion.Scope) error {
	out.ClusterId = in.ClusterId
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Container = in.Container
	out.Pod = in.Pod
	return nil
}

// Convert_v1_LogFileTreeSpec_To_logagent_LogFileTreeSpec is an autogenerated conversion function.
func Convert_v1_LogFileTreeSpec_To_logagent_LogFileTreeSpec(in *LogFileTreeSpec, out *logagent.LogFileTreeSpec, s conversion.Scope) error {
	return autoConvert_v1_LogFileTreeSpec_To_logagent_LogFileTreeSpec(in, out, s)
}

func autoConvert_logagent_LogFileTreeSpec_To_v1_LogFileTreeSpec(in *logagent.LogFileTreeSpec, out *LogFileTreeSpec, s conversion.Scope) error {
	out.ClusterId = in.ClusterId
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Container = in.Container
	out.Pod = in.Pod
	return nil
}

// Convert_logagent_LogFileTreeSpec_To_v1_LogFileTreeSpec is an autogenerated conversion function.
func Convert_logagent_LogFileTreeSpec_To_v1_LogFileTreeSpec(in *logagent.LogFileTreeSpec, out *LogFileTreeSpec, s conversion.Scope) error {
	return autoConvert_logagent_LogFileTreeSpec_To_v1_LogFileTreeSpec(in, out, s)
}
