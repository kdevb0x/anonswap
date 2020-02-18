// Copyright 2019 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

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
