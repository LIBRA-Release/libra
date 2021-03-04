cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

set(LIBRA_SRC
        src/libra.cpp)

add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})
        
set(CMAKE_INSTALL_RPATH_USE_LINK_PATH TRUE)

install(TARGETS common DESTINATION lib)
