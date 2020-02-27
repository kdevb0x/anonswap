package client

import (
	"errors"
	"io/ioutil"

	"github.com/alexmullins/zip"
	"github.com/awnumar/memguard"
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

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

}
