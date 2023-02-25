// Copyright 2021 github.com/gagliardetto
// This file has been modified by github.com/gagliardetto
//
// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package solana

import (
	"time"
)

// UnixTimeSeconds represents a UNIX second-resolution timestamp.
type UnixTimeSeconds int64

func (res UnixTimeSeconds) Time() time.Time {
	return time.Unix(int64(res), 0)
}

func (res UnixTimeSeconds) String() string {
	return res.Time().String()
}

// UnixTimeMilliseconds represents a UNIX millisecond-resolution timestamp.
type UnixTimeMilliseconds int64

func (res UnixTimeMilliseconds) Time() time.Time {
	return time.Unix(0, int64(res)*int64(time.Millisecond))
}

func (res UnixTimeMilliseconds) String() string {
	return res.Time().String()
}

// DurationSeconds represents a duration in seconds.
type DurationSeconds int64

func (res DurationSeconds) Duration() time.Duration {
	return time.Duration(res) * time.Second
}

func (res DurationSeconds) String() string {
	return res.Duration().String()
}

// DurationMilliseconds represents a duration in milliseconds.
type DurationMilliseconds int64

func (res DurationMilliseconds) Duration() time.Duration {
	return time.Duration(res) * time.Millisecond
}

func (res DurationMilliseconds) String() string {
	return res.Duration().String()
}
