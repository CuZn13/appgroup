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
	time "time"

	appv1 "github.com/cuzn/appgroup/pkg/apis/app/v1"
	versioned "github.com/cuzn/appgroup/pkg/client/clientset/versioned"
	internalinterfaces "github.com/cuzn/appgroup/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/cuzn/appgroup/pkg/client/listers/app/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AppGroupInformer provides access to a shared informer and lister for
// AppGroups.
type AppGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AppGroupLister
}

type appGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAppGroupInformer constructs a new informer for AppGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAppGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAppGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAppGroupInformer constructs a new informer for AppGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAppGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppV1().AppGroups(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppV1().AppGroups(namespace).Watch(options)
			},
		},
		&appv1.AppGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *appGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAppGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *appGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appv1.AppGroup{}, f.defaultInformer)
}

func (f *appGroupInformer) Lister() v1.AppGroupLister {
	return v1.NewAppGroupLister(f.Informer().GetIndexer())
}
