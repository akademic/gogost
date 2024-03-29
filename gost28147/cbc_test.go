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

func TestCBCCrypter(t *testing.T) {
	f := func(key [KeySize]byte, iv [BlockSize]byte, pt []byte) bool {
		c := NewCipher(key[:], SboxDefault)
		for i := 0; i < BlockSize; i++ {
			pt = append(pt, pt...)
		}
		ct := make([]byte, len(pt))
		e := cipher.NewCBCEncrypter(c, iv[:])
		e.CryptBlocks(ct, pt)
		d := cipher.NewCBCDecrypter(c, iv[:])
		d.CryptBlocks(ct, ct)
		return bytes.Equal(pt, ct)
	}
	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
