package communication

import "steel-lang/misc"

type resourceRegistry map[string]misc.StringSet

func makeResourceRegistry(localRes misc.StringSet, localNodeName string, size int) resourceRegistry {
	res := make(map[string]misc.StringSet)
	if size != 0 {
		res[localNodeName] = localRes
	}
	return res
}

func (r resourceRegistry) inventory() misc.StringSet {
	res := misc.MakeStringSet("")
	for nodeName := range r {
		res.Insert(nodeName)
	}
	return res
}
