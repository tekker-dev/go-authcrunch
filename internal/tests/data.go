// Copyright 2024 Paul Greenberg greenpau@outlook.com
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

package tests

import (
	"encoding/json"
)

// UnpackDict unpacks interface into a map.
func UnpackDict(i interface{}) (map[string]interface{}, error) {
	var m map[string]interface{}
	switch v := i.(type) {
	case string:
		if err := json.Unmarshal([]byte(v), &m); err != nil {
			return nil, err
		}
	default:
		b, err := json.Marshal(i)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(b, &m); err != nil {
			return nil, err
		}
	}
	return m, nil
}
