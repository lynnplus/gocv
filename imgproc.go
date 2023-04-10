//go:build !no_cv_imgproc

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

//#include "stdccv/cv_imgproc.h"
import "C"

type LineType int

const (
	LineFilled LineType = -1
	Lint4      LineType = 4  //!< 4-connected line
	Line8      LineType = 8  //!< 8-connected line
	LineAA     LineType = 16 //!< antialiased line
)

type HersheyFont int

const (
	FontHersheySimplex       HersheyFont = iota //!< normal size sans-serif font
	FontHersheyPlain                            //!< small size sans-serif font
	FontHersheyDuplex                           //!< normal size sans-serif font (more complex than FONT_HERSHEY_SIMPLEX)
	FontHersheyComplex                          //!< normal size serif font
	FontHersheyTriplex                          //!< normal size serif font (more complex than FONT_HERSHEY_COMPLEX)
	FontHersheyComplexSmall                     //!< smaller version of FONT_HERSHEY_COMPLEX
	FontHersheyScriptSimplex                    //!< hand-writing style font
	FontHersheyScriptComplex                    //!< more complex variant of FONT_HERSHEY_SCRIPT_SIMPLEX
	FontItalic               HersheyFont = 16   //!< flag for italic font
)

type MarkerType int

const (
	MarkerCross        MarkerType = iota //!< A crosshair marker shape
	MarkerTiltedCross                    //!< A 45 degree tilted crosshair marker shape
	MarkerStar                           //!< A star marker shape, combination of cross and tilted cross
	MarkerDiamond                        //!< A diamond marker shape
	MarkerSquare                         //!< A square marker shape
	MarkerTriangleUp                     //!< An upwards pointing triangle marker shape
	MarkerTriangleDown                   //!< A downwards pointing triangle marker shape
)

type drawsParams[C any] struct {
	color     Color
	thickness int
	lineType  LineType
	sub       C
}

type DrawsParams struct {
	drawsParams[*DrawsParams]
}

//= drawsParams[*drawsParams[any]]

func NewDrawsParams() *DrawsParams {
	p := &DrawsParams{
		drawsParams: defaultDrawsParams[*DrawsParams](),
	}
	p.init(p)
	return p
}

func (d *drawsParams[C]) init(sub C) {
	d.sub = sub
}

func (d *drawsParams[C]) Thickness(val int) C {
	d.thickness = val
	return d.sub
}

func (d *drawsParams[C]) LineType(line LineType) C {
	d.lineType = line
	return d.sub
}

func (d *drawsParams[C]) Color(val Color) C {
	d.color = val
	return d.sub
}

type MarkerParams struct {
	drawsParams[*MarkerParams]
	markerType MarkerType
	markerSize int
}

func NewMarkerParams() *MarkerParams {
	return defaultMarkerParams()
}

func (mp *MarkerParams) MarkerType(mt MarkerType) *MarkerParams {
	mp.markerType = mt
	return mp
}

func (mp *MarkerParams) MarkerSize(size int) *MarkerParams {
	mp.markerSize = size
	return mp
}

type ArrowedLineParams struct {
	drawsParams[*ArrowedLineParams]
	tipLength float64
}

func NewArrLineParams() *ArrowedLineParams {
	return defaultArrowedLineParams()
}

func (lp *ArrowedLineParams) TipLength(v float64) *ArrowedLineParams {
	lp.tipLength = v
	return lp
}

type TextParams struct {
	drawsParams[*TextParams]
	fontFace         HersheyFont
	fontScale        float64
	bottomLeftOrigin bool
}

func NewTextParams() *TextParams {
	return defaultTextParams()
}

func (tp *TextParams) FontFace(font HersheyFont) *TextParams {
	tp.fontFace = font
	return tp
}

func (tp *TextParams) FontScale(scale float64) *TextParams {
	tp.fontScale = scale
	return tp
}

func (tp *TextParams) BottomLeftOrigin(enable bool) *TextParams {
	tp.bottomLeftOrigin = enable
	return tp
}

func defaultDrawsParams[T any]() drawsParams[T] {
	return drawsParams[T]{color: Black, thickness: 1, lineType: Line8}
}

