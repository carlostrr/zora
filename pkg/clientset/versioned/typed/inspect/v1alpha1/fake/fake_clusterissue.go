// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/getupio-undistro/inspect/apis/inspect/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterIssues implements ClusterIssueInterface
type FakeClusterIssues struct {
	Fake *FakeInspectV1alpha1
	ns   string
}

var clusterissuesResource = schema.GroupVersionResource{Group: "inspect", Version: "v1alpha1", Resource: "clusterissues"}

var clusterissuesKind = schema.GroupVersionKind{Group: "inspect", Version: "v1alpha1", Kind: "ClusterIssue"}

// Get takes name of the clusterIssue, and returns the corresponding clusterIssue object, and an error if there is any.
func (c *FakeClusterIssues) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterIssue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterissuesResource, c.ns, name), &v1alpha1.ClusterIssue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIssue), err
}

// List takes label and field selectors, and returns the list of ClusterIssues that match those selectors.
func (c *FakeClusterIssues) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterIssueList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterissuesResource, clusterissuesKind, c.ns, opts), &v1alpha1.ClusterIssueList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterIssueList{ListMeta: obj.(*v1alpha1.ClusterIssueList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterIssueList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterIssues.
func (c *FakeClusterIssues) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterissuesResource, c.ns, opts))

}

// Create takes the representation of a clusterIssue and creates it.  Returns the server's representation of the clusterIssue, and an error, if there is any.
func (c *FakeClusterIssues) Create(ctx context.Context, clusterIssue *v1alpha1.ClusterIssue, opts v1.CreateOptions) (result *v1alpha1.ClusterIssue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterissuesResource, c.ns, clusterIssue), &v1alpha1.ClusterIssue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIssue), err
}

// Update takes the representation of a clusterIssue and updates it. Returns the server's representation of the clusterIssue, and an error, if there is any.
func (c *FakeClusterIssues) Update(ctx context.Context, clusterIssue *v1alpha1.ClusterIssue, opts v1.UpdateOptions) (result *v1alpha1.ClusterIssue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterissuesResource, c.ns, clusterIssue), &v1alpha1.ClusterIssue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIssue), err
}

// Delete takes name of the clusterIssue and deletes it. Returns an error if one occurs.
func (c *FakeClusterIssues) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusterissuesResource, c.ns, name, opts), &v1alpha1.ClusterIssue{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterIssues) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterissuesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterIssueList{})
	return err
}

// Patch applies the patch and returns the patched clusterIssue.
func (c *FakeClusterIssues) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterIssue, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterissuesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterIssue{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterIssue), err
}
