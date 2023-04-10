// Copyright (c) 2023 Lynn <lynnplus90@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.


#ifndef STDCCV_CV_IMGPROC_H
#define STDCCV_CV_IMGPROC_H

#ifdef __cplusplus

#include <opencv2/opencv_modules.hpp>

#ifdef HAVE_OPENCV_IMGPROC

    #include <opencv2/imgproc.hpp>

#else
//#error missing opencv imgcodecs library
#endif

extern "C" {
#endif

#include "c_helper.h"
#include "cv_core.h"

void cv_put_text(CMat img, CCString text, CIntPoint org, int fontFace, double fontScale,
                 CScalar color, int thickness, int lineType, bool bottomLeftOrigin);

double cv_get_font_scale_from_height(int fontFace, int pixelHeight, int thickness);

CIntSize cv_get_text_size(CCString text, int fontFace, double fontScale, int thickness, int *baseLine);

void cv_rectangle(CMat img, CIntRect rec, CScalar color, int thickness, int lineType);
void cv_circle(CMat img, CIntPoint center, int radius, CScalar color, int thickness, int lineType);
void cv_ellipse(CMat img, CIntPoint center, CIntSize axes, double angle, double startAngle,
                double endAngle, CScalar color, int thickness, int lineType);
void cv_draw_marker(CMat img, CIntPoint position, CScalar color, int markerType, int markerSize,
                    int thickness, int line_type);
void cv_line(CMat img, CIntPoint pt1, CIntPoint pt2, CScalar color, int thickness, int lineType);
void cv_arrowed_line(CMat img, CIntPoint pt1, CIntPoint pt2, CScalar color, int thickness,
                     int lineType, double tipLength);

#ifdef __cplusplus
}
#endif

#endif //STDCCV_CV_IMGPROC_H
