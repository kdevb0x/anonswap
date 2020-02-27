
// package zip implements AES encryption and decryption for zip files according
// to the specification found at: https://www.winzip.com/win/en/aes_info.html
package zip

import (
	z "archive/zip"
	"encoding/binary"
	"os"

	"golang.org/x/sys/cpu"
)

const bufferSize = 32_768

type ae1 z.Compressor


func PlaformByteOrder() binary.ByteOrder {
	cpu.
}
