package communication

import (
	"steel-lang/datastructure"
)

type resourceRegistry map[string]datastructure.StringSet

func makeResourceRegistry(localRes datastructure.StringSet, localNodeName string) resourceRegistry {
	res := make(map[string]datastructure.StringSet)
	res[localNodeName] = localRes
	return res
}

func (r resourceRegistry) inventory() datastructure.StringSet {
	res := datastructure.MakeStringSet("")
	for nodeName := range r {
		res.Insert(nodeName)
	}
	return res
}
