/*
Copyright 2018 The Kubernetes Authors.

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

/*
Package webhook provides functions to build and bootstrap an admission webhook server for a k8s cluster.

Build webhooks

	// mgr is the manager that runs the server.
	webhook1, err := NewWebhookBuilder().
		Name("foo.k8s.io").
		Type(WebhookTypeMutating).
		Path("/mutating-pods").
		Operations(admissionregistrationv1beta1.Create).
		ForType(&corev1.Pod{}).
		WithManager(mgr).
		Build(mutatingHandler1, mutatingHandler2)
	if err != nil {
		// handle error
	}

	webhook2, err := NewWebhookBuilder().
		Name("bar.k8s.io").
		Type(WebhookTypeValidating).
		Path("/validating-deployment").
		Operations(admissionregistrationv1beta1.Create, admissionregistrationv1beta1.Update).
		ForType(&appsv1.Deployment{}).
		WithManager(mgr).
		Build(validatingHandler1)
	if err != nil {
		// handle error
	}

Create a server for webhooks.

	as, err := NewServer("baz-admission-server", mrg, ServerOptions{})
	if err != nil {
		// handle error
	}

Register the webhooks in the server.

	err = as.Register(webhook1, webhook2)
	if err != nil {
		// handle error
	}

Start the server by starting the manager

	err := mrg.Start(signals.SetupSignalHandler())
	if err != nil {
		// handle error
	}

*/
package webhook

import (
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
)

var log = logf.KBLog.WithName("webhook")
