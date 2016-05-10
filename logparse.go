package main

import (
    "fmt"
    "os"
    "github.com/asdqwex/argparse_go"
)

a := Argument{"takes_input": 'string', "longname": 'function1', "required": 0, "conflicts": [], "dependencies": [], "only_once": 0, "reads_from_stdin": 0, "hungry": 1}
b := Argument{"takes_input": nil', "longname": 'function2', "required": 0, "conflicts": ['a'], "dependencies": ['c'], "only_once": 1, "reads_from_stdin": 0, "hungry": 0}
c := Argument{"takes_input": 'string', "longname": 'function3', "required": 0, "conflicts": ['a'], "dependencies": ['b'], "only_once": 1, "reads_from_stdin": 0, "hungry": 1}
d := Argument{"takes_input": nil, "longname": 'function4', "required": 1, "conflicts": [], "dependencies": [], "only_once": 1, "reads_from_stdin": 1, "hungry": 1}

arguments := [*a, *b, *c, *d]

func main() {
  fmt.Print(argparse.Parse(os.Args, arguments))

}
