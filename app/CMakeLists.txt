# Set the minimum version of CMake that can be used
# To find the cmake version run
# $ cmake --version
cmake_minimum_required(VERSION 3.5)

set(CMAKE_SYSTEM_NAME "Linux")

set(CMAKE_SYSTEM_PROCESSOR "ppc")
set(CROSSCOMPILE /opt/powerpc-e500mc--glibc--stable-2020.08-1/bin/powerpc-buildroot-linux-gnu-)
#set(CMAKE_SYSTEM_PROCESSOR "arm")
#set(CROSSCOMPILE /opt/gcc-arm-none-eabi-10-2020-q4-major/bin/arm-none-eabi-)

set(CMAKE_C_COMPILER ${CROSSCOMPILE}gcc)
set(CMAKE_CXX_COMPILER ${CROSSCOMPILE}g++)
set(AS_EXECUTABLE ${CROSSCOMPILE}as)


set(CONFIF_RUNTIME_CPU_DETECT 0 CACHE STRING "")

# Set the project name
project (app)

# Add an executable
add_executable(app main.c)
