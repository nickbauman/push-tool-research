# Go Client and Server for benchmarking

## Compiling

From within the root of the `async_golang` sub project, run

    ./bandi.sh

### Issues compiling

The Go language compiler requires very strict layout of your code and
environment. Therefore this language package looks the way it does.

The following is my `go env` output:

    nick@nickb-gup:~/workspaces/push-tool-research/async_golang$ go env
    GOARCH="amd64"
    GOBIN=""
    GOCHAR="6"
    GOEXE=""
    GOHOSTARCH="amd64"
    GOHOSTOS="linux"
    GOOS="linux"
    GOPATH="/home/nick/workspaces/push-tool-research/async_golang"
    GORACE=""
    GOROOT="/usr/lib/go"
    GOTOOLDIR="/usr/lib/go/pkg/tool/linux_amd64"
    CC="gcc"
    GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0"
    CXX="g++"
    CGO_ENABLED="1"
    GOMAXPROCS=8

YMMV

## Benchmarking

Starting the server is as easy as

    ./main server localhost:8080

from within the sub project root. This binds the server the `localhost` on port
8080, of course.

Benching the server or client is as easy as

    ./main client localhost:8080 10000000

Which will go after a server running on your local 8080 for 10000000, reporting
the benchmarks using STDOUT.
