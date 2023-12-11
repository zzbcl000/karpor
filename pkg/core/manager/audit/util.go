// Copyright The Karbour Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package audit

import (
	"math"
)

// P is the number of issues, and S is the sum of the severity (range 1-5) of
// the issue S will not be less than P.
//
// Example:
// - When there is one high-level issue, P=1 and S=3.
// - When there are three high-level issues, P=3 and S=9.
// - When there are ten low-level issues, P=10 and S=10.
func CalculateScore(p, s int) float64 {
	a, b := -0.04, -0.06
	param := a*float64(p) + b*float64(s)
	return 100 * math.Exp(param)
}
