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

package v1

import (
	context "context"
	time "time"

	apireportsv1 "github.com/kyverno/kyverno/api/reports/v1"
	versioned "github.com/kyverno/kyverno/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kyverno/kyverno/pkg/client/informers/externalversions/internalinterfaces"
	reportsv1 "github.com/kyverno/kyverno/pkg/client/listers/reports/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EphemeralReportInformer provides access to a shared informer and lister for
// EphemeralReports.
type EphemeralReportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() reportsv1.EphemeralReportLister
}

type ephemeralReportInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEphemeralReportInformer constructs a new informer for EphemeralReport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEphemeralReportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEphemeralReportInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEphemeralReportInformer constructs a new informer for EphemeralReport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEphemeralReportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReportsV1().EphemeralReports(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReportsV1().EphemeralReports(namespace).Watch(context.TODO(), options)
			},
		},
		&apireportsv1.EphemeralReport{},
		resyncPeriod,
		indexers,
	)
}

func (f *ephemeralReportInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEphemeralReportInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ephemeralReportInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apireportsv1.EphemeralReport{}, f.defaultInformer)
}

func (f *ephemeralReportInformer) Lister() reportsv1.EphemeralReportLister {
	return reportsv1.NewEphemeralReportLister(f.Informer().GetIndexer())
}
