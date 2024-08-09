// Copyright 2018 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package rawdb contains a collection of low level database accessors.

package rawdb

import (
	"github.com/ethereum/go-ethereum/common"
)

const WasmSchemaVersion byte = 0x00

var (
	wasmSchemaVersionKey = []byte("WasmSchemaVersion")

	// 0x00 prefix to avoid conflicts when wasmdb is not separate database
	activatedAsmWavmPrefix = []byte{0x00, 'w', 'w'} // (prefix, moduleHash) -> stylus module (wavm)
	activatedAsmArmPrefix  = []byte{0x00, 'w', 'r'} // (prefix, moduleHash) -> stylus asm for ARM system
	activatedAsmX86Prefix  = []byte{0x00, 'w', 'x'} // (prefix, moduleHash) -> stylus asm for x86 system
	activatedAsmHostPrefix = []byte{0x00, 'w', 'h'} // (prefix, moduleHash) -> stylus asm for system other then ARM and x86
)

// WasmKeyLen = CompiledWasmCodePrefix + moduleHash
const WasmKeyLen = 3 + common.HashLength

type WasmKey = [WasmKeyLen]byte

// key = prefix + moduleHash
func activatedKey(prefix []byte, moduleHash common.Hash) WasmKey {
	var key WasmKey
	copy(key[:3], prefix)
	copy(key[3:], moduleHash[:])
	return key
}
