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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	schedulingv1 "k8s.io/api/scheduling/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// PriorityClassLister helps list PriorityClasses.
// All objects returned here must be treated as read-only.
type PriorityClassLister interface {
	// List lists all PriorityClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*schedulingv1.PriorityClass, err error)
	// Get retrieves the PriorityClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*schedulingv1.PriorityClass, error)
	PriorityClassListerExpansion
}

// priorityClassLister implements the PriorityClassLister interface.
type priorityClassLister struct {
	listers.ResourceIndexer[*schedulingv1.PriorityClass]
}

// NewPriorityClassLister returns a new PriorityClassLister.
func NewPriorityClassLister(indexer cache.Indexer) PriorityClassLister {
	return &priorityClassLister{listers.New[*schedulingv1.PriorityClass](indexer, schedulingv1.Resource("priorityclass"))}
}
