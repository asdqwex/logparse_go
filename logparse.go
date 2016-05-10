// logparse - generate dynamic hash maps from log files
//
//
//
// logparse -pcv -t nginx -k 0.0.0.0 -k GET -f total -d 01012016 -d 02022016 -- data1.log data2.log data3.log
// logparse --print --color --verbose --type nginx --key 0.0.0.0 --key GET --type total --date 01012016 --date 02022016 -- data1.log data2.log data3.log
//
// -p = print output to stdout
// -c = colored output
// -v = versbose: display logic assertions
// -t = type: nginx || apache || varnish || raw
// -k = key: regex match to use as map key, in order of desired nesting
// -f = function to use when generatin gvalues: total || found || percent
// -d = date boundries
// -- = files to read from
//
// logparse --print --color --verbose --type nginx --key 0.0.0.0 --key GET --type total --date 01012016 --date 02022016 - < cat data1.log
//
// same as above but - means read from stdin
//

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
