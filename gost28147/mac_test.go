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
	"crypto/rand"
	"hash"
	"testing"
	"testing/quick"
)

func TestMACVectors(t *testing.T) {
	key := []byte("This is message\xFF length\x0032 bytes")
	c := NewCipher(key, SboxDefault)
	var iv [8]byte
	m, err := c.NewMAC(8, iv[:])
	if err != nil {
		t.FailNow()
	}

	t.Run("a", func(t *testing.T) {
		m.Write([]byte("a"))
		if !bytes.Equal(m.Sum(nil), []byte{0xbd, 0x5d, 0x3b, 0x5b, 0x2b, 0x7b, 0x57, 0xaf}) {
			t.FailNow()
		}
	})

	t.Run("abc", func(t *testing.T) {
		m.Reset()
		m.Write([]byte("abc"))
		if !bytes.Equal(m.Sum(nil), []byte{0x28, 0x66, 0x1e, 0x40, 0x80, 0x5b, 0x1f, 0xf9}) {
			t.FailNow()
		}
	})

	t.Run("128U", func(t *testing.T) {
		m.Reset()
		for i := 0; i < 128; i++ {
			m.Write([]byte("U"))
		}
		if !bytes.Equal(m.Sum(nil), []byte{0x1a, 0x06, 0xd1, 0xba, 0xd7, 0x45, 0x80, 0xef}) {
			t.FailNow()
		}
	})

	t.Run("xxxxxxxxxxxxx", func(t *testing.T) {
		m.Reset()
		for i := 0; i < 13; i++ {
			m.Write([]byte("x"))
		}
		if !bytes.Equal(m.Sum(nil), []byte{0x91, 0x7e, 0xe1, 0xf1, 0xa6, 0x68, 0xfb, 0xd3}) {
			t.FailNow()
		}
	})
}

func TestMACRandom(t *testing.T) {
	var key [KeySize]byte
	rand.Read(key[:])
	c := NewCipher(key[:], SboxDefault)
	f := func(iv [BlockSize]byte, data []byte) bool {
		if len(data) == 0 {
			return true
		}
		m, err := c.NewMAC(8, iv[:])
		if err != nil {
			return false
		}

		var tag1 []byte
		var tag2 []byte

		for _, b := range data {
			m.Write([]byte{b})
		}
		m.Sum(tag1)

		m.Reset()
		m.Write(data)
		m.Sum(tag2)

		return bytes.Equal(tag1, tag2)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestMACInterface(t *testing.T) {
	var key [KeySize]byte
	var iv [8]byte
	c := NewCipher(key[:], SboxDefault)
	m, _ := c.NewMAC(8, iv[:])
	var _ hash.Hash = m
}

func BenchmarkMAC(b *testing.B) {
	key := make([]byte, KeySize)
	iv := make([]byte, BlockSize)
	rand.Read(key)
	rand.Read(iv)
	b1 := make([]byte, BlockSize)
	b2 := make([]byte, BlockSize)
	rand.Read(b1)
	rand.Read(b2)
	c := NewCipher(key[:], SboxDefault)
	mac, _ := c.NewMAC(BlockSize, iv)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mac.Write(b1)
		mac.Write(b2)
		mac.Sum(nil)
	}
}
