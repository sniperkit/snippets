# GoPic

[![License][License-Image]][License-Url]
[![Godoc][Godoc-Image]][Godoc-Url]
[![ReportCard][ReportCard-Image]][ReportCard-Url]
[![Build][Build-Status-Image]][Build-Status-Url]

Secure picture indexer and web-viewer powered by Golang (alpha)

* Single picture unique id. Which is independed of the picture filepath because the id is calculated
  based on the picture content.
* Folder unique id (which contains one or more pictures)

The GoPic index is just a view of already existing folders with pictures. It is up to you
 how you organize the pictures inside folders. You just use the file-browser to organize
 and GoPic will index them for you!

## Installation

For installation from source you need a working Golang toolchain and run the following command:

```
go get github.com/xor-gate/gopic
go install github.com/xor-gate/gopic
$GOPATH/bin/gopic -path /path/to/pictures
```

## Usage

The first run of gopic will index the folder path (with possible subfolders) of jpeg pictures.
Keep in mind indexing time is dependend on amount of pictures and the speed of the harddisk or
network share! For big picture database this can take up more than 15 minutes (46GB/16K files) 
on an SSD.

```
$ gopic -path /path/to/pictures
Running gopic dev
Serving at http://0.0.0.0:8081
Indexing pictures...this will take some time

Summary:
         Started at: 2017-01-10 15:04:54.089956894 +0100 CET
        Finished at: 2017-01-10 15:04:54.978702674 +0100 CET
               Took: 888.781024ms

          New unique files: 6
        Total unique files: 6
           Duplicate files: 0

               Total files: 6
```

The following urls endpoints are available:

* `/admin/<secret>` : Admin panel to view all folders protected by `secret`
* `/by-id/<picture id>` : View single picture by unique ID
* `/by-folder-id/<folder id>` : View a folder gallery by unique ID

## Demo

A demo to the is available use "demo@demo.com:demo" as login:

[http://shulgin.xor-gate.org:8081](http://shulgin.xor-gate.org:8081)

## Design

GoPic uses only Golang packages and doesn't depend on libraries written in other languages (like C/C++).
So it cross-compiles easily and will run on Windows, *NIX, macOS. It uses the great BoltDB for key-value
 on-disk storage.

The Blake2s hash is used to generate unique picture ids which is cryptographic secure and the fastest currently
 in the public domain.

## Why

This project was created due to the lack of run-and-forget picture indexers.
 It is able to serve secure HTTP urls for folders and single pictures which don't 
 have the problem of [Full Path Disclosure](https://www.owasp.org/index.php/Full_Path_Disclosure).

## License

[MIT](LICENSE)

[License-Url]: https://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/badge/license-MIT-blue.svg?maxAge=2592000
[Build-Status-Url]: http://travis-ci.org/xor-gate/gopic
[Build-Status-Image]: https://travis-ci.org/xor-gate/gopic.svg?branch=master
[Godoc-Url]: https://godoc.org/github.com/xor-gate/gopic
[Godoc-Image]: https://godoc.org/github.com/xor-gate/gopic2?status.svg
[ReportCard-Url]: https://goreportcard.com/report/github.com/xor-gate/gopic
[ReportCard-Image]: https://goreportcard.com/badge/github.com/xor-gate/gopic
