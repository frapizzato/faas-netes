/*
Copyright 2019-2021 OpenFaaS Authors

Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "github.com/openfaas/faas-netes/pkg/apis/openfaas/v1"
	openfaasv1 "github.com/openfaas/faas-netes/pkg/client/applyconfiguration/openfaas/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProfiles implements ProfileInterface
type FakeProfiles struct {
	Fake *FakeOpenfaasV1
	ns   string
}

var profilesResource = v1.SchemeGroupVersion.WithResource("profiles")

var profilesKind = v1.SchemeGroupVersion.WithKind("Profile")

// Get takes name of the profile, and returns the corresponding profile object, and an error if there is any.
func (c *FakeProfiles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(profilesResource, c.ns, name), &v1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Profile), err
}

// List takes label and field selectors, and returns the list of Profiles that match those selectors.
func (c *FakeProfiles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(profilesResource, profilesKind, c.ns, opts), &v1.ProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ProfileList{ListMeta: obj.(*v1.ProfileList).ListMeta}
	for _, item := range obj.(*v1.ProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested profiles.
func (c *FakeProfiles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(profilesResource, c.ns, opts))

}

// Create takes the representation of a profile and creates it.  Returns the server's representation of the profile, and an error, if there is any.
func (c *FakeProfiles) Create(ctx context.Context, profile *v1.Profile, opts metav1.CreateOptions) (result *v1.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(profilesResource, c.ns, profile), &v1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Profile), err
}

// Update takes the representation of a profile and updates it. Returns the server's representation of the profile, and an error, if there is any.
func (c *FakeProfiles) Update(ctx context.Context, profile *v1.Profile, opts metav1.UpdateOptions) (result *v1.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(profilesResource, c.ns, profile), &v1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Profile), err
}

// Delete takes name of the profile and deletes it. Returns an error if one occurs.
func (c *FakeProfiles) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(profilesResource, c.ns, name, opts), &v1.Profile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProfiles) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(profilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ProfileList{})
	return err
}

// Patch applies the patch and returns the patched profile.
func (c *FakeProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Profile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(profilesResource, c.ns, name, pt, data, subresources...), &v1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Profile), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied profile.
func (c *FakeProfiles) Apply(ctx context.Context, profile *openfaasv1.ProfileApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Profile, err error) {
	if profile == nil {
		return nil, fmt.Errorf("profile provided to Apply must not be nil")
	}
	data, err := json.Marshal(profile)
	if err != nil {
		return nil, err
	}
	name := profile.Name
	if name == nil {
		return nil, fmt.Errorf("profile.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(profilesResource, c.ns, *name, types.ApplyPatchType, data), &v1.Profile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Profile), err
}
