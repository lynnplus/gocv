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


#ifndef STDCCV_CV_GUI_H
#define STDCCV_CV_GUI_H

#ifdef __cplusplus

#include <string>

typedef struct CVWindow {
    const std::string name;

    explicit CVWindow(std::string &name) : name(name) {};
} *CWindow;

    #include <opencv2/opencv_modules.hpp>

    #ifdef HAVE_OPENCV_HIGHGUI

        #include <opencv2/highgui.hpp>

    #else
//#error missing opencv highgui library
    #endif
#else

typedef void* CWindow;
#endif


#ifdef __cplusplus
extern "C" {
#endif

#include "c_helper.h"
#include "cv_mat.h"

CWindow Window_new(CCString name, int flag);
void Window_delete(CWindow win);
void Window_set_property(CWindow win, int prop_id, double prop_value);
void Window_set_title(CWindow win, CCString title);
void Window_resize(CWindow win, int width, int height);
void Window_move(CWindow win, int x, int y);
void Window_update(CWindow win);
void Window_imshow(CWindow win, CMat mat);
int Window_wait_key(int delay);
int Window_imshow_wait(CWindow win, CMat mat, int delay);


#ifdef __cplusplus
}
#endif


#endif //STDCCV_CV_GUI_H
