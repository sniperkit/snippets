# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.10

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/local/Cellar/cmake/3.10.2/bin/cmake

# The command to remove a file.
RM = /usr/local/Cellar/cmake/3.10.2/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug

# Include any dependencies generated for this target.
include CMakeFiles/libsyntax7-static.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/libsyntax7-static.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/libsyntax7-static.dir/flags.make

CMakeFiles/libsyntax7-static.dir/src/scanner.c.o: CMakeFiles/libsyntax7-static.dir/flags.make
CMakeFiles/libsyntax7-static.dir/src/scanner.c.o: ../../src/scanner.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/libsyntax7-static.dir/src/scanner.c.o"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/libsyntax7-static.dir/src/scanner.c.o   -c /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/src/scanner.c

CMakeFiles/libsyntax7-static.dir/src/scanner.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/libsyntax7-static.dir/src/scanner.c.i"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/src/scanner.c > CMakeFiles/libsyntax7-static.dir/src/scanner.c.i

CMakeFiles/libsyntax7-static.dir/src/scanner.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/libsyntax7-static.dir/src/scanner.c.s"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/src/scanner.c -o CMakeFiles/libsyntax7-static.dir/src/scanner.c.s

CMakeFiles/libsyntax7-static.dir/src/scanner.c.o.requires:

.PHONY : CMakeFiles/libsyntax7-static.dir/src/scanner.c.o.requires

CMakeFiles/libsyntax7-static.dir/src/scanner.c.o.provides: CMakeFiles/libsyntax7-static.dir/src/scanner.c.o.requires
	$(MAKE) -f CMakeFiles/libsyntax7-static.dir/build.make CMakeFiles/libsyntax7-static.dir/src/scanner.c.o.provides.build
.PHONY : CMakeFiles/libsyntax7-static.dir/src/scanner.c.o.provides

CMakeFiles/libsyntax7-static.dir/src/scanner.c.o.provides.build: CMakeFiles/libsyntax7-static.dir/src/scanner.c.o


CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o: CMakeFiles/libsyntax7-static.dir/flags.make
CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o: ../../src/dict/c_language.c
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building C object CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -o CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o   -c /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/src/dict/c_language.c

CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing C source to CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.i"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/src/dict/c_language.c > CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.i

CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling C source to assembly CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.s"
	/Library/Developer/CommandLineTools/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/src/dict/c_language.c -o CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.s

CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o.requires:

.PHONY : CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o.requires

CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o.provides: CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o.requires
	$(MAKE) -f CMakeFiles/libsyntax7-static.dir/build.make CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o.provides.build
.PHONY : CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o.provides

CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o.provides.build: CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o


# Object files for target libsyntax7-static
libsyntax7__static_OBJECTS = \
"CMakeFiles/libsyntax7-static.dir/src/scanner.c.o" \
"CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o"

# External object files for target libsyntax7-static
libsyntax7__static_EXTERNAL_OBJECTS =

libsyntax7.a: CMakeFiles/libsyntax7-static.dir/src/scanner.c.o
libsyntax7.a: CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o
libsyntax7.a: CMakeFiles/libsyntax7-static.dir/build.make
libsyntax7.a: CMakeFiles/libsyntax7-static.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Linking C static library libsyntax7.a"
	$(CMAKE_COMMAND) -P CMakeFiles/libsyntax7-static.dir/cmake_clean_target.cmake
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/libsyntax7-static.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/libsyntax7-static.dir/build: libsyntax7.a

.PHONY : CMakeFiles/libsyntax7-static.dir/build

CMakeFiles/libsyntax7-static.dir/requires: CMakeFiles/libsyntax7-static.dir/src/scanner.c.o.requires
CMakeFiles/libsyntax7-static.dir/requires: CMakeFiles/libsyntax7-static.dir/src/dict/c_language.c.o.requires

.PHONY : CMakeFiles/libsyntax7-static.dir/requires

CMakeFiles/libsyntax7-static.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/libsyntax7-static.dir/cmake_clean.cmake
.PHONY : CMakeFiles/libsyntax7-static.dir/clean

CMakeFiles/libsyntax7-static.dir/depend:
	cd /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7 /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7 /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug /Users/jerry/src/github.com/xor-gate/c-by-example/libsyntax7/build/Debug/CMakeFiles/libsyntax7-static.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/libsyntax7-static.dir/depend

