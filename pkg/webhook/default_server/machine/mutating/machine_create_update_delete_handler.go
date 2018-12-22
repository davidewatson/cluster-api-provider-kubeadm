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

package mutating

import (
	"context"
	"net/http"

	clusterv1alpha1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/runtime/inject"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/types"
)

func init() {
	webhookName := "mutating-create-update-delete-machine"
	if HandlerMap[webhookName] == nil {
		HandlerMap[webhookName] = []admission.Handler{}
	}
	HandlerMap[webhookName] = append(HandlerMap[webhookName], &MachineCreateUpdateDeleteHandler{})
}

// MachineCreateUpdateDeleteHandler handles Machine
type MachineCreateUpdateDeleteHandler struct {
	// Client  client.Client

	// Decoder decodes objects
	Decoder types.Decoder
}

func (h *MachineCreateUpdateDeleteHandler) mutatingMachineFn(ctx context.Context, obj *clusterv1alpha1.Machine) error {
	// TODO(user): implement your admission logic
	return nil
}

var _ admission.Handler = &MachineCreateUpdateDeleteHandler{}

// Handle handles admission requests.
func (h *MachineCreateUpdateDeleteHandler) Handle(ctx context.Context, req types.Request) types.Response {
	obj := &clusterv1alpha1.Machine{}

	err := h.Decoder.Decode(req, obj)
	if err != nil {
		return admission.ErrorResponse(http.StatusBadRequest, err)
	}
	copy := obj.DeepCopy()

	err = h.mutatingMachineFn(ctx, copy)
	if err != nil {
		return admission.ErrorResponse(http.StatusInternalServerError, err)
	}
	return admission.PatchResponse(obj, copy)
}

//var _ inject.Client = &MachineCreateUpdateDeleteHandler{}
//
//// InjectClient injects the client into the MachineCreateUpdateDeleteHandler
//func (h *MachineCreateUpdateDeleteHandler) InjectClient(c client.Client) error {
//	h.Client = c
//	return nil
//}

var _ inject.Decoder = &MachineCreateUpdateDeleteHandler{}

// InjectDecoder injects the decoder into the MachineCreateUpdateDeleteHandler
func (h *MachineCreateUpdateDeleteHandler) InjectDecoder(d types.Decoder) error {
	h.Decoder = d
	return nil
}
