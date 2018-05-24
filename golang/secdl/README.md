# secdl

Secdl is a Golang implementation of the famous [Lighttpd SecDownload](http://redmine.lighttpd.net/projects/1/wiki/Docs_ModSecDownload) module

![Stability][Stability-Image]
[![License][License-Image]][License-Url]
[![Godoc][Godoc-Image]][Godoc-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]
[![Build][Build-Status-Image]][Build-Status-Url]
[![Coverage][Coverage-Image]][Coverage-Url]

## Difference with ModSecDownload

The ModSecDownload module has a limitation to generate only "never"
expirable tokens. Because the timestamp is checked at the server.
The `secdl` package implements expiry with the following different
grades: 10 minutes (`10m`), 1 hour (`h`), 1 day (`d`),
1 week (`1w`), 2 weeks, (`2w`), 1 month (`m`) and never (`n`). These are carefully
chosen and also used by snippet paste-bins (e.g [pastebin.com](http://pastebin.com)).

When a token is decoded by `secdl` it loops over the list of known
expire times and then checks for a match against the re-generated token.

## Usage and installation

Make sure `GOPATH` and `GOBIN` is correctly set.

```
$ go get -v github.com/xor-gate/secdl/cmd/secdl
$ secdl -listen=":8181" -root="/tmp" -prefix="/"
[secdl] Serving HTTP at :8181 from path "/tmp" (prefix "/", secret "1234")
```

## Why

The Lighttpd ModSecDownload module was created because of the problem described below. And the `secdl` package is based on this ideas.

Serving secured downloads which will expire. There are multiple ways to handle this:

1. Use the webserver and the internal HTTP authentication
2. Use the application to authenticate and send the file
   through the application

Both ways have limitations:

* Webserver
  * Pros
    * Fast download
    * No additional system load
  * Cons
    * Inflexible authentication handling

* Application
  * Pros 
    * Integrated into the overall layout
    * Very flexible permission management
  * Cons
    * The download occupies an application thread/process

A simple way to combine the two ways could be:

1. App authenticates user and checks permissions to
download the file.
2. App redirects user to the file accessable by the webserver
for further downloading.
3. The webserver transfers the file to the user.

As the webserver doesn't know anything about the permissions
used in the app, the resulting URL would be available to every
user who knows the URL.

The `secdl` package removes this problem by introducing a way to
authenticate a URL for a specified time. The application has
to generate a token and a timestamp which are checked by the
webserver before it allows the file to be downloaded by the
webserver.

The generated URL has the following format:

```
<uri-prefix>/<token>/<timestamp-in-hex>/<relative-path>
```

Which looks like:

```
"yourserver.com/bf32df9cdb54894b22e09d0ed87326fc/435cc8cc/my/shared/folders/secure.tar.gz
```

**token** is an MD5 checksum of:

1. A secret string (user supplied)
2. Relative path (starts with /)
3. Unix timestamp, hexadecimal encoded

As you can see, the token is not bound to the user at all. The
only limiting factor is the timestamp which is used to
invalidate the URL.

If the user tries to fake the URL by choosing a random token,
HTTP status `403 Forbidden` will be sent out.

If the timeout is reached, HTTP status `410 Gone` will be
sent.

If token and timeout are valid, the `<relative-path>` is appended to
the configured `prefix` and passed to the
normal internal file transfer functionality. This might lead to
HTTP status `200 OK` or `404 Not Found`.

## Token generation

Ensure that the token is also hexadecimal encoded. Depending on
the programming language you use, there might be no extra
step for this. For instance, in PHP, the MD5 function
returns the Hex value of the digest. If, however, you use a
language such as Java or Python, the extra step of converting
the digest into Hex is needed.

## See also

A reference webserver implementation in Golang can be found here: [Gohide](https://github.com/xor-gate/gohide)

## License

[MIT](LICENSE)

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg?maxAge=2592000
[Stability-Image]: https://img.shields.io/badge/stability-unstable-yellow.svg
[Build-Status-Url]: http://travis-ci.org/xor-gate/secdl
[Build-Status-Image]: https://travis-ci.org/xor-gate/secdl.svg?branch=master
[Coverage-Url]: https://coveralls.io/r/xor-gate/secdl?branch=master
[Coverage-image]: https://img.shields.io/coveralls/xor-gate/secdl.svg
[Godoc-Url]: https://godoc.org/github.com/xor-gate/secdl
[Godoc-Image]: https://godoc.org/github.com/xor-gate/secdl?status.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/xor-gate/secdl
[ReportCard-Image]: https://goreportcard.com/badge/github.com/xor-gate/secdl
