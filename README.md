# gocv
golang pkg simple bind for opencv4

## Build

### Default
package uses cgo to compile the stdccv static library and link

### Dynamic link stdccv

set system env(or use golang build params)
```
CGO_CPPFLAGS=-I{stdccv include dir};
CGO_LDFLAGS=-LD{stdccv dynamic lib dir} -l{stdccv lib name}
```

use CMakeLists.txt (project root dir) compile dynamic lib,go build add 

### Ignore some opencv module
Some modules of opencv are included by default,can be ignored using the build tag

example:ignore opencv_highgui
`go build -tags=no_cv_highgui`

Supported tags and corresponding opencv modules

| Tag             | cv_module | default |
|-----------------|-----------|---------|
| no_cv_highgui   | highgui   | false   |
| no_cv_imgcodecs | imgcodecs | false   |
| no_cv_imgproc   | imgproc   | false   |

### Link opencv lib

#### use precompiled libraries

1. set `CGO_CPPFLAGS` and `CGO_LDFLAGS` link opencv include and lib dir
2. add build tags `custom_cv` for go build

#### MacOS

install opencv

`brew install opencv`

install pkgconfig

`brew install pkgconfig`



