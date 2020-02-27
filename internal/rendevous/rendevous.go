// Copyright 2020 kdevb0x ltd

// package rendevous implements the rendevous point for ephemeral data transfers
package rendevous

import (
	"log"
	"net"

	"golang.org/x/net/proxy"

	"github.com/awnumar/memguard"
)

func init() {
	log.SetPrefix("<rendevous-server>")

	// redundant?
	// log.SetOutput(os.Stderr)
}

type buffer struct {
	mem    *memguard.Enclave
	buff   *memguard.LockedBuffer
	stream *memguard.Stream
}

// XferServer is an ephemeral proxy server used for a single transfer session.
// They are designed to be quick to initialize and teardown.
type XferServer struct {
	Addr      string
	Buff      *buffer
	ChunkSize int
}

// NewXferServer initializes a new XferServer including allocating and zeroing
// the secure enclave.
func NewXferServer(addr string) *XferServer {
	net.SplitHostPort()
	phost := proxy.NewPerHost(proxy.Direct, nil)
	proxy.SOCKS5
	var x = new(XferServer)
	enc := memguard.NewEnclaveRandom(x.ChunkSize)
	lb, err := enc.Open()
	if err != nil {
		log.Fatal(err)
	}
	lb.Wipe()
	x.Buff = &buffer{mem: enc}
	x.Buff.buff = lb
	x.Buff.stream = memguard.NewStream()

	return x

}
