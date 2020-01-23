// Package hubserver provides functionality akin to a DNS server for clients to
// find content hosts.
package hubserver

import (
	"time"

	"github.com/kdevb0x/anonswap/internal/common"
)

type hostRecord struct {
	// addr is the address of the host providing this record.
	addr string

	// hosts public key
	pubkey []byte

	// shares are all the shares (content) the provider is sharing.
	shares []share

	// timeToDie is the point in time after which this Record is no longer
	// valid.
	timeToDie time.Time
}

func (h *hostRecord) Addr() string {
	return h.addr
}

func (h *hostRecord) PubKey() []byte {
	return h.pubkey
}

func (h *hostRecord) AllShares() []common.Share {
	var s []common.Share = make([]common.Share, len(h.shares), cap(h.shares))
	for i := 0; i < len(s); i++ {
		var c common.Share = &h.shares[i]
		s[i] = c
	}
	return s
}

func (h *hostRecord) Valid() bool {
	if time.Now().Before(h.timeToDie) {
		return true
	}
	return false
}

type share struct {
	// the files name
	filename string

	// the files sha512 hash
	filehash []byte

	// total size in bytes of the file
	length int

	// offset into the swapPkg where this file begins.
	offset int
}

func (s *share) FileName() string {
	return s.filename
}

func (s *share) FileHash() []byte {
	return s.filehash
}

func (s *share) Size() int {
	return s.length
}
