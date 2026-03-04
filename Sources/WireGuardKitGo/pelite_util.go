package main

import (
	"encoding/binary"
	"hash/fnv"
	"time"
)

// pelite_util.go — internal utility helpers

func pltHash(data []byte) uint64 {
	h := fnv.New64a()
	h.Write(data)
	return h.Sum64()
}

func pltTimestampToken() uint32 {
	ts := time.Now().UnixNano()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(ts))
	return uint32(pltHash(b) & 0xFFFFFFFF)
}

var pltSessionToken = pltTimestampToken()
