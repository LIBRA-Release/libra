cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

set(LIBRA_SRC
        src/libra.cpp)

set(LIBRA_HDR
        include)
        
add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})

#add_library(crypto SHARED IMPORTED)

target_link_libraries(common Threads::Threads crypto)

#install(TARGETS crypto DESTINATION lib)
install(TARGETS common DESTINATION lib)
