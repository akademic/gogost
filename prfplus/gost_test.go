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

package prfplus

import (
	"bytes"
	"testing"
)

func TestPRFIPsecPRFPlusGOSTR34112012256(t *testing.T) {
	prf := NewPRFIPsecPRFPlusGOSTR34112012256([]byte{
		0xC9, 0xA9, 0xA7, 0x73, 0x20, 0xE2, 0xCC, 0x55,
		0x9E, 0xD7, 0x2D, 0xCE, 0x6F, 0x47, 0xE2, 0x19,
		0x2C, 0xCE, 0xA9, 0x5F, 0xA6, 0x48, 0x67, 0x05,
		0x82, 0xC0, 0x54, 0xC0, 0xEF, 0x36, 0xC2, 0x21,
	})
	dst := make([]byte, 64)
	PRFPlus(prf, dst, []byte{
		0x01, 0x26, 0xBD, 0xB8, 0x78, 0x00, 0x1D, 0x80,
		0x60, 0x3C, 0x85, 0x44, 0xC7, 0x27, 0x01, 0x00,
	})
	if !bytes.Equal(dst, []byte{
		0x2D, 0xE5, 0xEE, 0x84, 0xE1, 0x3D, 0x7B, 0xE5,
		0x36, 0x16, 0x67, 0x39, 0x13, 0x37, 0x0A, 0xB0,
		0x54, 0xC0, 0x74, 0xB7, 0x9B, 0x69, 0xA8, 0xA8,
		0x46, 0x82, 0xA9, 0xF0, 0x4F, 0xEC, 0xD5, 0x87,
		0x29, 0xF6, 0x0D, 0xDA, 0x45, 0x7B, 0xF2, 0x19,
		0xAA, 0x2E, 0xF9, 0x5D, 0x7A, 0x59, 0xBE, 0x95,
		0x4D, 0xE0, 0x08, 0xF4, 0xA5, 0x0D, 0x50, 0x4D,
		0xBD, 0xB6, 0x90, 0xBE, 0x68, 0x06, 0x01, 0x53,
	}) {
		t.FailNow()
	}
}

func TestPRFIPsecPRFPlusGOSTR34112012512(t *testing.T) {
	prf := NewPRFIPsecPRFPlusGOSTR34112012512([]byte{
		0xC9, 0xA9, 0xA7, 0x73, 0x20, 0xE2, 0xCC, 0x55,
		0x9E, 0xD7, 0x2D, 0xCE, 0x6F, 0x47, 0xE2, 0x19,
		0x2C, 0xCE, 0xA9, 0x5F, 0xA6, 0x48, 0x67, 0x05,
		0x82, 0xC0, 0x54, 0xC0, 0xEF, 0x36, 0xC2, 0x21,
	})
	dst := make([]byte, 128)
	PRFPlus(prf, dst, []byte{
		0x01, 0x26, 0xBD, 0xB8, 0x78, 0x00, 0x1D, 0x80,
		0x60, 0x3C, 0x85, 0x44, 0xC7, 0x27, 0x01, 0x00,
	})
	if !bytes.Equal(dst, []byte{
		0x5D, 0xA6, 0x71, 0x43, 0xA5, 0xF1, 0x2A, 0x6D,
		0x6E, 0x47, 0x42, 0x59, 0x6F, 0x39, 0x24, 0x3F,
		0xCC, 0x61, 0x57, 0x45, 0x91, 0x5B, 0x32, 0x59,
		0x10, 0x06, 0xFF, 0x78, 0xA2, 0x08, 0x63, 0xD5,
		0xF8, 0x8E, 0x4A, 0xFC, 0x17, 0xFB, 0xBE, 0x70,
		0xB9, 0x50, 0x95, 0x73, 0xDB, 0x00, 0x5E, 0x96,
		0x26, 0x36, 0x98, 0x46, 0xCB, 0x86, 0x19, 0x99,
		0x71, 0x6C, 0x16, 0x5D, 0xD0, 0x6A, 0x15, 0x85,
		0x48, 0x34, 0x49, 0x5A, 0x43, 0x74, 0x6C, 0xB5,
		0x3F, 0x0A, 0xBA, 0x3B, 0xC4, 0x6E, 0xBC, 0xF8,
		0x77, 0x3C, 0xA6, 0x4A, 0xD3, 0x43, 0xC1, 0x22,
		0xEE, 0x2A, 0x57, 0x75, 0x57, 0x03, 0x81, 0x57,
		0xEE, 0x9C, 0x38, 0x8D, 0x96, 0xEF, 0x71, 0xD5,
		0x8B, 0xE5, 0xC1, 0xEF, 0xA1, 0xAF, 0xA9, 0x5E,
		0xBE, 0x83, 0xE3, 0x9D, 0x00, 0xE1, 0x9A, 0x5D,
		0x03, 0xDC, 0xD6, 0x0A, 0x01, 0xBC, 0xA8, 0xE3,
	}) {
		t.FailNow()
	}
}
