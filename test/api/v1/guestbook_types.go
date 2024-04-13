// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package v1

import (
	"github.com/elastic/crd-ref-docs/api/common"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

//+kubebuilder:object:root=true

type Embedded struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	A         string `json:"a,omitempty"`
	Embedded1 `json:",inline"`
}
type Embedded1 struct {
	Embedded2 `json:",inline"`
	// +hidefromdoc
	E         string `json:"e,omitempty"`
	EmbeddedX `json:",inline"`
}
type Embedded2 struct {
	B         string `json:"b,omitempty"`
	Embedded3 `json:",inline"`
}
type Embedded3 struct {
	Embedded4 `json:",inline"`
	D         string `json:"d,omitempty"`
}
type Embedded4 struct {
	C string `json:"c,omitempty"`
}
type EmbeddedX struct {
	X string `json:"x,omitempty"`
}

// Underlying tests that Underlying1's underlying type is Underlying2 instead of string.
// +kubebuilder:object:root=true
type Underlying struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:default="b"
	A Underlying1 `json:"a,omitempty"`
}

// Underlying1 has an underlying type with an underlying type
type Underlying1 Underlying2

// Underlying2 is a string alias
// +kubebuilder:validation:MaxLength=10
type Underlying2 string

// NOTE: Rating is placed here to ensure that it is parsed as a standalone type
// before it is parsed as a struct field.

// Rating is the rating provided by a guest.
// +kubebuilder:validation:Maximum=4
// +kubebuilder:validation:Minimum=1
// +kubebuilder:validation:Maximum=5
type Rating int

// GuestbookSpec defines the desired state of Guestbook.
// +kubebuilder:validation:XValidation:rule="self.page < 200", message="Please start a new book."
type GuestbookSpec struct {
	// Page indicates the page number
	// +kubebuilder:default=1
	// +kubebuilder:example=3
	Page *PositiveInt `json:"page,omitempty"`
	// Entries contain guest book entries for the page
	Entries []GuestbookEntry `json:"entries,omitempty"`
	// Selector selects something
	Selector metav1.LabelSelector `json:"selector,omitempty"`
	// Headers contains a list of header items to include in the page
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	Headers []GuestbookHeader `json:"headers,omitempty"`
	// CertificateRef is a reference to a secret containing a certificate
	CertificateRef gwapiv1b1.SecretObjectReference `json:"certificateRef"`
	String         common.CommonString             `json:"str"`
	// Enumeration is an example of an aliased enumeration type
	Enumeration MyEnum `json:"enum"`
}

// +kubebuilder:validation:Enum=MyFirstValue;MySecondValue
type MyEnum string

const (
	// MyFirstValue is an interesting value to use
	MyFirstValue MyEnum = "MyFirstValue"
	// MySecondValue is what you use when you can't use MyFirstValue
	MySecondValue MyEnum = "MySecondValue"
)

// +kubebuilder:validation:Minimum=1
type PositiveInt int

// GuestbookEntry defines an entry in a guest book.
type GuestbookEntry struct {
	// Name of the guest (pipe | should be escaped)
	// +kubebuilder:validation:MaxLength=80
	// +kubebuilder:validation:Pattern=`0*[a-z0-9]*[a-z]*[0-9]`
	Name string `json:"name,omitempty"`
	// Time of entry
	Time metav1.Time `json:"time,omitempty"`
	// Comment by guest. This can be a multi-line comment.
	// Like this one.
	// Now let's test a list:
	// * a
	// * b
	//
	// Another isolated comment.
	//
	// Looks good?
	//
	// +kubebuilder:validation:Pattern=`0*[a-z0-9]*[a-z]*[0-9]*`
	Comment string `json:"comment,omitempty"`
	// Rating provided by the guest
	Rating Rating `json:"rating,omitempty"`
}

// GuestbookStatus defines the observed state of Guestbook.
type GuestbookStatus struct {
	// +kubebuilder:validation:Enum={OK, Error}
	Status Status              `json:"status"`
	String common.CommonString `json:"str"`
}

// +kubebuilder:validation:Enum={OK, Unknown, Error}
type Status string

// GuestbookHeaders are strings to include at the top of a page.
type GuestbookHeader string

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Guestbook is the Schema for the guestbooks API.
type Guestbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +kubebuilder:default={page: 1}
	Spec   GuestbookSpec   `json:"spec,omitempty"`
	Status GuestbookStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GuestbookList contains a list of Guestbook.
type GuestbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Guestbook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Guestbook{}, &GuestbookList{})
}
