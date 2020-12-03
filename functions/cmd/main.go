package main

import "github.com/cmorales95/golang_api/functions"

func execute(name string, f functions.MyFunc) {
	f(name)
}

func main() {
	// first function order
	name := "Cristian"
	execute(name, functions.MiddlewareLog(functions.Greeting))
	execute(name,functions.MiddlewareLog(functions.SayBye))
}


