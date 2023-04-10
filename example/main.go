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

package main

import (
	"fmt"
	cv "github.com/lynnplus/gocv"
)

func main() {
	img := cv.IMRead("example/bin/face_person.jpg")
	if img == nil {
		panic("image empty")
	}
	fmt.Println(img.Size())

	win := cv.NewWindow()

	pam := cv.NewTextParams().Thickness(3)

	text := "text"
	size, _ := cv.GetTextSize(text, pam)

	cv.PutText(img, "text", cv.Point{X: 10, Y: size.Height}, pam)

	win.ShowImageWithWaitKey(img, -1)

	params := map[cv.IMWriteFlag]int{
		cv.IMWriteJpegQuality: 10,
	}
	cv.IMWriteWithParams("example/bin/test.jpeg", img, params)
}
