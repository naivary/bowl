package main

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/naivary/bowl/pkg/bowl"
)

func main() {

	fn := func() int {
		return rand.Int()

	}

	b := bowl.New(3, fn)
	i := 0
	wg := new(sync.WaitGroup)
	wg.Add(22)

	for i < 22 {

		go func() {
			defer wg.Done()
			s := rand.Int()
			b.Return(s)
		}()
		i++
	}

	wg.Wait()

	i = 0
	wg.Add(3)

	for i < 3 {
		go func() {
			defer wg.Done()
			j := b.Get()
			fmt.Println(j)
		}()
		i++
	}
}
