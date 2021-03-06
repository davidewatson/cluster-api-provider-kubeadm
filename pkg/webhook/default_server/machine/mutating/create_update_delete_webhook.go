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
	admissionregistrationv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	clusterv1alpha1 "sigs.k8s.io/cluster-api/pkg/apis/cluster/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission/builder"
)

func init() {
	builderName := "mutating-create-update-delete-machine"
	Builders[builderName] = builder.
		NewWebhookBuilder().
		Name(builderName+".k8s.io").
		Path("/"+builderName).
		Mutating().
		Operations(admissionregistrationv1beta1.Create, admissionregistrationv1beta1.Update, admissionregistrationv1beta1.Delete).
		FailurePolicy(admissionregistrationv1beta1.Fail).
		ForType(&clusterv1alpha1.Machine{})
}
