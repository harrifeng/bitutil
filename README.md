# bitutil
Golang bit relative operation utility


+ Hello World example

```
package main

import (
	"fmt"
	"os"

	"github.com/harrifeng/bitutil"
)

func main() {

	fmt.Println(bitutil.GetTwoComplement8(127))
	fmt.Println(bitutil.GetTwoComplement8(1))
	fmt.Println(bitutil.GetTwoComplement8(-1))
	fmt.Println(bitutil.GetTwoComplement8(-128))
	os.Exit(0)
}

// <===================OUTPUT===================>
// 01111111
// 00000001
// 11111111
// 10000000
```
