# study.go



- install go

`brew install go`

- workspace

`mkdir ~/go`

`export GOPATH=~/go`

`export PATH=$PATH:$GOPATH/bin`

`mkdir -p ~/go/src/github.com/mingz2013`

- hello go

`cd ~/go/src/github.com/mingz2013/`

`git clone https://github.com/mingz2013/study.go.git`

`cd study.go`

`mkdir hello`

`cd hello`

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







