/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
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

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_ChartGroup = map[string]string{
	"":     "ChartGroup is an chart group.",
	"spec": "Spec defines the desired identities of namespaces in this set.",
}

func (ChartGroup) SwaggerDoc() map[string]string {
	return map_ChartGroup
}

var map_ChartGroupList = map[string]string{
	"":      "ChartGroupList is the whole list of all chart groups which owned by a tenant.",
	"items": "List of namespaces",
}

func (ChartGroupList) SwaggerDoc() map[string]string {
	return map_ChartGroupList
}

var map_ChartGroupSpec = map[string]string{
	"":           "ChartGroupSpec represents an chart group.",
	"finalizers": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
}

func (ChartGroupSpec) SwaggerDoc() map[string]string {
	return map_ChartGroupSpec
}

var map_ChartGroupStatus = map[string]string{
	"":                   "ChartGroupStatus represents information about the status of an chart group.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (ChartGroupStatus) SwaggerDoc() map[string]string {
	return map_ChartGroupStatus
}

var map_ConfigMap = map[string]string{
	"":           "ConfigMap holds configuration data for tke to consume.",
	"data":       "Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.",
	"binaryData": "BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process.",
}

func (ConfigMap) SwaggerDoc() map[string]string {
	return map_ConfigMap
}

var map_ConfigMapList = map[string]string{
	"":      "ConfigMapList is a resource containing a list of ConfigMap objects.",
	"items": "Items is the list of ConfigMaps.",
}

func (ConfigMapList) SwaggerDoc() map[string]string {
	return map_ConfigMapList
}

var map_HardQuantity = map[string]string{
	"": "HardQuantity is a straightforward wrapper of ResourceList.",
}

func (HardQuantity) SwaggerDoc() map[string]string {
	return map_HardQuantity
}

var map_ImageNamespace = map[string]string{
	"":     "ImageNamespace is an image namespace.",
	"spec": "Spec defines the desired identities of namespaces in this set.",
}

func (ImageNamespace) SwaggerDoc() map[string]string {
	return map_ImageNamespace
}

var map_ImageNamespaceList = map[string]string{
	"":      "ImageNamespaceList is the whole list of all image namespaces which owned by a tenant.",
	"items": "List of namespaces",
}

func (ImageNamespaceList) SwaggerDoc() map[string]string {
	return map_ImageNamespaceList
}

var map_ImageNamespaceSpec = map[string]string{
	"":           "ImageNamespaceSpec represents an image namespace.",
	"finalizers": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
}

func (ImageNamespaceSpec) SwaggerDoc() map[string]string {
	return map_ImageNamespaceSpec
}

var map_ImageNamespaceStatus = map[string]string{
	"":                   "ImageNamespaceStatus represents information about the status of an image namespace.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (ImageNamespaceStatus) SwaggerDoc() map[string]string {
	return map_ImageNamespaceStatus
}

var map_Namespace = map[string]string{
	"":     "Namespace is a namespace in cluster.",
	"spec": "Spec defines the desired identities of namespaces in this set.",
}

func (Namespace) SwaggerDoc() map[string]string {
	return map_Namespace
}

var map_NamespaceCert = map[string]string{
	"": "NamespaceCert represents a x509 certificate of a namespace in project.",
}

func (NamespaceCert) SwaggerDoc() map[string]string {
	return map_NamespaceCert
}

var map_NamespaceCertOptions = map[string]string{
	"":          "NamespaceCertOptions is query options of getting namespace with a x509 certificate.",
	"validDays": "Pay attention to const CertOptionValiddays!",
}

func (NamespaceCertOptions) SwaggerDoc() map[string]string {
	return map_NamespaceCertOptions
}

var map_NamespaceList = map[string]string{
	"":      "NamespaceList is the whole list of all namespaces which owned by a tenant.",
	"items": "List of namespaces",
}

func (NamespaceList) SwaggerDoc() map[string]string {
	return map_NamespaceList
}

var map_NamespaceSpec = map[string]string{
	"":           "NamespaceSpec represents a namespace in cluster of a project.",
	"finalizers": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
	"hard":       "Hard represents the total resources of a namespace.",
}

func (NamespaceSpec) SwaggerDoc() map[string]string {
	return map_NamespaceSpec
}

