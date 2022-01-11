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

package host

import "os"

var (
	hostname = ""
)

func init() {
	// hostname variables assumes that it will not change
	// at runtime
	// make 1 initial syscall to get the hostname value
	hostname, _ = os.Hostname()
}

// Name returns current hostname information
func Name() string {
	return hostname
}

// Reload thread safe hostname reloading function
func Reload() error {
	h, err := os.Hostname()
	if err == nil && h != "" {
		// TODO a read-write data race may occur here
		hostname = h
	}
	return err
}
