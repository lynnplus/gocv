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

#include "cv_imgproc.h"
#include "cv_mat.h"
#include <opencv2/core/matx.hpp>
#include <opencv2/core/types.hpp>
#include <opencv2/imgproc.hpp>


#ifdef HAVE_OPENCV_IMGPROC


void cv_put_text(CMat img, CCString text, CIntPoint org,
                 int fontFace, double fontScale, CScalar color,
                 int thickness, int lineType, bool bottomLeftOrigin) {
    auto str = cstr2str(text);
    cv::putText(
            *img,
            str,
            cv::Point(org.x, org.y),
            fontFace,
            fontScale,
            cv::Scalar(color.v0, color.v1, color.v2, color.v3),
            thickness, lineType, bottomLeftOrigin);
}


double cv_get_font_scale_from_height(int fontFace, int pixelHeight, int thickness) {
    return cv::getFontScaleFromHeight(fontFace, pixelHeight, thickness);
}

void cv_rectangle(CMat img, CIntRect rec, CScalar color, int thickness, int lineType) {
    cv::Rect rect(rec.x, rec.y, rec.width, rec.height);
    cv::Scalar scalar(color.v0, color.v1, color.v2, color.v3);
    cv::rectangle(*img, rect, scalar, thickness, lineType);
}

void cv_circle(CMat img, CIntPoint center, int radius,
               CScalar color, int thickness, int lineType) {
    cv::Scalar scalar(color.v0, color.v1, color.v2, color.v3);
    cv::circle(*img, cv::Point(center.x, center.y), radius, scalar, thickness, lineType);
}

void cv_ellipse(CMat img, CIntPoint center, CIntSize axes,
                double angle, double startAngle, double endAngle,
                CScalar color, int thickness, int lineType) {
    cv::ellipse(*img,
                cv::Point(center.x, center.y),
                cv::Size(axes.width, axes.height),
                angle, startAngle, endAngle,
                cv::Scalar(color.v0, color.v1, color.v2, color.v3),
                thickness, lineType);
}

void cv_draw_marker(CMat img, CIntPoint position, CScalar color,
                    int markerType, int markerSize, int thickness, int line_type) {
    cv::drawMarker(*img,
                   cv::Point(position.x, position.y),
                   cv::Scalar(color.v0, color.v1, color.v2, color.v3),
                   markerType, markerSize, thickness, line_type);
}

void cv_line(CMat img, CIntPoint pt1, CIntPoint pt2, CScalar color,
             int thickness, int lineType) {
    cv::line(*img,
             cv::Point(pt1.x, pt1.y),
             cv::Point(pt2.x, pt2.y),
             cv::Scalar(color.v0, color.v1, color.v2, color.v3),
             thickness, lineType);
}

void cv_arrowed_line(CMat img, CIntPoint pt1, CIntPoint pt2, CScalar color,
                     int thickness, int lineType, double tipLength) {
    cv::arrowedLine(*img,
                    cv::Point(pt1.x, pt1.y),
                    cv::Point(pt2.x, pt2.y),
                    cv::Scalar(color.v0, color.v1, color.v2, color.v3),
                    thickness, lineType, 0, tipLength);
}

CIntSize cv_get_text_size(CCString text, int fontFace,
                          double fontScale, int thickness, int *baseLine) {
    auto str = cstr2str(text);
    auto size = cv::getTextSize(str, fontFace, fontScale, thickness, baseLine);
    return {
            .width=size.width,
            .height=size.height
    };
}

#endif
