[![GoDoc](https://godoc.org/github.com/mickep76/encoding?status.svg)](https://godoc.org/github.com/mickep76/encoding)
[![codecov](https://codecov.io/gh/mickep76/encoding/branch/master/graph/badge.svg)](https://codecov.io/gh/mickep76/encoding)
[![Build Status](https://travis-ci.org/mickep76/encoding.svg?branch=master)](https://travis-ci.org/mickep76/encoding)
[![Go Report Card](https://goreportcard.com/badge/github.com/mickep76/encoding)](https://goreportcard.com/report/github.com/mickep76/encoding)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/mickep76/mlfmt/blob/master/LICENSE)

# encoding

Package provides a generic interface to encoders and decoders

## Example

```go
package main
  
import (
        "flag"
        "fmt"
        "log"
        "strings"

        "github.com/mickep76/encoding"
        _ "github.com/mickep76/encoding/json"
        _ "github.com/mickep76/encoding/toml"
        _ "github.com/mickep76/encoding/yaml"
)

type Message struct {
        Name, Text string
}

type Messages struct {
        Messages []*Message
}

func main() {
        enc := flag.String("enc", "json", fmt.Sprintf("Encodings: [%s].", strings.Join(encoding.Encodings(), ", ")))

        flag.Parse()

        in := Messages{
                Messages: []*Message{
                        &Message{Name: "Ed", Text: "Knock knock."},
                        &Message{Name: "Sam", Text: "Who's there?"},
                        &Message{Name: "Ed", Text: "Go fmt."},
                        &Message{Name: "Sam", Text: "Go fmt who?"},
                        &Message{Name: "Ed", Text: "Go fmt yourself!"},
                },
        }

        e, err := encoding.NewEncoding(*enc)

        b, err := e.Encode(in)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Printf("Encoded:\n%s\n", string(b))

        out := Messages{}
        if err := e.Decode(b, &out); err != nil {
                log.Fatal(err)
        }

        fmt.Println("Decoded:")
        for _, m := range out.Messages {
                fmt.Printf("%s: %s\n", m.Name, m.Text)
        }
}
```
