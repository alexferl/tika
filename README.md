# tika [![Go Report Card](https://goreportcard.com/badge/github.com/admiralobvious/tika)](https://goreportcard.com/report/github.com/admiralobvious/tika)

tika is a Golang client for the [Apache Tika](https://tika.apache.org/) (1.16) REST server.

## Installing
`go get github.com/admiralobvious/tika`

## Using
This assumes you have a Tika 1.16 server running locally on port 9998.

Simple client:

``` go
package main

import (
	"fmt"
	"log"

	"github.com/admiralobvious/tika"
)

func main() {
	c := tika.NewClient(&tika.Options{Url: "http://localhost:9998"})

	hi, err := c.Hello()
	if err != nil {
		log.Fatalf("Error getting hello: %v", err)
	}

	fmt.Printf("Server replied: %s", hi)
}

```

Output:
```
Server replied: This is Tika Server (Apache Tika 1.16). Please PUT
```

More complex client in [examples](examples).
