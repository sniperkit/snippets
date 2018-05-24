This is the eRESP library for Arduino

Installation
--------------------------------------------------------------------------------

To install this library, just place this entire folder as a subfolder in your
Arduino/lib/targets/libraries folder.

When installed, this library should look like:

Arduino/lib/targets/libraries/eRESP              (this library's folder)
Arduino/lib/targets/libraries/eRESP/eresp.cpp    (the library implementation file)
Arduino/lib/targets/libraries/eRESP/eresp.h      (the library description file)
Arduino/lib/targets/libraries/eRESP/keywords.txt (the syntax coloring file)
Arduino/lib/targets/libraries/eRESP/examples     (the examples in the "open" menu)
Arduino/lib/targets/libraries/eRESP/readme.txt   (this file)

Building
--------------------------------------------------------------------------------

After this library is installed, you just have to start the Arduino application.
You may see a few warning messages as it's built.

To use this library in a sketch, go to the Sketch | Import Library menu and
select eRESP.  This will add a corresponding line to the top of your sketch:
#include <eresp.h>

To stop using this library, delete that line from your sketch.

Geeky information:
After a successful build of this library, a new file named "eresp.o" will appear
in "Arduino/lib/targets/libraries/eRESP". This file is the built/compiled library
code.

If you choose to modify the code for this library (i.e. "eresp.cpp" or "eresp.h"),
then you must first 'unbuild' this library by deleting the "eRESP.o" file. The
new "eRESP.o" with your code will appear after the next press of "verify"

