cmake_minimum_required(VERSION 3.20)
#std c wrapper for opencv
project(stdccv)

set(CMAKE_C_STANDARD 17)
set(CMAKE_CXX_STANDARD 20)

FILE(GLOB_RECURSE sourceFiles ${PROJECT_SOURCE_DIR}/stdccv/*.cpp)
FILE(GLOB_RECURSE headerFiles ${PROJECT_SOURCE_DIR}/stdccv/*.h)

find_package(OpenCV 4.5.0 REQUIRED)

add_library(${PROJECT_NAME} SHARED ${sourceFiles})
#add_library(${PROJECT_NAME} SHARED ${sourceFiles} ${headerFiles})
target_include_directories(${PROJECT_NAME} PUBLIC ${headerFiles})
target_include_directories(${PROJECT_NAME} PRIVATE ${OpenCV_INCLUDE_DIRS})
target_link_libraries(${PROJECT_NAME} ${OpenCV_LIBS})

