cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

set(LIBRA_SRC
        src/libra.cpp)

link_directories(${CMAKE_SOURCE_DIR}/lib/linux)

add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})
      
install(TARGETS common DESTINATION lib)
