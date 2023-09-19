# Sluggo - Convert free text to url-safe strings

## Summary
Sluggo sanitizes snippets of text to produce URL safe slugs compliant with section 2 of [RFC 3986](https://www.ietf.org/rfc/rfc3986.txt)

## Usage

```go
package main

import (
	"fmt"
	"github.com/sapient/sluggo"
)

func main() {
	str := "123 Hello World! 😊"
	fmt.Println(str + ": " + sluggo.Sluggo(str))
}
```

## Todo

 * Options for custom separator (underscore instead of dash, etc.)
 * Option to allow case preservation
 * Option for replacing emojis with custom text via dictionary lookup