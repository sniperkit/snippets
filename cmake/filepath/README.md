# filepath

The filepath example demonstrates a solution for fixing the absolute
 path of compiler macro `__FILE__`. CMake always generates absolute
 paths to files which resolves a lot of compiling problems. When
 writing a logging facility which uses this macro this is not desirable.

A trick is created to create a pointer macro inside the `__FILE__` string
 where it points after the known absolute path to the sources `${CMAKE_SOURCE_DIR}`.

I got this from the cmake developer mailing list ([here](https://cmake.org/pipermail/cmake-developers/2015-January/024202.html)):

	[cmake-developers] Adding an option for relative compiler invocations

	Brad King brad.king at kitware.com 
	Fri Jan 23 16:38:08 EST 2015
	On 01/22/2015 04:46 AM, Michael Enßlin wrote:
	> (1.2) Using compile-time string manipulations to sanitize the filename.
	>       Due to limitations of C++, this requires template metaprogramming,
	>       leading to unreasonable complexity and compile times.

	See below.

	> Over the last several decades, at least on the POSIX platform, it has
	> become common practice to invoke compilers with relative file paths

	This is only true for in-source builds.  CMake recommends out-of-source,
	and then full paths are much more reliable.  Even if one used relative
	paths then your messages would still have a bunch of "../" in them for
	an out-of-source build.  Therefore I'll assume you're using in-source
	builds.

	Side note: To make relative paths behave right with __FILE__ you would
	also need all include paths (-I) to be relative.  Otherwise headers will
	still get full paths.  This would require much more work than just
	trying to get compile lines to refer to source files with a relative
	path.

	So, assuming you have an in-source build then all sources and headers
	must sit under a single prefix (the top of the source tree).  With that
	you can do something like:

	 string(LENGTH "${CMAKE_SOURCE_DIR}/" SRC_DIR_LEN)
	 add_definitions(-DSRC_DIR_LEN=${SRC_DIR_LEN})

	and then in the source code do:

	 #define MY__FILE__ (__FILE__+SRC_DIR_LEN)

	That will give you a compile-time constant __FILE__ with a relative path
	and no runtime overhead.  Use MY__FILE__ in your log macros.

	-Brad
