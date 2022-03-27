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

    gs "github.com/andrewbenton/go-secrets"
)

type myConfig struct {
    Username string            `json:"username"`
    Password gs.Secret[string] `json:"password"`
}

func main() {
    input := `{"username": "andrew", "password": "super secret password"}`

    var cfg myConfig
    if err := json.Unmarshal(input, &cfg); err != nil {
        panic(err)
    }

    // secret is hidden when marshaling
    data, err := json.Marshal(cfg)
    if err != nil { panic(err) }
    fmt.Println(string(data))

    // recursive reflection based printing won't reveal the secret
    fmt.Printf("%#v\n", cfg)

    // secret isn't shown!
    fmt.Println(cfg)

    // secret still isn't shown!
    fmt.Println(cfg.Password)

    // secret can be directly accessed
    fmt.Println(cfg.Password.Value)
}
```

Performance
-----------

Performance isn't the primary concern of this library, however it's important
to recognize that the additional functionality required to keep secrets hidden
costs more.  The tests below contrast the JSON marshaling and unmarshaling
performance of a `string` vs a `Secret[string]`.

```
BenchmarkJsonMarshal-32             	 3077210	       384.4 ns/op
BenchmarkJsonUnmarshal-32           	 1799340	       664.4 ns/op
BenchmarkSecret_MarshalJSON-32      	 1720758	       732.2 ns/op
BenchmarkSecret_UnmarshalJSON-32    	  710397	      1537   ns/op
```
