package main

import "fmt"           // std package
import "rsc.io/quote" 

import "github.com/ivanaspi88/vlib"  

func main() {

  fmt.Println("AIS inception!") // std package "fmt"

  fmt.Println(quote.Go())       // ext module  "rsc.io/quote"

  var a int = vlib.Vad(2,3)   
  fmt.Println("Sum vlib =",a)
}

