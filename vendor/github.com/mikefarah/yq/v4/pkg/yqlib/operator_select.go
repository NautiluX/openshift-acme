package yqlib

import (
	"container/list"
)

func SelectOperator(d *dataTreeNavigator, matchingNodes *list.List, pathNode *PathTreeNode) (*list.List, error) {

	log.Debugf("-- selectOperation")
	var results = list.New()

	for el := matchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)

		rhs, err := d.GetMatchingNodes(nodeToMap(candidate), pathNode.Rhs)

		if err != nil {
			return nil, err
		}

		// grab the first value
		first := rhs.Front()

		if first != nil {
			result := first.Value.(*CandidateNode)
			includeResult, errDecoding := isTruthy(result)
			if errDecoding != nil {
				return nil, errDecoding
			}

			if includeResult {
				results.PushBack(candidate)
			}
		}
	}
	return results, nil
}
