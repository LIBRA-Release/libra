cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-common)

add_library(common-lib SHARED IMPORTED GLOBAL)
set_target_properties(
  common-lib
  PROPERTIES
  IMPORTED_LOCATION "./../lib/linux/crypto.so"
)

set(LIBRA_SRC
        src/libra.cpp)

add_library(common STATIC
        ${LIBRA_SRC} ${CMAKE_THREAD_LIBS_INIT})
       
install(TARGETS common DESTINATION lib)
