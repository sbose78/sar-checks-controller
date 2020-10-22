package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ExampleSpec defines the desired state of Example
type ExampleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Username is the user on behalf of whom SAR checks need to be done.
	Username string `json:"username"`

	// TestSubject is the resource that would be validated for permissions.
	TestSubject `json:"testSubject"`
}

// TestSubject is the resource that would be validated for permissions.
type TestSubject struct {
	metav1.GroupVersionResource `json:",inline"`
	corev1.LocalObjectReference `json:",inline"`
}

// ExampleStatus defines the observed state of Example
type ExampleStatus struct {
	Allowed bool `json:"allowed"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Example is the Schema for the examples API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=examples,scope=Namespaced
type Example struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExampleSpec   `json:"spec,omitempty"`
	Status ExampleStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ExampleList contains a list of Example
type ExampleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Example `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Example{}, &ExampleList{})
}
