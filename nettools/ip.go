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

package nettools

import (
	"log"
	"net"
)

// IP address lengths (bytes).
const (
	IPv4len = 4
	// Bigger than we need, not too big to worry about overflow
	big = 0xFFFFFF
)

var (
	outboundIP string
)

func init() {
	// This will be executed
	// only once in the entire lifetime of the program
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	outboundIP = localAddr.IP.String()
}

// GetOutboundIP Get preferred outbound ip of this machine
func GetOutboundIP() string {
	return outboundIP
}

// IsIpv4 returns true if provided IP string is a valid IPv4
func IsIpv4(s string) bool {
	for i := 0; i < IPv4len; i++ {
		if len(s) == 0 {
			// Missing octets.
			return false
		}
		if i > 0 {
			if s[0] != '.' {
				return false
			}
			s = s[1:]
		}
		var n int
		var i int
		for i = 0; i < len(s) && '0' <= s[i] && s[i] <= '9'; i++ {
			n = n*10 + int(s[i]-'0')
			if n >= big {
				n = big
			}
		}
		if i == 0 {
			n = 0
			i = 0
		}
		if n > 0xFF {
			return false
		}
		s = s[i:]
	}
	return len(s) == 0
}
