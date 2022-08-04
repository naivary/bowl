# Bowl <br>

Bowl is a simple and lighweight in-memory pool without any dependencies.
It is safe for concurrent use and is using generics to reseve type safety.

# Install

```shell
go get github.com/naivary/bowl
```

# Example usage

```go

import (
  "math/rand"
  "github.com/naivary/bowl"
)

fn := func() int {
  return rand.Int()
}


func main() {

  // create a new bowl with the default
  // limit which is 5. You can change it by
  // setting bowl.DefaultMax. `fn` is the
  // factory that will be used.
  bw := bowl.New(0, fn)

  // Get an object from the bowl.
  // If there no objects in the pool
  // an object will be created using `fn`
  i := bw.Get()

  // do some work...
  worker(i)

  // return `i` back to the bowl
  // after worker is done. If no
  // clean function is set using
  // bw.SetClean, no cleanig will
  // be done. Its recommended to
  // defer the bw.Return(i).
  bw.Return(i)

}
```
