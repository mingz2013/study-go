# study.go



- install go

`brew install go`

- workspace

`git clone https://github.com/mingz2013/study.go.git ~/go`

`export GOPATH=~/go`

`export PATH=$PATH:$GOPATH/bin`

`mkdir -p ~/go/src/github.com/mingz2013`

- hello go

`mkdir ~/go/src/github.com/mingz2013/hello`

`cd ~/go/src/github.com/mingz2013/hello`

`touch hello.go`

`vim hello.go`

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.\n")
}
```

`go install`

`hello`







