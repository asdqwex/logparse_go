// logparse.go - generate dynamic hash maps from log files
//
// a practice project by Matthew Ellsworth - https://github.com/asdqwex
//
// logparse [options] [outfile] <infile>
//
// logparse -cv -t nginx -k 0.0.0.0 -k GET -f total -d 01012016 -d 02022016  -o outfile.log data1.log data2.log data3.log
// logparse --verbose --type=nginx --key=0.0.0.0 --key=GET --type=total --date=01012016 --date=02022016 --outfile=outfile.log -- data1.log data2.log data3.log
//
// -o = outfile
// -c = colored output
// -v = versbose: display logic assertions
// -t = type: nginx || apache || varnish || raw
// -k = key: regex match to use as map key, in order of desired nesting
// -f = function to use when generating values: total || found || percent
// -d = date boundries
// -- = files to read from
//
//
// same as above but - means read from stdin
//
//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

package main

import (
    "fmt"
    "os"
    "github.com/asdqwex/argparse_go"
)

type Argument struct {
    takes_input string
    longname  string
    required bool
    conflicts []string
    dependencies []string
    only_once bool
    reads_from_stdin bool
    hungry bool
}

// Example arg_def_map
//  my_apps_args = {
//    a = {
//      takes_input = 'string' # type of data this function needs
//      longname = 'function1' # if called with -- what is it called
//      required = 0           # does this option always need to be defined
//      conflicts = ['b', 'd'] # other options this function cannot be run with
//      dependencies = ['a']   # other options this function requires
//      only_once = 1          # this function should only be called once
//      reads_from_stdin: 0    # this function read from stdin
//      hungry = 1             # this function takes the rest of the input data
//    },
//    b = {
//      ...
//      ...
//    },
//    c = {
//      ...
//      ...
//    }
// }

a := argparse.Argument{"takes_input": 'string', "longname": 'function1', "required": 0, "conflicts": ['a', 'b'], "dependencies": ['c'], "only_once": 0, "reads_from_stdin": 0, "hungry": 1}
b := argparse.Argument{"takes_input": nil', "longname": 'function2', "required": 0, "conflicts": ['a'], "dependencies": ['c'], "only_once": 1, "reads_from_stdin": 0, "hungry": 0}
c := argparse.Argument{"takes_input": 'string', "longname": 'function3', "required": 0, "conflicts": ['a'], "dependencies": ['b'], "only_once": 1, "reads_from_stdin": 0, "hungry": 1}
d := argparse.Argument{"takes_input": nil, "longname": 'function4', "required": 1, "conflicts": [], "dependencies": [], "only_once": 1, "reads_from_stdin": 1, "hungry": 1}

modules := [*a, *b, *c, *d]

func main() {
  options := argparse.Parse(os.Args)
  // for option in options
    // if the option exists
      // run function with provided data
    // else "ERROR :  no function by that name"
}

// -o = print output to file
func outfile (string) {

}

// -c = colored output
func color (bool, string) {

}

// -v = versbose: display logic assertions
func verbose (bool) {

}

// -t = type: nginx || apache || varnish || raw
func file_type (string) {

}

// -k = key: regex match to use as map key, in order of desired nesting
func key (string) {

}

// -f = function to use when generatin gvalues: total || found || percent
func function_type (string) {

}

// -d = date boundries
func date (string) {

}

// -- = files to read from
func double_dash (bool, string) {

}
