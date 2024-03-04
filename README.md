Generic Golang Secrets
======================

This package uses Go 1.18's secrets to implement a wrapping structure that will
unmarshal secret values but marshals them to the default / blank value.  This
is useful when dealing with values that you don't want to have accidentally
exposed in a type safe manner.

Example
-------

```go
import (
    "encoding/json"
    "fmt"

    secret "github.com/andrewbenton/go-secrets"
)

type myConfig struct {
    Username string                `json:"username"`
    Password secret.Secret[string] `json:"password"`
}

func main() {
    input := `{"username": "andrew", "password": "super secret password"}`

    var cfg myConfig
    if err := json.Unmarshal(input, &cfg); err != nil {
        panic(err)
    }

    // secret is hidden when marshaling
    data, err := json.Marshal(cfg)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(data))

    // recursive reflection based printing won't reveal the secret
    fmt.Printf("%#v\n", cfg)

    // secret isn't shown!
    fmt.Println(cfg)

    // secret still isn't shown!
    fmt.Println(cfg.Password)

    // secret can be directly accessed
    fmt.Println(cfg.Password.Get())
}
```

Performance
-----------

Performance isn't the primary concern of this library, however it's important
to recognize that the additional functionality required to keep secrets hidden
costs more.  The tests below contrast the JSON marshaling and unmarshaling
performance of a `string` vs a `Secret[string]`.  These results were executed
against an AMD Ryzen 9 7900X running in 105W Eco mode.

Go 1.22.0:
```
BenchmarkJsonMarshal-24                 10483046               104.5 ns/op
BenchmarkJsonUnmarshal-24                6155026               175.8 ns/op
BenchmarkSecret_MarshalJSON-24           7724763               147.7 ns/op
BenchmarkSecret_UnmarshalJSON-24         3029302               397.4 ns/op
```
