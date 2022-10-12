# CLID

CLID is an decoder for the [urfave/cli](https://github.com/urfave/cli) that provide some basic decoder from the `*cli.Context` into a struct using reflection and `cli` tags. This library is very much experimental and should likely be vendored into your code base instead of using directly as a package dependency.

**Example**

```go
type Nested struct {
	FlagString string `cli:"nested-string"`
}

type Struct struct {
	Nested      Nested
	FlagString  string  `cli:"string"`
}

func main() {
    app := cli.NewApp()
    app.Flags = []cli.Flag{
        cli.StringFlag{
            Name:  "string",
            Value: "default",
        },
        cli.StringFlag{
            Name:  "nested-string",
            Value: "default",
        },
    }

    app.Action = func(c *cli.Context) error {
        var s Struct
        if err := clid.Decode(c, &s); err != nil {
            return err
        }
        fmt.Printf("%+v", s)
        return nil
    }

    app.Run(os.Args)
}
```

```