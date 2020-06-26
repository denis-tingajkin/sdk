// Copyright (c) 2020 Doc.ai and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package expire

import "time"

type configurable interface {
	setPeriod(time.Duration)
}

// Option is expire registry configuration option
type Option interface {
	apply(configurable)
}

type applierFunc func(configurable)

func (f applierFunc) apply(c configurable) {
	f(c)
}

// WithPeriod sets specific period to checking expiration
func WithPeriod(duration time.Duration) Option {
	return applierFunc(func(c configurable) {
		c.setPeriod(duration)
	})
}
