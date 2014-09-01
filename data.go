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

package leap

// A Frame contains a snapshot of the data detected by the device.
type Frame struct {
	CurrentFrameRate float64        `json:"currentFrameRate"`
	ID               float64        `json:"id"`
	R                [][]float64    `json:"r"`
	S                float64        `json:"s"`
	T                []float64      `json:"t"`
	Timestamp        int64          `json:"timestamp"`
	Gestures         []Gesture      `json:"gestures"`
	Hands            []Hand         `json:"hands"`
	InteractionBox   InteractionBox `json:"interactionBox"`
	Pointables       []Pointable    `json:"pointables"`
}

type GestureState string

const (
	Start  GestureState = "start"
	Update GestureState = "update"
	Stop   GestureState = "stop"
)

type GestureType string

const (
	Circle    GestureType = "circle"
	Swipe     GestureType = "swipe"
	KeyTap    GestureType = "keyTap"
	ScreenTap GestureType = "screenType"
)

type Gesture struct {
	Center        []float64    `json:"center"`    // circle only
	Direction     []float64    `json:"direction"` // swipe, keyTap, screenTap only
	Duration      int64        `json:"duration"`
	HandsIDs      []int        `json:"handIds"`
	ID            int          `json:"id"`
	Normal        []float64    `json:"normal"` // circle only
	PointableIDs  []int        `json:"pointableIds"`
	Position      []float64    `json:"position"`      // swipe, keyTap, screenTap only
	Progress      float64      `json:"progress"`      // circle, keyTap, screenTap only
	Radius        float64      `json:"radius"`        // circle only
	Speed         float64      `json:"speed"`         // swipe only
	StartPosition []float64    `json:"startPosition"` // swipe only
	State         GestureState `json:"state"`
	Type          GestureType  `json:"type"`
}

type Hand struct {
	Direction              []float64   `json:"direction"`
	ID                     int         `json:"id"`
	PalmNormal             []float64   `json:"palmNormal"`
	PalmPosition           []float64   `json:"palmPosition"`
	PalmVelocity           []float64   `json:"palmVelocity"`
	R                      [][]float64 `json:"r"`
	S                      float64     `json:"s"`
	SphereCenter           []float64   `json:"sphereCenter"`
	SphereRadius           float64     `json:"sphereRadius"`
	StabilizedPalmPosition []float64   `json:"stabilizedPalmPosition"`
	T                      []float64   `json:"t"`
	TimeVisible            float64     `json:"timeVisible"`
}

type InteractionBox struct {
	Center []float64 `json:"center"`
	Size   []float64 `json:"size"`
}

type TouchZone string

const (
	None     TouchZone = "none"
	Hovering TouchZone = "hovering"
	Touching TouchZone = "touching"
)

type Pointable struct {
	Direction             []float64 `json:"direction"`
	HandID                int       `json:"handId"`
	ID                    int       `json:"id`
	Length                float64   `json:"length"`
	StabilizedTipPosition []float64 `json:"stabilizedTipPosition"`
	TimeVisible           float64   `json:"timeVisible"`
	TipPosition           []float64 `json:"tipPosition"`
	TipVelocity           []float64 `json:"tipVelocity"`
	Tool                  bool      `json:"tool"`
	TouchDistance         float64   `json:"touchDistance"`
	TouchZone             TouchZone `json:"touchZone"`
}
