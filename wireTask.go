// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package goabu

import (
	"bytes"
	"encoding/gob"

	"github.com/abu-lang/goabu/ecarule"
	"github.com/abu-lang/goabu/memory"
	"github.com/abu-lang/goabu/stringset"
)

// wireTasks groups a list of [ecarule.RemoteTask] along with a list of values for their remote resources.
type wireTasks struct {
	memory.Resources
	Tasks []ecarule.RemoteTask
}

// marshalWireTasks marshalls w allowing for network transfer.
func marshalWireTasks(w wireTasks) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(w)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

// unmarshalWireTasks performs the unmarshalling of a received wireTasks.
func unmarshalWireTasks(b []byte) (wireTasks, error) {
	var res wireTasks
	decoder := gob.NewDecoder(bytes.NewBuffer(b))
	err := decoder.Decode(&res)
	return res, err
}

// getRemoteResources returns the names of all the remote resources of the tasks in w.Tasks.
func (w wireTasks) getRemoteResources() []string {
	set := stringset.Make()
	for _, rTask := range w.Tasks {
		set.Add(stringset.Make(rTask.RemoteResources...))
	}
	return set.Slice()
}
