/*
 * Copyright (c) 2023 Lynn <lynnplus90@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gocv

import (
	"github.com/lynnplus/gotypes/geom"
	"runtime"
)

type Releaser interface {
	Release()
}

func autoRelease(obj Releaser) {
	runtime.SetFinalizer(obj, func(o Releaser) {
		o.Release()
	})
}

func resetFinalizer(obj Releaser) {
	runtime.SetFinalizer(obj, nil)
}

type (
	Point     = geom.Point[int]
	Size      = geom.Size[int]
	Rectangle = geom.Rectangle[int]
)

type Color struct {
	R, G, B, A uint8
}

func RGBColor(r, g, b uint8) Color {
	return Color{r, g, b, 255}
}

func HexColor() Color {
	return Color{}
}

func HexStrColor(hex string) Color {
	return Color{}
}

var (
	Black  = Color{0, 0, 0, 255}
	White  = Color{255, 255, 255, 255}
	Red    = Color{255, 0, 0, 255}
	Green  = Color{0, 255, 0, 255}
	Blue   = Color{0, 0, 255, 255}
	Yellow = Color{255, 255, 0, 255}
)
