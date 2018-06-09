package main

import (
	"fmt"
	"time"
	"os"

	"github.com/egymgmbh/go-prefix-writer/prefixer"
)

func main() {
	p := prefixer.New(os.Stdout, func() string {
		return time.Now().Format("2006-01-02 15:04:05: ")
	})

	fmt.Fprint(p, "hello")
	time.Sleep(time.Second)
	fmt.Fprint(p, "\nworld")
}
