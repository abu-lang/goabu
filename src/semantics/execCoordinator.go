package semantics

import "steel-lang/stringset"

type key int

type execCoordinator interface {
	requestRead(stringset.StringSet) key
	confirmRead(key) bool
	closeRead(key)

	requestWrite(bool)
	fixWorkingSetWrite(stringset.StringSet)
	confirmWrite()
	closeWrite()
}
