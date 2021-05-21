package communication

import (
	"steel-lang/datastructure"
)

type resourceRegistry map[string]datastructure.StringSet

func makeResourceRegistry(localRes datastructure.StringSet, localNodeName string, size int) resourceRegistry {
	res := make(map[string]datastructure.StringSet)
	if size != 0 {
		res[localNodeName] = localRes
	}
	return res
}

func (r resourceRegistry) inventory() datastructure.StringSet {
	res := datastructure.MakeStringSet("")
	for nodeName := range r {
		res.Insert(nodeName)
	}
	return res
}
