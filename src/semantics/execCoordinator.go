package semantics

import "steel-lang/misc"

type key int

type execCoordinator interface {
	requestRead(misc.StringSet) key
	confirmRead(key) bool
	closeRead(key)

	requestWrite(bool)
	fixWorkingSetWrite(misc.StringSet)
	confirmWrite()
	closeWrite()
}
