# Be more quiet
ifneq ($(V),1)
Q    := @
NQMAKE := $(MAKE) --no-print-directory
MAKE := $(Q)$(MAKE) --no-print-directory
NULL := 2>/dev/null
endif

all: debug

debug: build/Debug
	cd $^ && cmake -DCMAKE_BUILD_TYPE=Debug ../../ && make
release: build/Release
	cd $^ && cmake -DCMAKE_BUILD_TYPE=Release ../../ && make

ci: debug release coverage

coverage: test
	cd build/Debug && make coverage

test: debug
	cd build/Debug && make test

build/Debug:
	mkdir -p $@

build/Release:
	mkdir -p $@

clean:
	rm -Rf build

.PHONY: clean
