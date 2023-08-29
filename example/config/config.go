/*******************************************************************************
 * Copyright 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package config

import (
	"encoding/json"
)

var (
	gConfig *Config
)

type Config struct {
}

func Initialize(s string) {
	gConfig = &Config{}
	if err := json.Unmarshal([]byte(s), &gConfig); err != nil {
		panic(err.Error())
	}
}

func GetConfig() *Config {
	return gConfig
}
