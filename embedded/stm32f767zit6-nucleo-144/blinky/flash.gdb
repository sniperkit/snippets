target extended-remote :4242
monitor reset halt
file main.elf
load main.elf
monitor reset init
#continue
