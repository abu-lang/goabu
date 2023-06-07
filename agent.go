// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

package goabu

type Agent interface {
	Start() error
	Join() error
	ForAll([]byte) error
	ReceivedActions() (<-chan chan []byte, <-chan chan string)
	Stop() error
	IsRunning() bool
	SetLogLevel(int)
}
