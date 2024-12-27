package main

import "fmt"          // std package
import "rsc.io/quote" // https://pkg.go.dev 

//import "vlib1"   // local packages via "replace" - from my catalog for projects_

func main() {

  fmt.Println("AIS inception!") // std package "fmt"

  fmt.Println(quote.Go())       // ext module  "rsc.io/quote"

  //var a int = vlb1.Vad(2,3)   // local package "vlb1" from LIB1 catalog
  //fmt.Println("Sum1 =",a)
}

