//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Embedded) DeepCopyInto(out *Embedded) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Embedded1 = in.Embedded1
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Embedded.
func (in *Embedded) DeepCopy() *Embedded {
	if in == nil {
		return nil
	}
	out := new(Embedded)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Embedded) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Embedded1) DeepCopyInto(out *Embedded1) {
	*out = *in
	out.Embedded2 = in.Embedded2
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Embedded1.
func (in *Embedded1) DeepCopy() *Embedded1 {
	if in == nil {
		return nil
	}
	out := new(Embedded1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Embedded2) DeepCopyInto(out *Embedded2) {
	*out = *in
	out.Embedded3 = in.Embedded3
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Embedded2.
func (in *Embedded2) DeepCopy() *Embedded2 {
	if in == nil {
		return nil
	}
	out := new(Embedded2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Embedded3) DeepCopyInto(out *Embedded3) {
	*out = *in
	out.Embedded4 = in.Embedded4
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Embedded3.
func (in *Embedded3) DeepCopy() *Embedded3 {
	if in == nil {
		return nil
	}
	out := new(Embedded3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Embedded4) DeepCopyInto(out *Embedded4) {
	*out = *in
	out.EmbeddedX = in.EmbeddedX
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Embedded4.
func (in *Embedded4) DeepCopy() *Embedded4 {
	if in == nil {
		return nil
	}
	out := new(Embedded4)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedX) DeepCopyInto(out *EmbeddedX) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedX.
func (in *EmbeddedX) DeepCopy() *EmbeddedX {
	if in == nil {
		return nil
	}
	out := new(EmbeddedX)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Guestbook) DeepCopyInto(out *Guestbook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Guestbook.
func (in *Guestbook) DeepCopy() *Guestbook {
	if in == nil {
		return nil
	}
	out := new(Guestbook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Guestbook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuestbookEntry) DeepCopyInto(out *GuestbookEntry) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuestbookEntry.
func (in *GuestbookEntry) DeepCopy() *GuestbookEntry {
	if in == nil {
		return nil
	}
	out := new(GuestbookEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuestbookList) DeepCopyInto(out *GuestbookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Guestbook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuestbookList.
func (in *GuestbookList) DeepCopy() *GuestbookList {
	if in == nil {
		return nil
	}
	out := new(GuestbookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GuestbookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuestbookSpec) DeepCopyInto(out *GuestbookSpec) {
	*out = *in
	if in.Page != nil {
		in, out := &in.Page, &out.Page
		*out = new(int)
		**out = **in
	}
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make([]GuestbookEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Selector.DeepCopyInto(&out.Selector)
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]GuestbookHeader, len(*in))
		copy(*out, *in)
	}
	in.CertificateRef.DeepCopyInto(&out.CertificateRef)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuestbookSpec.
func (in *GuestbookSpec) DeepCopy() *GuestbookSpec {
	if in == nil {
		return nil
	}
	out := new(GuestbookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GuestbookStatus) DeepCopyInto(out *GuestbookStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GuestbookStatus.
func (in *GuestbookStatus) DeepCopy() *GuestbookStatus {
	if in == nil {
		return nil
	}
	out := new(GuestbookStatus)
	in.DeepCopyInto(out)
	return out
}
