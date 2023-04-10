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


#ifndef STDCCV_C_HELPER_H
#define STDCCV_C_HELPER_H


#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>
#include <stdbool.h>

typedef struct {
    uint8_t *data;
    size_t length;
} CByteArray;

typedef struct {
    const char *data;
    size_t length;
} CCString;

typedef struct {
    const int *data;
    size_t length;
} CIntArray;

typedef struct {
    int x, y;
} CIntPoint;

typedef struct {
    double v0, v1, v2, v3;//map to color blue,green,red,alpha
} CScalar;

typedef struct {
    int width, height;
} CIntSize;

typedef struct {
    int x, y, width, height;
} CIntRect;


#ifdef __cplusplus
}
#endif


#ifdef __cplusplus

#include <string>

static inline std::string cstr2str(const CCString &src) {
    return std::string(src.data, src.length);
}

#endif

#endif //STDCCV_C_HELPER_H
