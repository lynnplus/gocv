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

//
//

#ifndef STDCCV_CV_IMGCODECS_H
#define STDCCV_CV_IMGCODECS_H

#ifdef __cplusplus

#include <opencv2/opencv_modules.hpp>

#ifdef HAVE_OPENCV_IMGCODECS

    #include <opencv2/imgcodecs.hpp>

#else
//#error missing opencv imgcodecs library
#endif

extern "C" {
#endif

#include "c_helper.h"
#include "cv_core.h"

CMat cv_image_read(CCString filename, int flags);
bool cv_image_write(CCString filename, CMat img, CIntArray params);

#ifdef __cplusplus
}
#endif

#endif //STDCCV_CV_IMGCODECS_H
