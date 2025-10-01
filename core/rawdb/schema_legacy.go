// Copyright 2024 The go-ethereum Authors
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

package rawdb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/ethdb"
)

var (
	// The fields below are database keys used by the legacy offline pruner,
	// which is superseded by the online pruner. These fields are reserved
	// to avoid database key conflicts.
	offSetOfCurrentAncientFreezer = []byte("-offSetOfCurrentAncientFreezer")
	offSetOfLastAncientFreezer    = []byte("-offSetOfLastAncientFreezer")
	frozenOfAncientDBKey          = []byte("-frozenOfAncientDB")
	pruneAncientKey               = []byte("-pruneAncient")
)

// ReadLegacyOffset reads the legacy pruner metadata from the database and
// returns the highest offset found.
func ReadLegacyOffset(db ethdb.KeyValueReader) uint64 {
	var (
		maxOffset uint64
		data      []byte
	)
	data, _ = db.Get(offSetOfCurrentAncientFreezer)
	if len(data) > 0 {
		offset := new(big.Int).SetBytes(data).Uint64()
		if offset > maxOffset {
			maxOffset = offset
		}
	}
	data, _ = db.Get(offSetOfLastAncientFreezer)
	if len(data) > 0 {
		offset := new(big.Int).SetBytes(data).Uint64()
		if offset > maxOffset {
			maxOffset = offset
		}
	}
	data, _ = db.Get(frozenOfAncientDBKey)
	if len(data) > 0 {
		offset := new(big.Int).SetBytes(data).Uint64()
		if offset > maxOffset {
			maxOffset = offset
		}
	}
	return maxOffset
}

// CleanLegacyOffset removes the legacy pruner metadata from the database.
func CleanLegacyOffset(db ethdb.KeyValueWriter) {
	db.Delete(offSetOfCurrentAncientFreezer)
	db.Delete(offSetOfLastAncientFreezer)
	db.Delete(frozenOfAncientDBKey)
	db.Delete(pruneAncientKey)
}
