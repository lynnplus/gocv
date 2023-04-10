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


#include "cv_mat.h"
#include "c_helper.h"
#include <opencv2/core/mat.hpp>

CMat Mat_new() {
    return new cv::Mat();
}

CMat Mat_new_from_size(int rows, int cols, int _type) {
    return new cv::Mat(rows, cols, _type);
}

CMat Mat_new_from_data(int rows, int cols, int _type, CByteArray buf) {
    return new cv::Mat(rows, cols, _type, buf.data);
}

void Mat_delete(CMat mat) {
    delete mat;
}

bool Mat_empty(CMat mat) {
    return mat->empty();
}

size_t Mat_total(CMat mat) {
    return mat->total();
}

CMat Mat_clone(CMat mat) {
    return new cv::Mat(mat->clone());
}

CByteArray Mat_data(CMat mat) {
    return CByteArray{
            reinterpret_cast<uint8_t *>(mat->data),
            mat->total() * mat->elemSize()
    };
}

void Mat_rows_cols(CMat mat, int row_col[2]) {
    row_col[0] = mat->rows;
    row_col[1] = mat->cols;
}
