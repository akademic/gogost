// GoGOST -- Pure Go GOST cryptographic functions library
// Copyright (C) 2015-2024 Sergey Matveev <stargrave@stargrave.org>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, version 3 of the License.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package gost28147

import (
	"bytes"
	"crypto/cipher"
	"testing"
	"testing/quick"
)

func TestCFBCryptomanager(t *testing.T) {
	key := []byte{
		0x75, 0x71, 0x31, 0x34, 0xB6, 0x0F, 0xEC, 0x45,
		0xA6, 0x07, 0xBB, 0x83, 0xAA, 0x37, 0x46, 0xAF,
		0x4F, 0xF9, 0x9D, 0xA6, 0xD1, 0xB5, 0x3B, 0x5B,
		0x1B, 0x40, 0x2A, 0x1B, 0xAA, 0x03, 0x0D, 0x1B,
	}
	sbox := &SboxIdGostR341194TestParamSet
	pt := []byte{
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88,
		0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0x80, 0x00, 0x00,
	}
	ct := []byte{
		0x6E, 0xE8, 0x45, 0x86, 0xDD, 0x2B, 0xCA, 0x0C,
		0xAD, 0x36, 0x16, 0x94, 0x0E, 0x16, 0x42, 0x42,
	}
	c := NewCipher(key, sbox)
	iv := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	tmp := make([]byte, 16)
	fe := c.NewCFBEncrypter(iv)
	fe.XORKeyStream(tmp, pt)
	if !bytes.Equal(tmp, ct) {
		t.Fatal("encryption failed")
	}
	fd := c.NewCFBDecrypter(iv)
	fd.XORKeyStream(tmp, ct)
	if !bytes.Equal(tmp, pt) {
		t.Fatal("decryption failed")
	}
}

func TestCFBRandom(t *testing.T) {
	f := func(key [KeySize]byte, iv [BlockSize]byte, pt []byte) bool {
		if len(pt) == 0 {
			return true
		}
		c := NewCipher(key[:], SboxDefault)
		ct := make([]byte, len(pt))
		fe := c.NewCFBEncrypter(iv[:])
		fe.XORKeyStream(ct, pt)
		fd := c.NewCFBDecrypter(iv[:])
		pt2 := make([]byte, len(ct))
		fd.XORKeyStream(pt2, ct)
		return bytes.Equal(pt2, pt)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestCFBInterface(t *testing.T) {
	var key [32]byte
	var iv [8]byte
	c := NewCipher(key[:], SboxDefault)
	var _ cipher.Stream = c.NewCFBEncrypter(iv[:])
	var _ cipher.Stream = c.NewCFBDecrypter(iv[:])
}
