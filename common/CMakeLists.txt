cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_LIB
        lib/linux)       

add_library(common STATIC
        ${LIBRA_LIB} ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})
      
install(TARGETS common DESTINATION lib)
