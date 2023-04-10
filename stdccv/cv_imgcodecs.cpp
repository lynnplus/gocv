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

#include "cv_imgcodecs.h"
#include <vector>

#ifdef HAVE_OPENCV_IMGCODECS

CMat cv_image_read(CCString filename, int flags) {
    auto fn = cstr2str(filename);
    cv::Mat src = cv::imread(fn, flags);
    if (src.empty()) {
        return nullptr;
    }
    return new cv::Mat(src);
}

bool cv_image_write(CCString filename, CMat img, CIntArray params) {
    auto fn = cstr2str(filename);
    std::vector<int> paramList;
    if (params.length > 0) {
        const int *ptr = params.data;
        paramList.assign(ptr, ptr + params.length);
    }
    return cv::imwrite(fn, *img, paramList);
}

#endif
