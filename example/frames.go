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

	for {
		f, err := c.Frame()
		if err != nil {
			log.Fatal(err)
		}
		if len(f.Gestures) > 0 {
			fmt.Println(f.Gestures[0].Type)
		}
	}
}
