// package client implements the functionality for the client application.
package client

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"time"
)

type msgtunnel struct {
	conn    net.Conn
	minrate time.Duration
}

func (t *msgtunnel) Marshal() ([]byte, error) {

}

type message struct {
	data []byte
}

func (msg *message) UnmarshalBinary(data []byte) error {
	if len(msg.data) <= 0 {
		return errors.New("client: error unmarshaling binary: msg contains no data to unmarshal")
	}

	var buff = bytes.NewBuffer(msg.data)
	var b = make([]byte, buff.Len(), buff.Cap())
	err := binary.Read(buff, binary.LittleEndian, b)
	if err != nil {
		return err
	}

	var key [32]byte
	c, err := aes.NewCipher(key[:])
	if err != nil {
		return err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return err
	}
	nonce := make([]byte, gcm.NonceSize())
	nonce = b[len(b)-gcm.NonceSize():]
	cleartxt, err := gcm.Open(b, nonce, buff.Bytes(), nil)
	if err != nil {
		return err
	}
	data = cleartxt[:]
	return nil

}

func (msg *message) MarshalBinary() (data []byte, err error) {
	if len(msg.data) <= 0 {
		return nil, errors.New("client: encountered error marshaling binary: msg contains no data to marshal")
	}

	var key [32]byte
	b, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(b)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, err
	}

	// TODO: seal data
	buff := new(bytes.Buffer)
	err = binary.Write(buff, binary.LittleEndian, append(gcm.Seal(data, nonce, msg.data, nil), nonce...))
	if err != nil {
		return nil, err
	}
	msg.data = buff.Bytes()
	return msg.data, nil
}
