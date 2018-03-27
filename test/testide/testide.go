package main

import (
	"fmt"
)

func main() {
	fmt.Println("VSCode debug tool works!")
}

/*
How to install the dev tool
go env GOPATH
git clone https://github.com/derekparker/delve.git ~/go/src/github.com/derekparker/delve
cd ~/go/src/github.com/derekparker/delve/
make install
add $GOPATH/bin to PATH
*/
