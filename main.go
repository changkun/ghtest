package main

import (
	"fmt"

	"changkun.de/x/dummydep"
	"golang.org/x/sync/errgroup"
)

func main() {
	g := errgroup.Group{}
	g.Go(func() error {
		fmt.Println(dummydep.Hello())
		return nil
	})
	g.Wait()

}
