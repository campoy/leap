// Copyright 2014 Google Inc. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to writing, software distributed
// under the License is distributed on a "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"

	"github.com/campoy/leap"
)

func main() {
	c, err := leap.Connect("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	var f struct {
		Hands []struct{} `json:"hands"`
	}

	old := 0
	for {
		if err := c.Decode(&f); err != nil {
			log.Fatal(err)
		}
		if n := len(f.Hands); n != old {
			old = n
			fmt.Printf("I see %v hands\n", n)
		}
	}
}
