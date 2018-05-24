lighttpd:
	lighttpd -D -f lighttpd.conf
secdl:
	$(GOBIN)/secdl-http -port=8123 -prefix=/dl/ -static=../../static --root=../../storage
bench-concurrent:
	ab -n 10000 -c 100 http://127.0.0.1:8123/dl/f5fbba124cfd2367b78e7dd2591b55b4/5665f363/secret-page.html
bench-single:
	ab -n 10000 http://127.0.0.1:8123/dl/f5fbba124cfd2367b78e7dd2591b55b4/5665f363/secret-page.html
