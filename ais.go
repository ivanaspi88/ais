package main

import "fmt"
import "rsc.io/quote"

//import "github.com/ivanaspi88/vlib"
import "github.com/ivanaspi88/vlib/VP1"
import "github.com/ivanaspi88/vlib/VP1/VP2"

//import "github.com/ivanaspi88/dbf"

func main() {

	vp1.Vprint1()
	vp2.Vprint1()
	//vlib.VprintM()

	fmt.Println("AIS inception!") // std package "fmt"

	fmt.Println(quote.Go()) // ext module  "rsc.io/quote"

	//var a int = vlib.Vad(2, 3)
	//fmt.Println("Sum vlib =", a)

}
