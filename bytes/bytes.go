//
// Copyright zerjioang. 2021 All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package bytes

import (
	"bytes"
	"encoding/binary"
	"unsafe"
)

const hextable = "0123456789abcdef"

// RemoveDoubleQuotes removes double quotes and replace by single quotes
// in error messages
func RemoveDoubleQuotes(src []byte) []byte {
	for i := 0; i < len(src); i++ {
		if src[i] == '"' {
			src[i] = '\''
		}
	}
	return src
}

// BytesToHex converts byte
func BytesToHex(src []byte) string {
	if src == nil || len(src) == 0 {
		return ""
	}
	dst := make([]byte, len(src)*2)
	var j uint = 0
	_ = src[len(src)-1]
	for i := 0; i < len(src); i++ {
		v := src[i]
		dst[j] = hextable[v>>4]
		dst[j+1] = hextable[v&0x0f]
		j += 2
	}
	return string(dst)
}

var nativeEndian binary.ByteOrder

func init() {
	checkEndian()
}

func checkEndian() {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0x0001)
	endians := []binary.ByteOrder{binary.LittleEndian, binary.BigEndian}
	nativeEndian = endians[buf[1]]
}

/*
converts an int value to bytearray
*/

// deprecated
func uint32ToBytes(a int32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, nativeEndian, a)
	return buf.Bytes(), err
}

// Uint32ToBytes converts an int value to bytearray
func Uint32ToBytes(v int32) [4]byte {
	var b [4]byte
	_ = b[3] // early bounds check to guarantee safety of writes below
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	return b
}

// Endianess returns current system endianess
func Endianess() binary.ByteOrder {
	return nativeEndian
}
