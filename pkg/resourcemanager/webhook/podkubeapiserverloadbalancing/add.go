// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package podkubeapiserverloadbalancing

import (
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

const (
	// HandlerName is the name of this webhook handler.
	HandlerName = "pod-kube-apiserver-load-balancing"
	// WebhookPath is the path at which the handler should be registered.
	WebhookPath = "/webhooks/pod-kube-apiserver-load-balancing"
)

// AddToManager adds Handler to the given manager.
func (h *Handler) AddToManager(mgr manager.Manager) error {
	webhook := admission.
		WithCustomDefaulter(mgr.GetScheme(), &corev1.Pod{}, h).
		WithRecoverPanic(true)

	mgr.GetWebhookServer().Register(WebhookPath, webhook)
	return nil
}
