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

package cloud

import (
	infrav1 "github.com/aws/cluster-api-provider-cloudstack/api/v1beta1"
	"github.com/pkg/errors"
)

const (
	AntiAffinityGroupType = "host anti-affinity"
	AffinityGroupType     = "host affinity"
)

type AffinityGroup struct {
	Type string
	Name string
	Id   string
}

type AffinityGroupIFace interface {
	FetchAffinityGroup(*AffinityGroup) error
	GetOrCreateAffinityGroup(*infrav1.CloudStackCluster, *AffinityGroup) error
	DeleteAffinityGroup(*AffinityGroup) error
	AssociateAffinityGroup(*infrav1.CloudStackMachine, AffinityGroup) error
	DissassociateAffinityGroup(*infrav1.CloudStackMachine, AffinityGroup) error
}

func (c *client) FetchAffinityGroup(group *AffinityGroup) (reterr error) {
	if group.Id != "" {
		affinityGroup, count, err := c.cs.AffinityGroup.GetAffinityGroupByID(group.Id)
		if err != nil {
			// handle via multierr
			return err
		} else if count > 1 {
			// handle via creating a new error.
			return errors.New("Count bad")
		} else {
			group.Name = affinityGroup.Name
			group.Type = affinityGroup.Type
			return nil
		}
	}
	if group.Name != "" {
		affinityGroup, count, err := c.cs.AffinityGroup.GetAffinityGroupByName(group.Name)
		if err != nil {
			// handle via multierr
			return err
		} else if count > 1 {
			// handle via creating a new error.
			return errors.New("Count bad")
		} else {
			group.Id = affinityGroup.Id
			group.Type = affinityGroup.Type
			return nil
		}
	}
	return errors.Errorf(`could not fetch AffinityGroup by name "%s" or id "%s"`, group.Name, group.Id)
}

func (c *client) GetOrCreateAffinityGroup(csCluster *infrav1.CloudStackCluster, group *AffinityGroup) (retErr error) {
	if err := c.FetchAffinityGroup(group); err != nil { // Group not found?
		p := c.cs.AffinityGroup.NewCreateAffinityGroupParams(group.Name, group.Type)
		setIfNotEmpty(csCluster.Spec.Account, p.SetAccount)
		setIfNotEmpty(csCluster.Status.DomainID, p.SetDomainid)
		if resp, err := c.cs.AffinityGroup.CreateAffinityGroup(p); err != nil {
			return err
		} else {
			group.Id = resp.Id
		}
	}
	return nil
}

func (c *client) DeleteAffinityGroup(group *AffinityGroup) (retErr error) {
	p := c.cs.AffinityGroup.NewDeleteAffinityGroupParams()
	setIfNotEmpty(group.Id, p.SetId)
	setIfNotEmpty(group.Name, p.SetName)
	_, retErr = c.cs.AffinityGroup.DeleteAffinityGroup(p)
	return retErr
}

type AffinityGroups []AffinityGroup

func (c *client) GetCurrentAffinityGroups(csMachine *infrav1.CloudStackMachine) (AffinityGroups, error) {
	// Start by fetching VM details which includes an array of currently associated affinity groups.
	if virtM, count, err := c.cs.VirtualMachine.GetVirtualMachineByID(*csMachine.Spec.InstanceID); err != nil {
		return nil, err
	} else if count > 1 {
		return nil, errors.Errorf("found more than one VM for ID: %s", *csMachine.Spec.InstanceID)
	} else {
		groups := make([]AffinityGroup, 0, len(virtM.Affinitygroup))
		for _, v := range virtM.Affinitygroup {
			groups = append(groups, AffinityGroup{Name: v.Name, Type: v.Type, Id: v.Id})
		}
		return groups, nil
	}
}

func (ags *AffinityGroups) ToArrayOfIDs() []string {
	groupIds := make([]string, 0, len(*ags))
	for _, group := range *ags {
		groupIds = append(groupIds, group.Id)
	}
	return groupIds
}

func (ags *AffinityGroups) AddGroup(addGroup AffinityGroup) {
	// This is essentially adding to a set followed by array conversion.
	groupSet := map[string]AffinityGroup{addGroup.Id: addGroup}
	for _, group := range *ags {
		groupSet[group.Id] = group
	}
	*ags = make([]AffinityGroup, 0, len(groupSet))
	for _, group := range groupSet {
		*ags = append(*ags, group)
	}
}

func (ags *AffinityGroups) RemoveGroup(removeGroup AffinityGroup) {
	// This is essentially subtracting from a set followed by array conversion.
	groupSet := map[string]AffinityGroup{}
	for _, group := range *ags {
		groupSet[group.Id] = group
	}
	delete(groupSet, removeGroup.Id)
	*ags = make([]AffinityGroup, 0, len(groupSet))
	for _, group := range groupSet {
		*ags = append(*ags, group)
	}
}

func (c *client) StopAndModifyAffinityGroups(csMachine *infrav1.CloudStackMachine, groups AffinityGroups) (retErr error) {
	agp := c.cs.AffinityGroup.NewUpdateVMAffinityGroupParams(*csMachine.Spec.InstanceID)
	agp.SetAffinitygroupids(groups.ToArrayOfIDs())

	p1 := c.cs.VirtualMachine.NewStopVirtualMachineParams(string(*csMachine.Spec.InstanceID))
	if _, err := c.cs.VirtualMachine.StopVirtualMachine(p1); err != nil {
		return err
	}

	if _, err := c.cs.AffinityGroup.UpdateVMAffinityGroup(agp); err != nil {
		return err
	}

	p2 := c.cs.VirtualMachine.NewStartVirtualMachineParams(string(*csMachine.Spec.InstanceID))
	_, err := c.cs.VirtualMachine.StartVirtualMachine(p2)
	return err
}

func (c *client) AssociateAffinityGroup(csMachine *infrav1.CloudStackMachine, group AffinityGroup) (retErr error) {
	if groups, err := c.GetCurrentAffinityGroups(csMachine); err != nil {
		return err
	} else {
		groups.AddGroup(group)
		return c.StopAndModifyAffinityGroups(csMachine, groups)
	}
}

func (c *client) DissassociateAffinityGroup(csMachine *infrav1.CloudStackMachine, group AffinityGroup) (retErr error) {
	if groups, err := c.GetCurrentAffinityGroups(csMachine); err != nil {
		return err
	} else {
		groups.RemoveGroup(group)
		return c.StopAndModifyAffinityGroups(csMachine, groups)
	}
}
