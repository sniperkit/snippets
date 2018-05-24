#!/bin/bash
echo "-- C compilers available"
ls -1 /usr/bin/gcc*
ls -1 /usr/bin/clang*
ls -1 /usr/bin/scan-build*
echo "----"

if [ "$TRAVIS_OS_NAME" == "linux" ]; then
	sudo apt-get update -qq || true
fi

echo "=== Building Debug"
mkdir -p build/Debug && cd build/Debug && cmake -DCMAKE_BUILD_TYPE=Debug -DCMAKE_INSTALL_PREFIX=$PWD/_install ../../ && make && make test && cd -

echo "=== Building Release"
mkdir -p build/Release && cd build/Release && cmake -DCMAKE_BUILD_TYPE=Release -DCMAKE_INSTALL_PREFIX=$PWD/_install ../../ && make && cd -
