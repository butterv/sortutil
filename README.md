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
        ID uint64
        Capital string
        Latitude float32
        Longitude float32
    }{
        {
            ID:        3,
            Capital:   "Tokyo",
            Latitude:  35.67581,
            Longitude: 139.74507,
        },
        {
            ID:        5,
            Capital:   "Seoul",
            Latitude:  37.531084,
            Longitude: 126.91583,
        },
        {
            ID:        2,
            Capital:   "Beijing",
            Latitude:  39.908222,
            Longitude: 116.391027,
        },
        {
            ID:        1,
            Capital:   "Washington D.C.",
            Latitude:  38.897159,
            Longitude: -77.036207,
        },
        {
            ID:        4,
            Capital:   "London",
            Latitude:  51.499183,
            Longitude: -0.12464066,
        },
        {
            ID:        6,
            Capital:   "Brasilia",
            Latitude:  -15.799668,
            Longitude: -47.864154,
        },
    }

    sortutil.Order(countries).Asc("ID")

    // or
    // sortutil.Order(countries).Desc("ID")
}
```

## Installation

```
$ go get github.com/istsh/sortutil
```

## Author

istsh