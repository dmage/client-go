// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	helmv1beta1 "github.com/openshift/api/helm/v1beta1"
	versioned "github.com/openshift/client-go/helm/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/helm/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/openshift/client-go/helm/listers/helm/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ProjectHelmChartRepositoryInformer provides access to a shared informer and lister for
// ProjectHelmChartRepositories.
type ProjectHelmChartRepositoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ProjectHelmChartRepositoryLister
}

type projectHelmChartRepositoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewProjectHelmChartRepositoryInformer constructs a new informer for ProjectHelmChartRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewProjectHelmChartRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredProjectHelmChartRepositoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredProjectHelmChartRepositoryInformer constructs a new informer for ProjectHelmChartRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredProjectHelmChartRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HelmV1beta1().ProjectHelmChartRepositories(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HelmV1beta1().ProjectHelmChartRepositories(namespace).Watch(context.TODO(), options)
			},
		},
		&helmv1beta1.ProjectHelmChartRepository{},
		resyncPeriod,
		indexers,
	)
}

func (f *projectHelmChartRepositoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredProjectHelmChartRepositoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *projectHelmChartRepositoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&helmv1beta1.ProjectHelmChartRepository{}, f.defaultInformer)
}

func (f *projectHelmChartRepositoryInformer) Lister() v1beta1.ProjectHelmChartRepositoryLister {
	return v1beta1.NewProjectHelmChartRepositoryLister(f.Informer().GetIndexer())
}