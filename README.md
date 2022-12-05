# go-epoch

A simple go package / function to automatically convert Unix epoch timestamps (in seconds, milliseconds, or nanoseconds) to `time.Time`.

_note_ this method only reliably works for times after date 999999999 (2001-09-09 01:46:39 UTC).

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
