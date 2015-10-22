# goCounter

A webserver that implements a page count service on port 3000.s

## Run with Docker
```
$ docker build -t gocounter
$ docker run -p 3000:3000 gocounter
```

## Run natively (with Go installed)
```
$ go build
$ ./goCounter
```
