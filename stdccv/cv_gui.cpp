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

#include "cv_gui.h"
#include "c_helper.h"
#include <opencv2/highgui.hpp>

#ifdef HAVE_OPENCV_HIGHGUI

CWindow Window_new(const CCString name, int flag) {
    auto str = cstr2str(name);
    auto win = new CVWindow(str);
    cv::namedWindow(win->name, flag);
    return win;
}

void Window_delete(CWindow win) {
    if (win) {
        cv::destroyWindow(win->name);
    }
    delete win;
}

void Window_set_property(CWindow win, int prop_id, double prop_value) {
    cv::setWindowProperty(win->name, prop_id, prop_value);
}

void Window_set_title(CWindow win, const CCString title) {
    auto str = cstr2str(title);
    cv::setWindowTitle(win->name, str);
}

void Window_resize(CWindow win, int width, int height) {
    cv::resizeWindow(win->name, width, height);
}

void Window_move(CWindow win, int x, int y) {
    cv::moveWindow(win->name, x, y);
}

void Window_update(CWindow win) {
    cv::updateWindow(win->name);
}

void Window_imshow(CWindow win, CMat mat) {
    if (mat->empty()) {
        return;
    }
    cv::imshow(win->name, *mat);
}

int Window_wait_key(int delay) {
    return cv::waitKey(delay);
}

int Window_imshow_wait(CWindow win, CMat mat, int delay) {
    Window_imshow(win, mat);
    return Window_wait_key(delay);
}

#endif