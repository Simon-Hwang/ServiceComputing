package hello

import (
    "github.com/spf13/pflag" 
    "fmt"
    "os"
)

func hello() {
    var port int
    var bol bool
    pflag.IntVar(&port, "p", 8000, "specify port to use.  defaults to 8000.")
    pflag.BoolVar(&bol, "b", false, "specify port to use.  defaults to false.")
    pflag.Parse()
    for _, i := range os.Args[1:]{
        fmt.Println(i)
    }
    fmt.Printf("port = %d\n", port)
    fmt.Printf("other args: %+v\n", pflag.Args())
}