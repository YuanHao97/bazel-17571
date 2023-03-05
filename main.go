package main

import (
        "fmt"
        "log"

        // "github.com/bazelbuild/rules_go/go/runfiles"
		// the above import works well, however my own repo source code doesn't work
        "github.com/YuanHao97/rules_go/go/runfiles"
	)

func main() {
        loc, err := runfiles.Rlocation("go_sdk/bin/gofmt")
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println(loc)
}