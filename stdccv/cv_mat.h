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


#ifndef STDCCV_CV_MAT_H
#define STDCCV_CV_MAT_H

#ifdef __cplusplus

#include <opencv2/core.hpp>

typedef cv::Mat *CMat;

#else
typedef void* CMat;
#endif


#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>
#include "c_helper.h"

CMat Mat_new();
CMat Mat_new_from_size(int rows, int cols, int _type);
CMat Mat_new_from_data(int rows, int cols, int _type, CByteArray buf);
void Mat_delete(CMat mat);
bool Mat_empty(CMat mat);
CMat Mat_clone(CMat mat);
size_t Mat_total(CMat mat);
CByteArray Mat_data(CMat mat);
void Mat_rows_cols(CMat mat, int row_col[2]);

#ifdef __cplusplus
}
#endif

#endif //STDCCV_CV_MAT_H
