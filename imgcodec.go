//go:build !no_cv_imgcodecs

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

//#include "stdccv/cv_imgcodecs.h"
import "C"
import "unsafe"

type IMReadFlag int

const (
	IMReadUnChanged         IMReadFlag = -1
	IMReadGrayScale         IMReadFlag = 0
	IMReadColor             IMReadFlag = 1
	IMReadAnyDepth          IMReadFlag = 2
	IMReadAnyColor          IMReadFlag = 4
	IMReadLoadGDAL          IMReadFlag = 8
	IMReadReducedGrayscale2 IMReadFlag = 16
	IMReadReducedColor2     IMReadFlag = 17
	IMReadReducedGaryScale4 IMReadFlag = 32
	IMReadReducedColor4     IMReadFlag = 33
	IMReadReducedGaryScale8 IMReadFlag = 64
	IMReadReducedColor8     IMReadFlag = 65
	IMReadIgnoreOrientation IMReadFlag = 128
)

type IMWriteFlag int

const (
	IMWriteJpegQuality              IMWriteFlag = 1
	IMWriteJpegProgressive          IMWriteFlag = 2
	IMWriteJpegOptimize             IMWriteFlag = 3
	IMWriteJpegRstInterval          IMWriteFlag = 4
	IMWriteJpegLumaQuality          IMWriteFlag = 5
	IMWriteJpegChromaQuality        IMWriteFlag = 6
	IMWritePngCompression           IMWriteFlag = 16
	IMWritePngStrategy              IMWriteFlag = 17
	IMWritePngBiLevel               IMWriteFlag = 18
	IMWritePxmBinary                IMWriteFlag = 32
	IMWriteExrType                  IMWriteFlag = 48
	IMWriteExrCompression           IMWriteFlag = 49
	IMWriteWebpQuality              IMWriteFlag = 64
	IMWritePamTupleType             IMWriteFlag = 128
	IMWriteTiffResUnit              IMWriteFlag = 256
	IMWriteTiffXDpi                 IMWriteFlag = 257
	IMWriteTiffYDpi                 IMWriteFlag = 258
	IMWriteTiffCompression          IMWriteFlag = 259
	IMWriteJpeg2000CompressionX1000 IMWriteFlag = 272
)

func IMRead(file string) *Mat {
	return IMReadWithFlags(file, IMReadColor)
}

func IMReadWithFlags(file string, flags IMReadFlag) *Mat {
	str := toCString(file)
	p := C.cv_image_read(str, C.int(flags))
	if p == nil {
		return nil
	}
	return newMat(p, nil)
}

func IMWrite(file string, mat *Mat) bool {
	return IMWriteWithParams(file, mat, nil)
}

func IMWriteWithParams(file string, mat *Mat, params map[IMWriteFlag]int) bool {
	str := toCString(file)
	ps := C.CIntArray{}
	if len(params) > 0 {
		temp := make([]C.int, 0, len(params)*2)
		for flag, v := range params {
			temp = append(temp, C.int(flag), C.int(v))
		}
		ps.data = (*C.int)(unsafe.Pointer(&temp[0]))
		ps.length = C.size_t(len(temp))
	}
	return bool(C.cv_image_write(str, mat.p, ps))
}
