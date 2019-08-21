# sortutil

Easy-to-use sort library for go

## Usage

```go
package main

import (
    "github.com/istsh/sortutil"
)

func main() {
    countries := []struct{
        ID     uint64
        Name   string
        Age    uint
        Height float32
        Weight float32
    }{
        {
            ID:     5,
            Name:   "Zach",
            Age:    30,
            Height: 175.0,
            Weight: 75.0,
        },
        {
            ID:     3,
            Name:   "Angie",
            Age:    20,
            Height: 160.5,
            Weight: 55.0,
        },
        {
            ID:     2,
            Name:   "Taro",
            Age:    18,
            Height: 179.5,
            Weight: 70.0,
        },
        {
            ID:     4,
            Name:   "Kim",
            Age:    25,
            Height: 150.2,
            Weight: 48.0,
        },
        {
            ID:     1,
            Name:   "Tom",
            Age:    20,
            Height: 175.5,
            Weight: 70.0,
        },
    }

    sortutil.Order(countries).Asc("ID").Exec()

    // or
    // sortutil.Order(countries).Asc("ID").Desc("Name").Exec()
    // sortutil.Order(countries).Asc("ID").Desc("Name").Asc("Height").Exec()
}
```

## Installation

```
$ go get github.com/istsh/sortutil
```

## Author

istsh