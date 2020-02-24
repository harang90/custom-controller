package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	GroupName    string = "harang90.github.io"
	Kind         string = "TestResource"
	GroupVersion string = "v1beta1"
	Plural       string = "testresources"
	Singular     string = "testresource"
	Name         string = Plural + "." + GroupName
)

type TestResourceSpec struct {
	Command        string `json:"command"`
	CustomProperty string `json:"customproperty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status TestResourceStatus `json:"status"`
	Spec   TestResourceSpec   `json:"spec"`
}

type TestResourceStatus struct {
	Name string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type TestResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TestResource `json:"items"`
}
