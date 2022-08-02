package unit

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/naivary/bowl/pkg/bowl"
)

func fn() int {
	s := rand.Int()
	return s
}

func TestNew(t *testing.T) {
	b := bowl.New(5, fn)

	if b.Max() != 5 {
		t.Errorf("size is not 5. Expected 5 got: %d", b.Size())
		return
	}

}

func TestGetFromEmptyPool(t *testing.T) {
	b := bowl.New(5, fn)

	o := b.Get()

	fmt.Println(o)

}

func TestGetFromPool(t *testing.T) {
	i := 0
	b := bowl.New(5, fn)

	for i < 10 {
		s := rand.Int()
		b.Return(s)
		i++
	}

	if b.Size() != 5 {
		t.Errorf("size should be 5. Got: %d", b.Size())
		return
	}

	b.Get()

	if b.Size() != 4 {
		t.Errorf("size should be 4. Got: %b", b.Size())
		return
	}

}

func TestReturn(t *testing.T) {
	b := bowl.New(5, fn)
	b.Return(3)

	if b.Size() != 1 {
		t.Errorf("size sohuld be 1. Got: %d", b.Size())
		return
	}

}

func TestParallelReturn(t *testing.T) {
	b := bowl.New(5, fn)
	i := 0

	wg := new(sync.WaitGroup)
	wg.Add(5)

	for i < 5 {
		go func() {
			defer wg.Done()
			s := rand.Int()
			b.Return(s)
		}()
		i++
	}

	wg.Wait()

	if b.Size() != 5 {
		t.Errorf("size should be equal to 5. Got %d", b.Size())
		return
	}

}

func TestReturnLimit(t *testing.T) {
	b := bowl.New(5, fn)
	i := 0

	wg := new(sync.WaitGroup)
	wg.Add(20)

	for i < 20 {
		go func() {
			defer wg.Done()
			s := rand.Int()
			b.Return(s)
		}()
		i++
	}

	wg.Wait()

	if b.Size() != b.Max() {
		t.Errorf("size should be equal to %d. Got %d", b.Max(), b.Size())
		return
	}

}

func TetstGetConcurrently(t *testing.T) {
	b := bowl.New(5, fn)
	i := 0

	wg := new(sync.WaitGroup)
	wg.Add(20)

	for i < 20 {
		go func() {
			defer wg.Done()
			s := rand.Int()
			b.Return(s)
		}()
		i++
	}

	wg.Wait()

	i = 0

	if b.Size() != b.Max() {
		t.Errorf("size should be equal to %d. Got %d", b.Max(), b.Size())
		return
	}

	wg.Add(5)

	for i < 5 {
		go func() {
			defer wg.Done()
			b.Get()
		}()
		i++
	}

	wg.Wait()

	if b.Size() != 0 {
		t.Errorf("Size should be 0. Got: %d", b.Size())
		return
	}

}
