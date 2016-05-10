package main

import (
    "fmt"
    "os"
    "github.com/asdqwex/argparse_go"
)

func main() {
  fmt.Print(argparse.Parse(os.Args))

}
