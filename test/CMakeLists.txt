cmake_minimum_required(VERSION 3.9)
project(${CMAKE_PROJECT_NAME}-test)

set(THREADS_PREFER_PTHREAD_FLAG ON)

find_package(Threads REQUIRED)

add_executable(libra test_libra.cpp)

include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../common/include)

if ( NOT crypto )
        message( WARNING "crypto NOT FOUND")
else()
        message( STATUS "crypto library : ${crypto}")
endif( NOT crypto )

target_link_libraries(libra PRIVATE Threads::Threads common crypto)

#install(TARGETS libra DESTINATION bin)
