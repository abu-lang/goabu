// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package physical

import (
	"github.com/abu-lang/goabu/memory"
)

type IOdelegate interface {
	Start(IOadaptor, chan<- string, chan<- error) error
	Modified(IOadaptor, string, memory.Resources, chan<- error) *memory.Resources
}
