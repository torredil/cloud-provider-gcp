/*
Copyright 2022 The Kubernetes Authors.

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
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	networkv1 "k8s.io/cloud-provider-gcp/crd/apis/network/v1"
	versioned "k8s.io/cloud-provider-gcp/crd/client/network/clientset/versioned"
	internalinterfaces "k8s.io/cloud-provider-gcp/crd/client/network/informers/externalversions/internalinterfaces"
	v1 "k8s.io/cloud-provider-gcp/crd/client/network/listers/network/v1"
)

// NetworkInformer provides access to a shared informer and lister for
// Networks.
type NetworkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NetworkLister
}

type networkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewNetworkInformer constructs a new informer for Network type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkInformer constructs a new informer for Network type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().Networks().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().Networks().Watch(context.TODO(), options)
			},
		},
		&networkv1.Network{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkv1.Network{}, f.defaultInformer)
}

func (f *networkInformer) Lister() v1.NetworkLister {
	return v1.NewNetworkLister(f.Informer().GetIndexer())
}
