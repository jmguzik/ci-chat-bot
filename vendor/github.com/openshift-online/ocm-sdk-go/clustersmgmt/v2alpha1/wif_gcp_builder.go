/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v2alpha1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v2alpha1

// WifGcpBuilder contains the data and logic needed to build 'wif_gcp' objects.
type WifGcpBuilder struct {
	bitmap_              uint32
	impersonatorEmail    string
	projectId            string
	serviceAccounts      []*WifServiceAccountBuilder
	workloadIdentityPool *WifPoolBuilder
}

// NewWifGcp creates a new builder of 'wif_gcp' objects.
func NewWifGcp() *WifGcpBuilder {
	return &WifGcpBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *WifGcpBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// ImpersonatorEmail sets the value of the 'impersonator_email' attribute to the given value.
func (b *WifGcpBuilder) ImpersonatorEmail(value string) *WifGcpBuilder {
	b.impersonatorEmail = value
	b.bitmap_ |= 1
	return b
}

// ProjectId sets the value of the 'project_id' attribute to the given value.
func (b *WifGcpBuilder) ProjectId(value string) *WifGcpBuilder {
	b.projectId = value
	b.bitmap_ |= 2
	return b
}

// ServiceAccounts sets the value of the 'service_accounts' attribute to the given values.
func (b *WifGcpBuilder) ServiceAccounts(values ...*WifServiceAccountBuilder) *WifGcpBuilder {
	b.serviceAccounts = make([]*WifServiceAccountBuilder, len(values))
	copy(b.serviceAccounts, values)
	b.bitmap_ |= 4
	return b
}

// WorkloadIdentityPool sets the value of the 'workload_identity_pool' attribute to the given value.
func (b *WifGcpBuilder) WorkloadIdentityPool(value *WifPoolBuilder) *WifGcpBuilder {
	b.workloadIdentityPool = value
	if value != nil {
		b.bitmap_ |= 8
	} else {
		b.bitmap_ &^= 8
	}
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *WifGcpBuilder) Copy(object *WifGcp) *WifGcpBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.impersonatorEmail = object.impersonatorEmail
	b.projectId = object.projectId
	if object.serviceAccounts != nil {
		b.serviceAccounts = make([]*WifServiceAccountBuilder, len(object.serviceAccounts))
		for i, v := range object.serviceAccounts {
			b.serviceAccounts[i] = NewWifServiceAccount().Copy(v)
		}
	} else {
		b.serviceAccounts = nil
	}
	if object.workloadIdentityPool != nil {
		b.workloadIdentityPool = NewWifPool().Copy(object.workloadIdentityPool)
	} else {
		b.workloadIdentityPool = nil
	}
	return b
}

// Build creates a 'wif_gcp' object using the configuration stored in the builder.
func (b *WifGcpBuilder) Build() (object *WifGcp, err error) {
	object = new(WifGcp)
	object.bitmap_ = b.bitmap_
	object.impersonatorEmail = b.impersonatorEmail
	object.projectId = b.projectId
	if b.serviceAccounts != nil {
		object.serviceAccounts = make([]*WifServiceAccount, len(b.serviceAccounts))
		for i, v := range b.serviceAccounts {
			object.serviceAccounts[i], err = v.Build()
			if err != nil {
				return
			}
		}
	}
	if b.workloadIdentityPool != nil {
		object.workloadIdentityPool, err = b.workloadIdentityPool.Build()
		if err != nil {
			return
		}
	}
	return
}
