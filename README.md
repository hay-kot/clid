# CLID

CLID is an decoder for the [urfave/cli](https://github.com/urfave/cli) that provide some basic decoder from the `*cli.Context` into a struct using reflection and `cli` tags. This library is very much experimental and should likely be vendored into your code base instead of using directly as a package dependency.

Only some types are supported and tested

- String
- Bool
- Int (int, int8, int16, int32, int64)
- Float (float32, float64)
- Uint (uint, uint8, uint16, uint32, uint64)

**Example**

```go
package main

import (
	"fmt"
	"os"

	"github.com/hay-kot/clid"
	"github.com/urfave/cli/v2"
)

type Nested struct {
	FlagString string `cli:"nested-string"`
}

type Struct struct {
	Nested     Nested
	FlagString string `cli:"string"`
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "string",
			Value: "default_1",
		},
		&cli.StringFlag{
			Name:  "nested-string",
			Value: "default_2",
		},
	}

	app.Action = func(c *cli.Context) error {
		var s Struct
		if err := clid.Decode(c, &s); err != nil {
			return err
		}
		fmt.Printf("%+v\n", s)
		return nil
	}

	app.Run(os.Args)
    // OUTPUT:
    // {Nested:{FlagString:default_2} FlagString:default_1}
}

```