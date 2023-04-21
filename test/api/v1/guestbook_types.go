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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true

type Embedded struct {
	A         string `json:"a,omitempty"`
	Embedded1 `json:",inline"`
}
type Embedded1 struct {
	Embedded2 `json:",inline"`
	E         string `json:"e,omitempty"`
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
	C         string `json:"c,omitempty"`
	EmbeddedX `json:",inline"`
}
type EmbeddedX struct {
	X string `json:"x,omitempty"`
}

// GuestbookSpec defines the desired state of Guestbook.
type GuestbookSpec struct {
	// Page indicates the page number
	Page *int `json:"page,omitempty"`
	// Entries contain guest book entries for the page
	Entries []GuestbookEntry `json:"entries,omitempty"`
	// Selector selects something
	Selector metav1.LabelSelector `json:"selector,omitempty"`
	// Headers contains a list of header items to include in the page
	Headers []GuestbookHeader `json:"headers,omitempty"`
}

// GuestbookEntry defines an entry in a guest book.
type GuestbookEntry struct {
	// Name of the guest (pipe | should be escaped)
	Name string `json:"name,omitempty"`
	// Time of entry
	Time metav1.Time `json:"time,omitempty"`
	// Comment by guest
	Comment string `json:"comment,omitempty"`
	// Rating provided by the guest
	Rating Rating `json:"rating,omitempty"`
}

// GuestbookStatus defines the observed state of Guestbook.
type GuestbookStatus struct {
	Status string `json:"status"`
}

// GuestbookHeaders are strings to include at the top of a page.
type GuestbookHeader string

// +kubebuilder:object:root=true

// Guestbook is the Schema for the guestbooks API.
type Guestbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuestbookSpec   `json:"spec,omitempty"`
	Status GuestbookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuestbookList contains a list of Guestbook.
type GuestbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Guestbook `json:"items"`
}

// Rating is the rating provided by a guest.
type Rating string

func init() {
	SchemeBuilder.Register(&Guestbook{}, &GuestbookList{})
}
