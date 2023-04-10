//go:build !no_cv_highgui

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

// #include "stdccv/cv_gui.h"
import "C"
import (
	"fmt"
	"github.com/lynnplus/gotypes/geom"
	"runtime"
	"sync/atomic"
)

type WindowFlag int

const (
	WindowNormal      WindowFlag = 0x00000000 //!< the user can resize the window (no constraint) / also use to switch a fullscreen window to a normal size.
	WindowAutoSize    WindowFlag = 0x00000001 //!< the user cannot resize the window, the size is constrainted by the image displayed.
	WindowOpenGL      WindowFlag = 0x00001000 //!< window with opengl support.
	WindowFullScreen  WindowFlag = 1          //!< change the window to fullscreen.
	WindowFreeRatio   WindowFlag = 0x00000100 //!< the image expends as much as it can (no ratio constraint).
	WindowKeepRatio   WindowFlag = 0x00000000 //!< the ratio of the image is respected.
	WindowGuiExpanded WindowFlag = 0x00000000 //!< status bar and tool bar
	WindowGuiNormal   WindowFlag = 0x00000010 //!< old fashious way
)

type WindowPropertyFlag int

const (
	WinPropFullScreen  WindowPropertyFlag = iota //!< fullscreen property    (can be WINDOW_NORMAL or WINDOW_FULLSCREEN).
	WinPropAutoSize                              //!< autosize property      (can be WINDOW_NORMAL or WINDOW_AUTOSIZE).
	WinPropAspectRatio                           //!< window's aspect ration (can be set to WINDOW_FREERATIO or WINDOW_KEEPRATIO).
	WinPropOpenGL                                //!< opengl support.
	WinPropVisible                               //!< checks whether the window exists and is visible
	WinPropTopmost                               //!< property to toggle normal window being topmost or not
	WinPropVsync                                 //!< enable or disable VSYNC (in OpenGL mode)
)

type Window struct {
	p C.CWindow
}

var windowIdCounter atomic.Uint32

func NewWindow() *Window {
	return NewWindowWithFlags(WindowAutoSize)
}

func NewWindowWithFlags(flags WindowFlag) *Window {
	runtime.LockOSThread()
	name := fmt.Sprintf("win-%v", windowIdCounter.Add(1))

	p := C.Window_new(toCString(name), C.int(flags))
	win := &Window{p: p}
	autoRelease(win)
	return win
}

func (win *Window) SetTitle(title string) *Window {
	C.Window_set_title(win.p, toCString(title))
	return win
}

func (win *Window) SetProperty(propId WindowPropertyFlag, propVal WindowFlag) *Window {
	C.Window_set_property(win.p, C.int(propId), C.double(propVal))
	return win
}

func (win *Window) Resize(size geom.Size[int]) *Window {
	C.Window_resize(win.p, C.int(size.Width), C.int(size.Height))
	return win
}

func (win *Window) MoveTo(x, y int) *Window {
	C.Window_move(win.p, C.int(x), C.int(y))
	return win
}

// Update Force window to redraw its context
func (win *Window) Update() *Window {
	C.Window_update(win.p)
	return win
}

func (win *Window) ShowImage(img *Mat) *Window {
	C.Window_imshow(win.p, img.p)
	return win
}

func (win *Window) WaitKey(delayMs int) int {
	return int(C.Window_wait_key(C.int(delayMs)))
}

func (win *Window) ShowImageWithWaitKey(img *Mat, delayMs int) int {
	return int(C.Window_imshow_wait(win.p, img.p, C.int(delayMs)))
}

func (win *Window) Release() {
	resetFinalizer(win)
	C.Window_delete(win.p)
	if win.p != nil {
		runtime.UnlockOSThread()
	}
	win.p = nil
}
