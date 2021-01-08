package yqlib

import (
	"container/list"
	"fmt"

	"gopkg.in/yaml.v3"
)

type OperatorHandler func(d *dataTreeNavigator, matchingNodes *list.List, pathNode *PathTreeNode) (*list.List, error)

func UnwrapDoc(node *yaml.Node) *yaml.Node {
	if node.Kind == yaml.DocumentNode {
		return node.Content[0]
	}
	return node
}

func EmptyOperator(d *dataTreeNavigator, matchingNodes *list.List, pathNode *PathTreeNode) (*list.List, error) {
	return list.New(), nil
}

func createBooleanCandidate(owner *CandidateNode, value bool) *CandidateNode {
	valString := "true"
	if !value {
		valString = "false"
	}
	node := &yaml.Node{Kind: yaml.ScalarNode, Value: valString, Tag: "!!bool"}
	return &CandidateNode{Node: node, Document: owner.Document, Path: owner.Path}
}

func nodeToMap(candidate *CandidateNode) *list.List {
	elMap := list.New()
	elMap.PushBack(candidate)
	return elMap
}

func createTraversalTree(path []interface{}) *PathTreeNode {
	if len(path) == 0 {
		return &PathTreeNode{Operation: &Operation{OperationType: SelfReference}}
	} else if len(path) == 1 {
		return &PathTreeNode{Operation: &Operation{OperationType: TraversePath, Value: path[0], StringValue: fmt.Sprintf("%v", path[0])}}
	}
	return &PathTreeNode{
		Operation: &Operation{OperationType: ShortPipe},
		Lhs:       createTraversalTree(path[0:1]),
		Rhs:       createTraversalTree(path[1:])}
}
