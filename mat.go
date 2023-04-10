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

//#include "stdccv/cv_mat.h"
import "C"
import (
	"github.com/lynnplus/gotypes/geom"
	"unsafe"
)

const (
	cvChannelShift = 3
	cvDepthMax     = 1 << cvChannelShift
	cvMatDepthMask = cvDepthMax - 1
)

const (
	cvDepth8U = iota
	cvDepth8S
	cvDepth16U
	cvDepth16S
	cvDepth32S
	cvDepth32F
	cvDepth64F
	cvDepth16F
)

type MatType int

const (
	MatType8UC1 MatType = (cvDepth8U & cvMatDepthMask) + (((iota + 1) - 1) << cvChannelShift)
	MatType8UC2
	MatType8UC3
	MatType8UC4
	MatType8SC1 MatType = (cvDepth8S & cvMatDepthMask) + (((iota - 4 + 1) - 1) << cvChannelShift)
	MatType8SC2
	MatType8SC3
	MatType8SC4
	MatType16UC1 MatType = (cvDepth16U & cvMatDepthMask) + (((iota - 8 + 1) - 1) << cvChannelShift)
	MatType16UC2
	MatType16UC3
	MatType16UC4
	MatType16SC1 MatType = (cvDepth16S & cvMatDepthMask) + (((iota - 12 + 1) - 1) << cvChannelShift)
	MatType16SC2
	MatType16SC3
	MatType16SC4
	MatType32SC1 MatType = (cvDepth32S & cvMatDepthMask) + (((iota - 16 + 1) - 1) << cvChannelShift)
	MatType32SC2
	MatType32SC3
	MatType32SC4
	MatType32FC1 MatType = (cvDepth32F & cvMatDepthMask) + (((iota - 20 + 1) - 1) << cvChannelShift)
	MatType32FC2
	MatType32FC3
	MatType32FC4
	MatType64FC1 MatType = (cvDepth64F & cvMatDepthMask) + (((iota - 24 + 1) - 1) << cvChannelShift)
	MatType64FC2
	MatType64FC3
	MatType64FC4
	MatType16FC1 MatType = (cvDepth16F & cvMatDepthMask) + (((iota - 28 + 1) - 1) << cvChannelShift)
	MatType16FC2
	MatType16FC3
	MatType16FC4
)

func MatType8U(channel int) MatType {
	return makeMatType(cvDepth8U, channel)
}

func MatType8S(channel int) MatType {
	return makeMatType(cvDepth8S, channel)
}

func MatType16U(channel int) MatType {
	return makeMatType(cvDepth16U, channel)
}

func MatType16S(channel int) MatType {
	return makeMatType(cvDepth16S, channel)
}

func MatType32S(channel int) MatType {
	return makeMatType(cvDepth32S, channel)
}

func MatType32F(channel int) MatType {
	return makeMatType(cvDepth32F, channel)
}

func makeMatType(depth, cn int) MatType {
	return MatType((depth & cvMatDepthMask) + (((cn + 1) - 1) << cvChannelShift))
}

type Mat struct {
	p         C.CMat
	cacheData []byte
}

func NewMat() *Mat {
	p := C.Mat_new()
	return newMat(p, nil)
}

func NewMatFromSize(rows, cols int, mt MatType) *Mat {
	p := C.Mat_new_from_size(C.int(rows), C.int(cols), C.int(mt))
	return newMat(p, nil)
}

func NewMatFromData(rows, cols int, mt MatType, data []byte) *Mat {
	p := C.Mat_new_from_data(C.int(rows), C.int(cols), C.int(mt), toCByteArray(data))
	return newMat(p, data)
}

func newMat(p C.CMat, data []byte) *Mat {
	mat := &Mat{p: p, cacheData: data}
	autoRelease(mat)
	return mat
}

func (mat *Mat) UnsafePtr() unsafe.Pointer {
	return unsafe.Pointer(mat.p)
}

func (mat *Mat) Empty() bool {
	return bool(C.Mat_empty(mat.p))
}

func (mat *Mat) Clone() *Mat {
	p := C.Mat_clone(mat.p)
	return newMat(p, mat.cacheData)
}

func (mat *Mat) Total() int {
	return int(C.Mat_total(mat.p))
}

func (mat *Mat) Data() []byte {
	cbs := C.Mat_data(mat.p)
	return C.GoBytes(unsafe.Pointer(cbs.data), C.int(cbs.length))
}

func (mat *Mat) Size() geom.Size[int] {
	rc := [2]C.int{}
	C.Mat_rows_cols(mat.p, &rc[0])
	return geom.Size[int]{
		Width:  int(rc[1]),
		Height: int(rc[0]),
	}
}

// Release Manually release the memory held by mat
func (mat *Mat) Release() {
	resetFinalizer(mat)
	C.Mat_delete(mat.p)
	mat.p = nil
	mat.cacheData = nil
}

func toCString(str string) C.CCString {
	p := unsafe.Pointer(unsafe.StringData(str))
	return C.CCString{
		data:   (*C.char)(p),
		length: C.size_t(len(str)),
	}
}

func toCByteArray(data []byte) C.CByteArray {
	return C.CByteArray{
		data:   (*C.uint8_t)(unsafe.Pointer(&data[0])),
		length: C.size_t(len(data)),
	}
}

func toCIntPoint(pt Point) C.CIntPoint {
	return C.CIntPoint{
		x: C.int(pt.X),
		y: C.int(pt.Y),
	}
}

func toCIntSize(sz Size) C.CIntSize {
	return C.CIntSize{
		width:  C.int(sz.Width),
		height: C.int(sz.Height),
	}
}

func toCScalarFromColor(color Color) C.CScalar {
	return C.CScalar{
		v0: C.double(color.B),
		v1: C.double(color.G),
		v2: C.double(color.R),
		v3: C.double(color.A),
	}
}

func toCIntRect(rect Rectangle) C.CIntRect {
	sz := rect.Size()
	return C.CIntRect{
		x:      C.int(rect.Min.X),
		y:      C.int(rect.Min.Y),
		width:  C.int(sz.Width),
		height: C.int(sz.Height),
	}
}
