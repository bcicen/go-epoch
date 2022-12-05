# go-epoch

A simple go package / function to automatically convert Unix epoch timestamps (in seconds, milliseconds, or nanoseconds) to `time.Time`.

> note this method only reliably works for timestamps between `1970-04-17 18:02:52 UTC` and `2262-04-11 23:47:16 UTC`

## Usage

```go
package main

import (
	"fmt"

	"github.com/bcicen/epoch"
)

func main() {
	fmt.Println(epoch.Time(1667015475))           // 2022-10-29 03:51:15 +0000 UTC
	fmt.Println(epoch.Time(1667015475000))        // 2022-10-29 03:51:15 +0000 UTC
	fmt.Println(epoch.Time(1667015475000000000))  // 2022-10-29 03:51:15 +0000 UTC
}
```
