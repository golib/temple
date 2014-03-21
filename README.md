# Slim

Slim is an elegant templating engine for Go Programming Language
It is inspired from [slim lang](https://github.com/slim-template/slim)

## Usage

```go
package main

import (
  "os"
  "fmt"

  "github.com/golib/slim"
)

func main() {
  // html string:
  //   h1
  //     = Title
  slimFile = "index.html.slim"

  tpl, err := slim.Parse(slimFile)
  if err != nil {
    fmt.Println("Cannot parse slim file with ", err)
    return
  }

  tpl.Execute(os.Stdout, &map[string]string{"Title": "Slim, golang!"})
}
```
