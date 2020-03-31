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
var map_APIKeyReq = map[string]string{
	"":            "APIKeyReq contains expiration time used to apply the api key.",
	"expire":      "Expire is required, holds the duration of the api key become invalid. By default, 168h(= seven days)",
	"description": "Description describes api keys usage.",
}

func (APIKeyReq) SwaggerDoc() map[string]string {
	return map_APIKeyReq
}

var map_LogAgent = map[string]string{
	"":     "LogAgent is a manager to collect logs of workload.",
	"spec": "Spec defines the desired identities of LogCollector.",
}

func (LogAgent) SwaggerDoc() map[string]string {
	return map_LogAgent
}

var map_LogAgentList = map[string]string{
	"":      "LogAgentList is the whole list of all LogCollector which owned by a tenant.",
	"items": "List of volume decorators.",
}

func (LogAgentList) SwaggerDoc() map[string]string {
	return map_LogAgentList
}

var map_LogAgentSpec = map[string]string{
	"": "LogCollectorSpec describes the attributes of a LogCollector.",
}

func (LogAgentSpec) SwaggerDoc() map[string]string {
	return map_LogAgentSpec
}

var map_LogAgentStatus = map[string]string{
	"":                            "LogCollectorStatus is information about the current status of a LogCollector.",
	"phase":                       "Phase is the current lifecycle phase of the LogCollector of cluster.",
	"reason":                      "Reason is a brief CamelCase string that describes any failure.",
	"retryCount":                  "RetryCount is a int between 0 and 5 that describes the time of retrying initializing.",
	"lastReInitializingTimestamp": "LastReInitializingTimestamp is a timestamp that describes the last time of retrying initializing.",
}

func (LogAgentStatus) SwaggerDoc() map[string]string {
	return map_LogAgentStatus
}

var map_LogCollectorSpec = map[string]string{
	"": "LogCollectorSpec",
}

func (LogCollectorSpec) SwaggerDoc() map[string]string {
	return map_LogCollectorSpec
}

var map_LogFileTree = map[string]string{
	"": "LogFileTree",
}

func (LogFileTree) SwaggerDoc() map[string]string {
	return map_LogFileTree
}

// AUTO-GENERATED FUNCTIONS END HERE