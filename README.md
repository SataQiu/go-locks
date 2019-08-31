# go-locks

Some extended locks for GoLang.

## Locks

### ReentrantLock

A kind of Lock that allow repeated locking.

Example:

```go
package main

import (
	"github.com/SataQiu/go-locks"
)

func main() {
	reenLock := &locks.ReentrantLock{}
	reenLock.Lock()
	defer reenLock.Unlock()
	// do something ...
	reenLock.Lock()
	// do something ...
	reenLock.Unlock()
}
```

## Contribution

Any contribution is welcome! Feel free to send PRs and Issues.