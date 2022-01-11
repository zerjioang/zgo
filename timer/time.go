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

package timer

import (
	"time"

	"github.com/zerjioang/time32"
)

// Now returns current system time as int64 epoch value
func Now() int64 {
	return time32.ReuseUnix()
}

func Time() time.Time {
	return time32.ReuseTime()
}
