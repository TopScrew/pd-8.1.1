// Copyright 2016 TiKV Project Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package typeutil

import "time"

// ZeroTime is a zero time.
var ZeroTime = time.Time{}

// ParseTimestamp returns a timestamp for a given byte slice.
func ParseTimestamp(data []byte) (time.Time, error) {
	nano, err := BytesToUint64(data)
	if err != nil {
		return ZeroTime, err
	}

	return time.Unix(0, int64(nano)), nil
}

// SubRealTimeByWallClock returns the duration between two different time.Time structs.
// You should use it to compare the real-world system time.
// And DO NOT USE IT TO COMPARE two TSOs' physical times directly in some cases.
func SubRealTimeByWallClock(after, before time.Time) time.Duration {
	return time.Duration(after.UnixNano() - before.UnixNano())
}

// SubTSOPhysicalByWallClock returns the duration between two different TSOs' physical times with millisecond precision.
func SubTSOPhysicalByWallClock(after, before time.Time) int64 {
	return after.UnixNano()/int64(time.Millisecond) - before.UnixNano()/int64(time.Millisecond)
}
