// Copyright 2021 Massimo Comuzzo, Michele Pasqua and Marino Miculan
// SPDX-License-Identifier: Apache-2.0

// Package config contains GoAbU's global configuration.
package config

import "os"

var Production = os.Getenv("GOABU_DEBUG") != "1"
