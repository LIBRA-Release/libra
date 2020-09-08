# Project Libra
# author: apolo.yasuda@ge.com 
#
cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-golib)
find_library(golang REQUIRED)

set(TARGET shim_go)
set(GOPATH ${CMAKE_CURRENT_BINARY_DIR})

set(SRC crypto.go)
set(LIB crypto.so)

add_custom_command(OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/${LIB}
  DEPENDS ${SRCS}
  WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
  COMMAND env GOPATH=${GOPATH} go build -buildmode=c-archive
  -o "${CMAKE_CURRENT_BINARY_DIR}/${LIB}"
  ${CMAKE_GO_FLAGS} ./...
  COMMENT "Building Go library")

add_custom_target(${TARGET} DEPENDS ${LIB} ${HEADER})

add_library(goshim STATIC IMPORTED GLOBAL)
add_dependencies(goshim ${TARGET})
set_target_properties(goshim
  PROPERTIES
  IMPORTED_LOCATION ${CMAKE_CURRENT_BINARY_DIR}/${LIB}
  INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR})