func defaultTextParams() *TextParams {
	tp := &TextParams{
		drawsParams: defaultDrawsParams[*TextParams](),
		fontFace:    FontHersheySimplex,
		fontScale:   1,
	}
	tp.init(tp)
	return tp
}

func defaultArrowedLineParams() *ArrowedLineParams {
	lp := &ArrowedLineParams{
		drawsParams: defaultDrawsParams[*ArrowedLineParams](),
		tipLength:   0.1,
	}
	lp.init(lp)
	return lp
}

func defaultMarkerParams() *MarkerParams {
	mp := &MarkerParams{
		drawsParams: defaultDrawsParams[*MarkerParams](),
		markerType:  MarkerCross,
		markerSize:  20,
	}
	mp.init(mp)
	return mp
}

// PutText renders the specified text string in the image.
// param origin is Bottom-left corner of the text string in the image.
func PutText(mat *Mat, text string, origin Point, params *TextParams) {
	if params == nil {
		params = defaultTextParams()
	}
	C.cv_put_text(mat.p, toCString(text), toCIntPoint(origin), C.int(params.fontFace),
		C.double(params.fontScale), toCScalarFromColor(params.color), C.int(params.thickness),
		C.int(params.lineType), C.bool(params.bottomLeftOrigin))
}

func GetTextSize(text string, params *TextParams) (size Size, baseLine int) {
	if params == nil {
		params = defaultTextParams()
	}
	cbs := C.int(0)
	sz := C.cv_get_text_size(toCString(text), C.int(params.fontFace),
		C.double(params.fontScale), C.int(params.thickness), &cbs)
	return Size{Width: int(sz.width), Height: int(sz.height)}, int(cbs)
}

func GetFontScaleFromHeight(fontFace, pixelHeight, thickness int) float64 {
	fs := C.cv_get_font_scale_from_height(C.int(fontFace), C.int(pixelHeight), C.int(thickness))
	return float64(fs)
}

func PutRectangle(mat *Mat, rect Rectangle, params *DrawsParams) {
	if params == nil {
		params = NewDrawsParams()
	}
	C.cv_rectangle(mat.p, toCIntRect(rect), toCScalarFromColor(params.color), C.int(params.thickness),
		C.int(params.lineType))
}

func PutCircle(mat *Mat, center Point, radius int, params *DrawsParams) {
	if params == nil {
		params = NewDrawsParams()
	}

	C.cv_circle(mat.p, toCIntPoint(center), C.int(radius), toCScalarFromColor(params.color), C.int(params.thickness),
		C.int(params.lineType))
}

func PutEllipse(mat *Mat, center Point, axes Size, angle, startAngle, endAngle float64, params *DrawsParams) {
	if params == nil {
		params = NewDrawsParams()
	}
	C.cv_ellipse(mat.p, toCIntPoint(center), toCIntSize(axes),
		C.double(angle), C.double(startAngle), C.double(endAngle),
		toCScalarFromColor(params.color), C.int(params.thickness),
		C.int(params.lineType))
}

func PutMarker(mat *Mat, position Point, params *MarkerParams) {
	if params == nil {
		params = defaultMarkerParams()
	}
	C.cv_draw_marker(mat.p, toCIntPoint(position),
		toCScalarFromColor(params.color), C.int(params.markerType),
		C.int(params.markerSize), C.int(params.thickness), C.int(params.lineType))
}

func PutLine(mat *Mat, firstPt Point, secondPt Point, params *DrawsParams) {
	if params == nil {
		params = NewDrawsParams()
	}
	C.cv_line(mat.p, toCIntPoint(firstPt), toCIntPoint(secondPt),
		toCScalarFromColor(params.color), C.int(params.thickness),
		C.int(params.lineType))
}

func PutArrowedLine(mat *Mat, firstPt Point, secondPt Point, params *ArrowedLineParams) {
	if params == nil {
		params = defaultArrowedLineParams()
	}
	C.cv_arrowed_line(mat.p, toCIntPoint(firstPt), toCIntPoint(secondPt), toCScalarFromColor(params.color),
		C.int(params.thickness),
		C.int(params.lineType), C.double(params.tipLength))
}
