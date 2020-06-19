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

package storage

import (
	"context"
	"sort"

	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
	platforminternalclient "tkestack.io/tke/api/client/clientset/internalversion/typed/platform/internalversion"
	"tkestack.io/tke/pkg/platform/proxy"
	"tkestack.io/tke/pkg/platform/util"
)

// EventREST implements the REST endpoint for find events by a daemonset.
type EventREST struct {
	rest.Storage

	platformClient platforminternalclient.PlatformInterface
}

var _ rest.Getter = &EventREST{}
var _ rest.GroupVersionKindProvider = &EventREST{}

// GroupVersionKind is used to specify a particular GroupVersionKind to discovery.
func (r *EventREST) GroupVersionKind(containingGV schema.GroupVersion) schema.GroupVersionKind {
	return corev1.SchemeGroupVersion.WithKind("EventList")
}

// New returns an empty object that can be used with Create and Update after
// request data has been put into it.
func (r *EventREST) New() runtime.Object {
	return &corev1.EventList{}
}

// Get retrieves the object from the storage. It is required to support Patch.
func (r *EventREST) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
	client, err := proxy.ClientSet(ctx, r.platformClient)
	if err != nil {
		return nil, err
	}

	namespaceName, ok := request.NamespaceFrom(ctx)
	if !ok {
		return nil, errors.NewBadRequest("a namespace must be specified")
	}

	service, err := client.CoreV1().Services(namespaceName).Get(ctx, name, *options)
	if err != nil {
		return nil, errors.NewNotFound(extensionsv1beta1.Resource("services/events"), name)
	}

	selector := fields.AndSelectors(
		fields.OneTermEqualSelector("involvedObject.uid", string(service.UID)),
		fields.OneTermEqualSelector("involvedObject.name", service.Name),
		fields.OneTermEqualSelector("involvedObject.namespace", service.Namespace),
		fields.OneTermEqualSelector("involvedObject.kind", "Service"))
	listOptions := metav1.ListOptions{
		FieldSelector: selector.String(),
	}
	serviceEvents, err := client.CoreV1().Events(namespaceName).List(ctx, listOptions)
	if err != nil {
		return nil, err
	}

	var events util.EventSlice
	for _, serviceEvent := range serviceEvents.Items {
		events = append(events, serviceEvent)
	}

	if len(service.Spec.Selector) > 0 {
		podSelector := labels.FormatLabels(service.Spec.Selector)
		podListOptions := metav1.ListOptions{LabelSelector: podSelector}
		podListByRS, err := client.CoreV1().Pods(namespaceName).List(ctx, podListOptions)
		if err != nil {
			return nil, err
		}
		for _, pod := range podListByRS.Items {
			podEventsSelector := fields.AndSelectors(
				fields.OneTermEqualSelector("involvedObject.uid", string(pod.UID)),
				fields.OneTermEqualSelector("involvedObject.name", pod.Name),
				fields.OneTermEqualSelector("involvedObject.namespace", pod.Namespace),
				fields.OneTermEqualSelector("involvedObject.kind", "Pod"))
			podEventsListOptions := metav1.ListOptions{
				FieldSelector: podEventsSelector.String(),
			}
			podEvents, err := client.CoreV1().Events(namespaceName).List(ctx, podEventsListOptions)
			if err != nil {
				return nil, err
			}

			for _, podEvent := range podEvents.Items {
				events = append(events, podEvent)
			}
		}

		sort.Sort(events)
	}

	return &corev1.EventList{
		Items: events,
	}, nil
}
