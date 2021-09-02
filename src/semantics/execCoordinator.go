package semantics

import "steel-lang/stringset"

type key int

type execCoordinator interface {
	requestRead(stringset.Set) key
	confirmRead(key) bool
	closeRead(key)

	requestWrite(bool)
	fixWorkingSetWrite(stringset.Set)
	confirmWrite()
	closeWrite()
}
