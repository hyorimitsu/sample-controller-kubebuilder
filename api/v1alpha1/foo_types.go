/*

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FooSpec defines the desired state of Foo
// validation のマーカーでで、フィールドの Validataion を設定。
// マニフェストの自動生成ではこれを解析し、⽣成している。
// CRD(/config/crd/base/samplecontroller.k8s.io_foo.yaml)参照
type FooSpec struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Format:=string

	DeploymentName string `json:"deploymentName"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=0

	Replicas *int32 `json:"replicas"`
}

// FooStatus defines the observed state of Foo
type FooStatus struct {
	// +optional
	// オプション項目指定
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Foo is the Schema for the foos API
// subresourece:status のマーカーでStatus SubResource（/status を利⽤するためのもの）が有効化された
// CRDマニフェストファイルが生成される
type Foo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FooSpec   `json:"spec,omitempty"`
	Status FooStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FooList contains a list of Foo
type FooList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Foo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Foo{}, &FooList{})
}
