/*
Copyright The Karpor Authors.

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

package topology

import (
	"context"

	"github.com/KusionStack/karpor/pkg/util/ctxutil"
	topologyutil "github.com/KusionStack/karpor/pkg/util/topology"
	"github.com/dominikbraun/graph"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

// GetByJSONPath retrieves related resources based on JSON path from a given list of unstructured resources.
func GetByJSONPath(
	ctx context.Context,
	relatedResList *unstructured.UnstructuredList,
	relationshipType string,
	client *dynamic.DynamicClient,
	obj unstructured.Unstructured,
	relation *Relationship,
	relatedGVK schema.GroupVersionKind,
	objResourceNode ResourceGraphNode,
	relationshipGraph graph.Graph[string, RelationshipGraphNode],
	resourceGraph graph.Graph[string, ResourceGraphNode],
) (graph.Graph[string, ResourceGraphNode], error) {
	log := ctxutil.GetLogger(ctx)

	log.Info("Using direct references to find related resources...")
	var jpMatch bool
	var err error
	for _, relatedRes := range relatedResList.Items {
		if relation.AutoGenerated {
			jpMatch, err = topologyutil.JSONPathMatch(relatedRes, obj, relation.JSONPath)
		} else {
			jpMatch, err = topologyutil.JSONPathMatch(obj, relatedRes, relation.JSONPath)
		}
		if jpMatch && err == nil {
			log.Info("Resource found based on JSONPath.", "relationshipType", relationshipType, "kind", obj.GetKind(), "name", obj.GetName())
			log.Info("Resource is:", "relationshipType", relationshipType, "kind", relatedRes.GetKind(), "name", relatedRes.GetName())
			log.Info("---------------------------------------------------------------------------")
			rgv, _ := schema.ParseGroupVersion(relatedRes.GetAPIVersion())
			relatedResourceNode := ResourceGraphNode{
				Group:     rgv.Group,
				Version:   rgv.Version,
				Kind:      relatedRes.GetKind(),
				Name:      relatedRes.GetName(),
				Namespace: relatedRes.GetNamespace(),
			}
			resourceGraph.AddVertex(relatedResourceNode)
			if relationshipType == ParentTypeKey {
				resourceGraph.AddEdge(relatedResourceNode.GetHash(), objResourceNode.GetHash())
			} else {
				resourceGraph.AddEdge(objResourceNode.GetHash(), relatedResourceNode.GetHash())
			}
			relatedGVKOnGraph, _ := FindNodeOnGraph(relationshipGraph, relatedGVK.Group, relatedGVK.Version, relatedGVK.Kind)
			if relationshipType == ParentTypeKey && len(relatedGVKOnGraph.Parent) > 0 {
				// repeat for parent resources
				for _, parentRelation := range relatedGVKOnGraph.Parent {
					resourceGraph, _ = GetParents(ctx, client, relatedRes, parentRelation, relatedRes.GetNamespace(), relatedRes.GetName(), relatedResourceNode, relationshipGraph, resourceGraph)
				}
			} else if relationshipType == ChildTypeKey && len(relatedGVKOnGraph.Children) > 0 {
				// repeat for child resources
				for _, childRelation := range relatedGVKOnGraph.Children {
					resourceGraph, _ = GetChildren(ctx, client, relatedRes, childRelation, relatedRes.GetNamespace(), relatedRes.GetName(), relatedResourceNode, relationshipGraph, resourceGraph)
				}
			}
		}
	}
	return resourceGraph, nil
}

// GetByLabelSelector retrieves related resources based on label selectors from a given list of unstructured resources.
func GetByLabelSelector(
	ctx context.Context,
	relatedResList *unstructured.UnstructuredList,
	relationshipType string,
	client *dynamic.DynamicClient,
	obj unstructured.Unstructured,
	relation *Relationship,
	relatedGVK schema.GroupVersionKind,
	objResourceNode ResourceGraphNode,
	relationshipGraph graph.Graph[string, RelationshipGraphNode],
	resourceGraph graph.Graph[string, ResourceGraphNode],
) (graph.Graph[string, ResourceGraphNode], error) {
	log := ctxutil.GetLogger(ctx)

	log.Info("Using label selectors to find related resources...")
	var labelsMatch bool
	var err error
	for _, relatedRes := range relatedResList.Items {
		if relationshipType == ParentTypeKey {
			labelsMatch, err = topologyutil.LabelSelectorsMatch(relatedRes, obj, relation.SelectorPath)
		} else {
			labelsMatch, err = topologyutil.LabelSelectorsMatch(obj, relatedRes, relation.SelectorPath)
		}
		if labelsMatch && err == nil {
			log.Info("Resource found based on selector path.", "relationshipType", relationshipType, "kind", obj.GetKind(), "name", obj.GetName(), "selectorPath", relation.SelectorPath)
			log.Info("Resource is:", "relationshipType", relationshipType, "kind", relatedRes.GetKind(), "name", relatedRes.GetName())
			log.Info("---------------------------------------------------------------------------")
			rgv, _ := schema.ParseGroupVersion(relatedRes.GetAPIVersion())
			relatedResourceNode := ResourceGraphNode{
				Group:     rgv.Group,
				Version:   rgv.Version,
				Kind:      relatedRes.GetKind(),
				Name:      relatedRes.GetName(),
				Namespace: relatedRes.GetNamespace(),
			}
			resourceGraph.AddVertex(relatedResourceNode)
			if relationshipType == ParentTypeKey {
				resourceGraph.AddEdge(relatedResourceNode.GetHash(), objResourceNode.GetHash())
			} else {
				resourceGraph.AddEdge(objResourceNode.GetHash(), relatedResourceNode.GetHash())
			}
			relatedGVKOnGraph, _ := FindNodeOnGraph(relationshipGraph, relatedGVK.Group, relatedGVK.Version, relatedGVK.Kind)
			if relationshipType == ParentTypeKey && len(relatedGVKOnGraph.Parent) > 0 {
				for _, parentRelation := range relatedGVKOnGraph.Parent {
					resourceGraph, _ = GetParents(ctx, client, relatedRes, parentRelation, relatedRes.GetNamespace(), relatedRes.GetName(), relatedResourceNode, relationshipGraph, resourceGraph)
				}
			} else if relationshipType == ChildTypeKey && len(relatedGVKOnGraph.Children) > 0 {
				for _, childRelation := range relatedGVKOnGraph.Children {
					resourceGraph, _ = GetChildren(ctx, client, relatedRes, childRelation, relatedRes.GetNamespace(), relatedRes.GetName(), relatedResourceNode, relationshipGraph, resourceGraph)
				}
			}
		}
	}
	return resourceGraph, nil
}
