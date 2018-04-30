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

package internalversion

import (
	context "context"
	time "time"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	internalclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/internalclientset"
	internalinterfaces "k8s.io/apiextensions-apiserver/pkg/client/informers/internalversion/internalinterfaces"
	internalversion "k8s.io/apiextensions-apiserver/pkg/client/listers/apiextensions/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CustomResourceDefinitionInformer provides access to a shared informer and lister for
// CustomResourceDefinitions.
type CustomResourceDefinitionInformer interface {
	Informer(ctx context.Context) cache.SharedIndexInformer
	Lister(ctx context.Context) internalversion.CustomResourceDefinitionLister
}

type customResourceDefinitionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCustomResourceDefinitionInformer constructs a new informer for CustomResourceDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCustomResourceDefinitionInformer(ctx context.Context, client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCustomResourceDefinitionInformer(ctx, client, resyncPeriod, indexers, nil)
}

// NewFilteredCustomResourceDefinitionInformer test constructs a new informer for CustomResourceDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCustomResourceDefinitionInformer(ctx context.Context, client internalclientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Apiextensions().CustomResourceDefinitions().List(ctx, options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Apiextensions().CustomResourceDefinitions().Watch(ctx, options)
			},
		},
		&apiextensions.CustomResourceDefinition{},
		resyncPeriod,
		indexers,
	)
}

func (f *customResourceDefinitionInformer) defaultInformer(ctx context.Context, client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCustomResourceDefinitionInformer(ctx, client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *customResourceDefinitionInformer) Informer(ctx context.Context) cache.SharedIndexInformer {
	return f.factory.InformerFor(ctx, &apiextensions.CustomResourceDefinition{}, f.defaultInformer)
}

func (f *customResourceDefinitionInformer) Lister(ctx context.Context) internalversion.CustomResourceDefinitionLister {
	return internalversion.NewCustomResourceDefinitionLister(f.Informer(ctx).GetIndexer())
}
