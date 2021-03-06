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

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	notify "tkestack.io/tke/api/notify"
)

// MessagesGetter has a method to return a MessageInterface.
// A group's client should implement this interface.
type MessagesGetter interface {
	Messages() MessageInterface
}

// MessageInterface has methods to work with Message resources.
type MessageInterface interface {
	Create(*notify.Message) (*notify.Message, error)
	Update(*notify.Message) (*notify.Message, error)
	UpdateStatus(*notify.Message) (*notify.Message, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*notify.Message, error)
	List(opts v1.ListOptions) (*notify.MessageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.Message, err error)
	MessageExpansion
}

// messages implements MessageInterface
type messages struct {
	client rest.Interface
}

// newMessages returns a Messages
func newMessages(c *NotifyClient) *messages {
	return &messages{
		client: c.RESTClient(),
	}
}

// Get takes name of the message, and returns the corresponding message object, and an error if there is any.
func (c *messages) Get(name string, options v1.GetOptions) (result *notify.Message, err error) {
	result = &notify.Message{}
	err = c.client.Get().
		Resource("messages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Messages that match those selectors.
func (c *messages) List(opts v1.ListOptions) (result *notify.MessageList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &notify.MessageList{}
	err = c.client.Get().
		Resource("messages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested messages.
func (c *messages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("messages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a message and creates it.  Returns the server's representation of the message, and an error, if there is any.
func (c *messages) Create(message *notify.Message) (result *notify.Message, err error) {
	result = &notify.Message{}
	err = c.client.Post().
		Resource("messages").
		Body(message).
		Do().
		Into(result)
	return
}

// Update takes the representation of a message and updates it. Returns the server's representation of the message, and an error, if there is any.
func (c *messages) Update(message *notify.Message) (result *notify.Message, err error) {
	result = &notify.Message{}
	err = c.client.Put().
		Resource("messages").
		Name(message.Name).
		Body(message).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *messages) UpdateStatus(message *notify.Message) (result *notify.Message, err error) {
	result = &notify.Message{}
	err = c.client.Put().
		Resource("messages").
		Name(message.Name).
		SubResource("status").
		Body(message).
		Do().
		Into(result)
	return
}

// Delete takes name of the message and deletes it. Returns an error if one occurs.
func (c *messages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("messages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched message.
func (c *messages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.Message, err error) {
	result = &notify.Message{}
	err = c.client.Patch(pt).
		Resource("messages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
