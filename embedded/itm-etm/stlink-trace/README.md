stlink-trace
============

From: https://github.com/avrhack/stlink-trace
And: http://essentialscrap.com/tips/arm_trace/

ST-Link V2 ITM trace utility

This utility can be used with an ST-Link V2 JTAG device connected to an STM32Fxxx series microcontroller to capture the ITM data sent via the printf port (ITM stimulus port 0).

The ITM trace on the ST-Link uses the SWO pin which needs to be connected. On an STM32F303 for example it is PB3. The protocol is a simple 2Mb/s async serial so you must ensure the ITM trace is set up correctly depending on the clock frequency of your particular processor.

changes in this fork
====================

If you interrupt the running code it now correctly tidies up the stlink so you can easily access it again from e.g. st-util for debugging etc. This saves having to unplug/replug the stlink and/or reset it.

Use the '-c NN' option on the command line to set the core frequency where NN is in MHz ie for a 72MHz use stlink-trace -c 72. If you omit the -c option, stlink-trace will not set the ITM/SWO speed on the assumption that you have already done this in your own firmware - that's how I do mine as it means you can use firmware macros to ensure it's correct.

This fork creates the same two files as the original (see the code for default names) however the 'raw' file now includes the packet format according to the ARM ITM specs so you can post-parse that if necessary. The non-raw file only includes data bytes from the ITM 0 stimulus port.

There is now a very simple makefile to make building a little easier from the command line.

Build
-----
Eclipse project files can be used. Alternatively use the following:

gcc stlink-trace.c -lusb-1.0 -L/usr/local/lib -o stlink-trace

TODO
----
- Sort out timestamps (partly done offline, will be uploaded when it's working!)



All credit to obe1line for the original hard work; my little changes above are just usability ones but Chris did all the heavy lifting 2 years ago. Thanks Chris! :)

