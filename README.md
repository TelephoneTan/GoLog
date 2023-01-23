# GoLog
GoLang log utils with color and time.
## Usage
```go
package main

import "github.com/TelephoneTan/GoLog/log"

func main() {
	msg := "hello, world"
	log.I(msg)
	log.S(msg)
	log.W(msg)
	log.E(msg)
	log.F(msg)
	log.S(msg)
}

```