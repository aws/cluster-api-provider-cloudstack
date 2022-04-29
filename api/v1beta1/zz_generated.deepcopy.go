//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1beta1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackAffinityGroup) DeepCopyInto(out *CloudStackAffinityGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackAffinityGroup.
func (in *CloudStackAffinityGroup) DeepCopy() *CloudStackAffinityGroup {
	if in == nil {
		return nil
	}
	out := new(CloudStackAffinityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackAffinityGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackAffinityGroupList) DeepCopyInto(out *CloudStackAffinityGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudStackAffinityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackAffinityGroupList.
func (in *CloudStackAffinityGroupList) DeepCopy() *CloudStackAffinityGroupList {
	if in == nil {
		return nil
	}
	out := new(CloudStackAffinityGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackAffinityGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackAffinityGroupSpec) DeepCopyInto(out *CloudStackAffinityGroupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackAffinityGroupSpec.
func (in *CloudStackAffinityGroupSpec) DeepCopy() *CloudStackAffinityGroupSpec {
	if in == nil {
		return nil
	}
	out := new(CloudStackAffinityGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackAffinityGroupStatus) DeepCopyInto(out *CloudStackAffinityGroupStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackAffinityGroupStatus.
func (in *CloudStackAffinityGroupStatus) DeepCopy() *CloudStackAffinityGroupStatus {
	if in == nil {
		return nil
	}
	out := new(CloudStackAffinityGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackCluster) DeepCopyInto(out *CloudStackCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackCluster.
func (in *CloudStackCluster) DeepCopy() *CloudStackCluster {
	if in == nil {
		return nil
	}
	out := new(CloudStackCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackClusterList) DeepCopyInto(out *CloudStackClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudStackCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackClusterList.
func (in *CloudStackClusterList) DeepCopy() *CloudStackClusterList {
	if in == nil {
		return nil
	}
	out := new(CloudStackClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackClusterSpec) DeepCopyInto(out *CloudStackClusterSpec) {
	*out = *in
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]Zone, len(*in))
		copy(*out, *in)
	}
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
	if in.IdentityRef != nil {
		in, out := &in.IdentityRef, &out.IdentityRef
		*out = new(CloudStackIdentityReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackClusterSpec.
func (in *CloudStackClusterSpec) DeepCopy() *CloudStackClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CloudStackClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackClusterStatus) DeepCopyInto(out *CloudStackClusterStatus) {
	*out = *in
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make(ZoneStatusMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FailureDomains != nil {
		in, out := &in.FailureDomains, &out.FailureDomains
		*out = make(apiv1beta1.FailureDomains, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackClusterStatus.
func (in *CloudStackClusterStatus) DeepCopy() *CloudStackClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CloudStackClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackIdentityReference) DeepCopyInto(out *CloudStackIdentityReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackIdentityReference.
func (in *CloudStackIdentityReference) DeepCopy() *CloudStackIdentityReference {
	if in == nil {
		return nil
	}
	out := new(CloudStackIdentityReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackIsolatedNetwork) DeepCopyInto(out *CloudStackIsolatedNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackIsolatedNetwork.
func (in *CloudStackIsolatedNetwork) DeepCopy() *CloudStackIsolatedNetwork {
	if in == nil {
		return nil
	}
	out := new(CloudStackIsolatedNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackIsolatedNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackIsolatedNetworkList) DeepCopyInto(out *CloudStackIsolatedNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudStackIsolatedNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackIsolatedNetworkList.
func (in *CloudStackIsolatedNetworkList) DeepCopy() *CloudStackIsolatedNetworkList {
	if in == nil {
		return nil
	}
	out := new(CloudStackIsolatedNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackIsolatedNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackIsolatedNetworkSpec) DeepCopyInto(out *CloudStackIsolatedNetworkSpec) {
	*out = *in
	out.CloudStackResourceIdentifier = in.CloudStackResourceIdentifier
	out.ControlPlaneEndpoint = in.ControlPlaneEndpoint
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackIsolatedNetworkSpec.
func (in *CloudStackIsolatedNetworkSpec) DeepCopy() *CloudStackIsolatedNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(CloudStackIsolatedNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackIsolatedNetworkStatus) DeepCopyInto(out *CloudStackIsolatedNetworkStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackIsolatedNetworkStatus.
func (in *CloudStackIsolatedNetworkStatus) DeepCopy() *CloudStackIsolatedNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(CloudStackIsolatedNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackMachine) DeepCopyInto(out *CloudStackMachine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackMachine.
func (in *CloudStackMachine) DeepCopy() *CloudStackMachine {
	if in == nil {
		return nil
	}
	out := new(CloudStackMachine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackMachine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackMachineList) DeepCopyInto(out *CloudStackMachineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudStackMachine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackMachineList.
func (in *CloudStackMachineList) DeepCopy() *CloudStackMachineList {
	if in == nil {
		return nil
	}
	out := new(CloudStackMachineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackMachineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackMachineSpec) DeepCopyInto(out *CloudStackMachineSpec) {
	*out = *in
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	out.Offering = in.Offering
	out.Template = in.Template
	if in.Details != nil {
		in, out := &in.Details, &out.Details
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AffinityGroupIDs != nil {
		in, out := &in.AffinityGroupIDs, &out.AffinityGroupIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AffinityGroupRef != nil {
		in, out := &in.AffinityGroupRef, &out.AffinityGroupRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	if in.ProviderID != nil {
		in, out := &in.ProviderID, &out.ProviderID
		*out = new(string)
		**out = **in
	}
	out.CloudStackResourceIdentifier = in.CloudStackResourceIdentifier
	if in.IdentityRef != nil {
		in, out := &in.IdentityRef, &out.IdentityRef
		*out = new(CloudStackIdentityReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackMachineSpec.
func (in *CloudStackMachineSpec) DeepCopy() *CloudStackMachineSpec {
	if in == nil {
		return nil
	}
	out := new(CloudStackMachineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackMachineStatus) DeepCopyInto(out *CloudStackMachineStatus) {
	*out = *in
	if in.Addresses != nil {
		in, out := &in.Addresses, &out.Addresses
		*out = make([]v1.NodeAddress, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackMachineStatus.
func (in *CloudStackMachineStatus) DeepCopy() *CloudStackMachineStatus {
	if in == nil {
		return nil
	}
	out := new(CloudStackMachineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackMachineTemplate) DeepCopyInto(out *CloudStackMachineTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackMachineTemplate.
func (in *CloudStackMachineTemplate) DeepCopy() *CloudStackMachineTemplate {
	if in == nil {
		return nil
	}
	out := new(CloudStackMachineTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackMachineTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackMachineTemplateList) DeepCopyInto(out *CloudStackMachineTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudStackMachineTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackMachineTemplateList.
func (in *CloudStackMachineTemplateList) DeepCopy() *CloudStackMachineTemplateList {
	if in == nil {
		return nil
	}
	out := new(CloudStackMachineTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackMachineTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackMachineTemplateResource) DeepCopyInto(out *CloudStackMachineTemplateResource) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackMachineTemplateResource.
func (in *CloudStackMachineTemplateResource) DeepCopy() *CloudStackMachineTemplateResource {
	if in == nil {
		return nil
	}
	out := new(CloudStackMachineTemplateResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackMachineTemplateSpec) DeepCopyInto(out *CloudStackMachineTemplateSpec) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackMachineTemplateSpec.
func (in *CloudStackMachineTemplateSpec) DeepCopy() *CloudStackMachineTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(CloudStackMachineTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackResourceIdentifier) DeepCopyInto(out *CloudStackResourceIdentifier) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackResourceIdentifier.
func (in *CloudStackResourceIdentifier) DeepCopy() *CloudStackResourceIdentifier {
	if in == nil {
		return nil
	}
	out := new(CloudStackResourceIdentifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackZone) DeepCopyInto(out *CloudStackZone) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackZone.
func (in *CloudStackZone) DeepCopy() *CloudStackZone {
	if in == nil {
		return nil
	}
	out := new(CloudStackZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackZone) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackZoneList) DeepCopyInto(out *CloudStackZoneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudStackZone, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackZoneList.
func (in *CloudStackZoneList) DeepCopy() *CloudStackZoneList {
	if in == nil {
		return nil
	}
	out := new(CloudStackZoneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudStackZoneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackZoneSpec) DeepCopyInto(out *CloudStackZoneSpec) {
	*out = *in
	out.CloudStackResourceIdentifier = in.CloudStackResourceIdentifier
	out.Network = in.Network
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackZoneSpec.
func (in *CloudStackZoneSpec) DeepCopy() *CloudStackZoneSpec {
	if in == nil {
		return nil
	}
	out := new(CloudStackZoneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudStackZoneStatus) DeepCopyInto(out *CloudStackZoneStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudStackZoneStatus.
func (in *CloudStackZoneStatus) DeepCopy() *CloudStackZoneStatus {
	if in == nil {
		return nil
	}
	out := new(CloudStackZoneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Zone) DeepCopyInto(out *Zone) {
	*out = *in
	out.CloudStackResourceIdentifier = in.CloudStackResourceIdentifier
	out.Network = in.Network
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Zone.
func (in *Zone) DeepCopy() *Zone {
	if in == nil {
		return nil
	}
	out := new(Zone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ZoneStatusMap) DeepCopyInto(out *ZoneStatusMap) {
	{
		in := &in
		*out = make(ZoneStatusMap, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZoneStatusMap.
func (in ZoneStatusMap) DeepCopy() ZoneStatusMap {
	if in == nil {
		return nil
	}
	out := new(ZoneStatusMap)
	in.DeepCopyInto(out)
	return *out
}
