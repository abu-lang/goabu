// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package memory implements GoAbU logical resources.
package memory

// ResourceController is the interface modeling an AbU node's state and its interaction with the environment.
type ResourceController interface {
	// Start shall be called as soon as the node is ready to process inputs from the environment.
	Start() error
	// Inputs returns a channel providing the inputs received from the environment as strings of the form "<resource_name> = <value>;".
	Inputs() <-chan string
	// Errors returns a channel handing the errors that occurs during operation.
	Errors() <-chan error
	// Modified shall be called when the resource with the given identifier is set to a different value.
	Modified(string)
	// Extract returns a shallow copy of only the resources specified by the provided identifiers.
	Extract([]string) Resources
	// Enclose adds the provided resources to the ResourceController, overwriting previous values if present.
	Enclose(Resources)
	// HasDuplicates verifies if the ResourceController has multiple resources sharing the same identifier.
	HasDuplicates() bool
	// Has checks if the ResourceController contains a resource identified by the provided string.
	Has(string) bool
	// Types returns a map with an entry for each resource specifying its type (one of the following: "Bool", "Integer", "Float", "Text", "Time", "Other").
	// Prerequisite: !HasDuplicates()
	Types() map[string]string
	// GetResources provides access to the resources.
	GetResources() Resources
	// ResourceNames returns the list of all the managed resources' identifiers (without repeated elements).
	ResourceNames() []string
	// String returns a string representation of the ResourceController for debugging purposes.
	String() string
	// Copy returns a shallow copy of the ResourceController.
	Copy() ResourceController
}
