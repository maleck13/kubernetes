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

	autoscaling_v1 "k8s.io/api/autoscaling/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/listers/autoscaling/v1"
	cache "k8s.io/client-go/tools/cache"
)

// HorizontalPodAutoscalerInformer provides access to a shared informer and lister for
// HorizontalPodAutoscalers.
type HorizontalPodAutoscalerInformer interface {
	Informer(ctx context.Context) cache.SharedIndexInformer
	Lister(ctx context.Context) v1.HorizontalPodAutoscalerLister
}

type horizontalPodAutoscalerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHorizontalPodAutoscalerInformer constructs a new informer for HorizontalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHorizontalPodAutoscalerInformer(ctx context.Context, client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHorizontalPodAutoscalerInformer(ctx, client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHorizontalPodAutoscalerInformer test constructs a new informer for HorizontalPodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHorizontalPodAutoscalerInformer(ctx context.Context, client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1().HorizontalPodAutoscalers(namespace).List(ctx, options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1().HorizontalPodAutoscalers(namespace).Watch(ctx, options)
			},
		},
		&autoscaling_v1.HorizontalPodAutoscaler{},
		resyncPeriod,
		indexers,
	)
}

func (f *horizontalPodAutoscalerInformer) defaultInformer(ctx context.Context, client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHorizontalPodAutoscalerInformer(ctx, client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *horizontalPodAutoscalerInformer) Informer(ctx context.Context) cache.SharedIndexInformer {
	return f.factory.InformerFor(ctx, &autoscaling_v1.HorizontalPodAutoscaler{}, f.defaultInformer)
}

func (f *horizontalPodAutoscalerInformer) Lister(ctx context.Context) v1.HorizontalPodAutoscalerLister {
	return v1.NewHorizontalPodAutoscalerLister(f.Informer(ctx).GetIndexer())
}
