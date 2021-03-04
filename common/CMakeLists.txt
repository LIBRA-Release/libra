cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

find_library(crytolib crypto ./../lib/linux)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)
add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})


add_library(crytolib SHARED IMPORTED GLOBAL)
        
target_link_libraries(common crytolib)

install(TARGETS common DESTINATION lib)
