/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	context "context"
	json "encoding/json"
	fmt "fmt"

	v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	extensionsv1beta1 "k8s.io/client-go/applyconfigurations/extensions/v1beta1"
	gentype "k8s.io/client-go/gentype"
	typedextensionsv1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	testing "k8s.io/client-go/testing"
)

// fakeReplicaSets implements ReplicaSetInterface
type fakeReplicaSets struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.ReplicaSet, *v1beta1.ReplicaSetList, *extensionsv1beta1.ReplicaSetApplyConfiguration]
	Fake *FakeExtensionsV1beta1
}

func newFakeReplicaSets(fake *FakeExtensionsV1beta1, namespace string) typedextensionsv1beta1.ReplicaSetInterface {
	return &fakeReplicaSets{
		gentype.NewFakeClientWithListAndApply[*v1beta1.ReplicaSet, *v1beta1.ReplicaSetList, *extensionsv1beta1.ReplicaSetApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("replicasets"),
			v1beta1.SchemeGroupVersion.WithKind("ReplicaSet"),
			func() *v1beta1.ReplicaSet { return &v1beta1.ReplicaSet{} },
			func() *v1beta1.ReplicaSetList { return &v1beta1.ReplicaSetList{} },
			func(dst, src *v1beta1.ReplicaSetList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.ReplicaSetList) []*v1beta1.ReplicaSet { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.ReplicaSetList, items []*v1beta1.ReplicaSet) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}

// GetScale takes name of the replicaSet, and returns the corresponding scale object, and an error if there is any.
func (c *fakeReplicaSets) GetScale(ctx context.Context, replicaSetName string, options v1.GetOptions) (result *v1beta1.Scale, err error) {
	emptyResult := &v1beta1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceActionWithOptions(c.Resource(), c.Namespace(), "scale", replicaSetName, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *fakeReplicaSets) UpdateScale(ctx context.Context, replicaSetName string, scale *v1beta1.Scale, opts v1.UpdateOptions) (result *v1beta1.Scale, err error) {
	emptyResult := &v1beta1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(c.Resource(), "scale", c.Namespace(), scale, opts), &v1beta1.Scale{})

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.Scale), err
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *fakeReplicaSets) ApplyScale(ctx context.Context, replicaSetName string, scale *extensionsv1beta1.ScaleApplyConfiguration, opts v1.ApplyOptions) (result *v1beta1.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}
	emptyResult := &v1beta1.Scale{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(c.Resource(), c.Namespace(), replicaSetName, types.ApplyPatchType, data, opts.ToPatchOptions(), "scale"), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1beta1.Scale), err
}
