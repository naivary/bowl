# BOWL <br>

Bowl is a simple and lighweight in-memory pool without any dependencies.
It is safe for concurrent use and is using generics to reseve
the advantages of type safety provided by Go.

# Install

```shell
go get github.com/naivary/bowl
```

# Example usage

```Go

import (
  "math/rand"
  "github.com/naivary/bowl"
)

newFn := func() int {
  return rand.Int()
}


func main() {

  // create a new bowl with the default
  // limit which is 5. You can change it by
  // setting bowl.DefaultMax. newFn is the
  // factory that will be used.
  bw := bowl.New(0, newFn)

  // get an object from
  // the bowl. If tehere are none newFn
  // will be used to create one
  i := bw.Get()

  // worker is doing some work
  worker(i)

  // return `i` back to the bowl
  // after worker is done with the work
  bw.Return(i)
}
```
