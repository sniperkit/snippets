ARDUINO_DIR   = /Applications/Arduino.app/Contents/Java
MONITOR_PORT  = /dev/tty.usbserial-A800etxu
BOARD_TAG     = diecimila
MCU           = atmega328

include /usr/local/opt/arduino-mk/Arduino.mk

# TODO for some reason the arduino makefile castle is unable to create a working upload target so we have burn now
burn:
	/Applications/Arduino.app/Contents/Java/hardware/tools/avr/bin/avrdude -q -V -p m328p -C /Applications/Arduino.app/Contents/Java\
/hardware/tools/avr/etc/avrdude.conf -D -c arduino -b 57600 -P /dev/tty.usbserial-A800etxu \
      -U flash:w:build-diecimila/arduino.hex:i
