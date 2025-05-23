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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	context "context"
	time "time"

	apipolicyreportv1alpha2 "github.com/kyverno/kyverno/api/policyreport/v1alpha2"
	versioned "github.com/kyverno/kyverno/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kyverno/kyverno/pkg/client/informers/externalversions/internalinterfaces"
	policyreportv1alpha2 "github.com/kyverno/kyverno/pkg/client/listers/policyreport/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterPolicyReportInformer provides access to a shared informer and lister for
// ClusterPolicyReports.
type ClusterPolicyReportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() policyreportv1alpha2.ClusterPolicyReportLister
}

type clusterPolicyReportInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterPolicyReportInformer constructs a new informer for ClusterPolicyReport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterPolicyReportInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterPolicyReportInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterPolicyReportInformer constructs a new informer for ClusterPolicyReport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterPolicyReportInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Wgpolicyk8sV1alpha2().ClusterPolicyReports().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Wgpolicyk8sV1alpha2().ClusterPolicyReports().Watch(context.TODO(), options)
			},
		},
		&apipolicyreportv1alpha2.ClusterPolicyReport{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterPolicyReportInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterPolicyReportInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterPolicyReportInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apipolicyreportv1alpha2.ClusterPolicyReport{}, f.defaultInformer)
}

func (f *clusterPolicyReportInformer) Lister() policyreportv1alpha2.ClusterPolicyReportLister {
	return policyreportv1alpha2.NewClusterPolicyReportLister(f.Informer().GetIndexer())
}
