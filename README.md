# gomodfinder

gomodfinder - library for finding go.mod file

**Example**

```go
package main

import (
	"fmt"

	"github.com/artarts36/gomodfinder"
)

func main() {
	goMod, err := gomodfinder.Find("./", 5)
	if err != nil {
		panic(err)
	}

	fmt.Println(goMod.Module.Mod.Path)
}
```
