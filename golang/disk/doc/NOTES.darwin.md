// http://internals.exposed/blog/dtrace-vs-sip.html

root@Jerrys-MacBook-Pro ~ $ csrutil enable --without dtrace
csrutil: requesting an unsupported configuration. This is likely to break in the future and leave your machine in an unknown state.
csrutil: failed to modify system integrity configuration. This tool needs to be executed from the Recovery OS.
root@Jerrys-MacBook-Pro ~ $ dtruss -f -t open diskutil -l

dtrace: failed to execute diskutil: dtrace cannot control executables signed with restricted entitlements