var map_NamespaceStatus = map[string]string{
	"":                   "NamespaceStatus represents information about the status of a namespace in project.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
	"used":               "Used represents the resources of a namespace that are used.",
}

func (NamespaceStatus) SwaggerDoc() map[string]string {
	return map_NamespaceStatus
}

var map_NsEmigration = map[string]string{
	"":     "NsEmigration is an namespace emigration.",
	"spec": "Spec defines the desired identities of emigrations in this set.",
}

func (NsEmigration) SwaggerDoc() map[string]string {
	return map_NsEmigration
}

var map_NsEmigrationList = map[string]string{
	"":      "NsEmigrationList is the whole list of all namespace emigrations which owned by a tenant.",
	"items": "List of namespace emigrations",
}

func (NsEmigrationList) SwaggerDoc() map[string]string {
	return map_NsEmigrationList
}

var map_NsEmigrationSpec = map[string]string{
	"": "NsEmigrationSpec represents a namespace emigration.",
}

func (NsEmigrationSpec) SwaggerDoc() map[string]string {
	return map_NsEmigrationSpec
}

var map_NsEmigrationStatus = map[string]string{
	"":                   "NsEmigrationStatus represents information about the status of a namespace emigration.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (NsEmigrationStatus) SwaggerDoc() map[string]string {
	return map_NsEmigrationStatus
}

var map_Platform = map[string]string{
	"":     "Platform is a platform in TKE.",
	"spec": "Spec defines the desired identities of platforms in this set.",
}

func (Platform) SwaggerDoc() map[string]string {
	return map_Platform
}

var map_PlatformList = map[string]string{
	"":      "PlatformList is the whole list of all platforms which owned by a tenant.",
	"items": "List of platform.",
}

func (PlatformList) SwaggerDoc() map[string]string {
	return map_PlatformList
}

var map_PlatformSpec = map[string]string{
	"": "PlatformSpec is a description of a platform.",
}

func (PlatformSpec) SwaggerDoc() map[string]string {
	return map_PlatformSpec
}

var map_Portal = map[string]string{
	"":              "Portal is a user in TKE.",
	"administrator": "Administrator indicates whether the user is a platform administrator",
	"projects":      "Projects represents the list of projects to which the user belongs, where the key represents project name and the value represents the project display name.",
	"extension":     "Extension is extension info. for projects.",
}

func (Portal) SwaggerDoc() map[string]string {
	return map_Portal
}

var map_PortalProject = map[string]string{
	"":       "PortalProject is a project extension info for portal.",
	"phase":  "Phases of projects.",
	"parent": "Parents of projects.",
}

func (PortalProject) SwaggerDoc() map[string]string {
	return map_PortalProject
}

var map_Project = map[string]string{
	"":     "Project is a project in TKE.",
	"spec": "Spec defines the desired identities of project in this set.",
}

func (Project) SwaggerDoc() map[string]string {
	return map_Project
}

var map_ProjectList = map[string]string{
	"":      "ProjectList is the whole list of all projects which owned by a tenant.",
	"items": "List of projects",
}

func (ProjectList) SwaggerDoc() map[string]string {
	return map_ProjectList
}

var map_ProjectSpec = map[string]string{
	"":                  "ProjectSpec is a description of a project.",
	"finalizers":        "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
	"members":           "Users represents the user list of project.",
	"parentProjectName": "ParentProjectName indicates the superior project name of this service.",
	"clusters":          "Clusters represents clusters that can be used and the resource limits of each cluster.",
}

func (ProjectSpec) SwaggerDoc() map[string]string {
	return map_ProjectSpec
}

var map_ProjectStatus = map[string]string{
	"":                   "ProjectStatus represents information about the status of a project.",
	"clusters":           "Clusters represents clusters that have been used and the resource usage of each cluster.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"reason":             "The reason for the condition's last transition.",
	"message":            "A human readable message indicating details about the transition.",
}

func (ProjectStatus) SwaggerDoc() map[string]string {
	return map_ProjectStatus
}

var map_UsedQuantity = map[string]string{
	"": "UsedQuantity is a straightforward wrapper of ResourceList.",
}

func (UsedQuantity) SwaggerDoc() map[string]string {
	return map_UsedQuantity
}

// AUTO-GENERATED FUNCTIONS END HERE
