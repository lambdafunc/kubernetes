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
	gentype "k8s.io/client-go/gentype"
	v1alpha1 "k8s.io/metrics/pkg/apis/metrics/v1alpha1"
	metricsv1alpha1 "k8s.io/metrics/pkg/client/clientset/versioned/typed/metrics/v1alpha1"
)

// fakeNodeMetricses implements NodeMetricsInterface
type fakeNodeMetricses struct {
	*gentype.FakeClientWithList[*v1alpha1.NodeMetrics, *v1alpha1.NodeMetricsList]
	Fake *FakeMetricsV1alpha1
}

func newFakeNodeMetricses(fake *FakeMetricsV1alpha1) metricsv1alpha1.NodeMetricsInterface {
	return &fakeNodeMetricses{
		gentype.NewFakeClientWithList[*v1alpha1.NodeMetrics, *v1alpha1.NodeMetricsList](
			fake.Fake,
			"",
			v1alpha1.SchemeGroupVersion.WithResource("nodes"),
			v1alpha1.SchemeGroupVersion.WithKind("NodeMetrics"),
			func() *v1alpha1.NodeMetrics { return &v1alpha1.NodeMetrics{} },
			func() *v1alpha1.NodeMetricsList { return &v1alpha1.NodeMetricsList{} },
			func(dst, src *v1alpha1.NodeMetricsList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.NodeMetricsList) []*v1alpha1.NodeMetrics {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.NodeMetricsList, items []*v1alpha1.NodeMetrics) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
