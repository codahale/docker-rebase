docker-rebase
=============

[![Build Status](https://travis-ci.org/codahale/docker-rebase.png?branch=master)](https://travis-ci.org/codahale/docker-rebase)

`docker-rebase` is used to "rebase" a Docker image against an base image. The
resulting image contains only those layers which are unique to the downstream
image:

```
docker save base > base.tar
docker save app | docker-rebase base.tar > app.tar
```

To install, run `go get -u github.com/codahale/docker-rebase`.

For documentation, check [godoc](http://godoc.org/github.com/codahale/docker-rebase).
