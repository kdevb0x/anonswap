package client

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/alexmullins/zip"
	"github.com/awnumar/memguard"
	hub "github.com/kdevb0x/anonswap/internal/hub-server"
)

type archive struct {
	mem      []byte // memory enclave
	f        *zip.Writer
	streamer *memguard.Stream
	created  bool
}

func NewArchive() *archive {
	enc := memguard.NewEnclaveRandom(1024 * 1000)
	buf, err := enc.Open()
	if err != nil {
		panic(err)
	}
	buf.Melt()
	buf.Wipe()
	var a = new(archive)
	a.mem = buf.Bytes()
	a.streamer = memguard.NewStream()
	return a
}

func (a archive) Create() error {
	a.f = zip.NewWriter(a.streamer)
	a.created = true
	return nil

}

func (a *archive) AddFile(path string) error {
	if !a.created {
		return errors.New("unable to add file to uninitialized archive; must create() archive first!")
	}

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	var share = new(hub.Share)
	share.SetFileName(f.Name())

	buff, err := a.f.Create(f.Name())
	if err != nil {
		return err
	}
	size, err := io.ReadFull(f, a.mem)
	if err != nil {
		return err
	}
	m, err := buff.Write(a.mem)

	size2, err := buff.Write(a.mem)
	if err != nil {
		return err
	}
	if size != size2 {
		switch {
		case size > size2:

			return fmt.Errorf("error short write: only wrote %d out of %d bytes", size2, size)
		}
	}

}
