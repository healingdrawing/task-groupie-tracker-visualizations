package main

import (
	"flag"
	"fmt"
	"groupie-tracker/server"
	"log"
	"net"
	"net/http"
)

const setBold = "\033[1m"
const reset = "\033[0m"

const HELP = `-------- GROUPIE-TRACKER API Server -------- 

Usage: ` + setBold + `go run . [PORT]` + reset + ` to start API server on the specified port
Example: ` + setBold + `go run .` + reset + ` to start server on default port 8080

Flags:
`

func main() {
	isDebugMode := flag.Bool("debug", false, "debug mode, show path to log call statement")
	flag.Usage = func() {
		fmt.Print(HELP)
		flag.VisitAll(func(f *flag.Flag) {
			fmt.Printf("\t--%v | %v\n", f.Name, f.Usage)
		})
		fmt.Println("\n--------\nMade by Maksim Mikhailov (max@mer.pw) and Sergei Ivanov (kyznector@gmail.com) | grit:lab coding school, Åland Islands, November 2022")
	}
	flag.Parse()
	if *isDebugMode {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}
	address := ":8080"
	if flag.NArg() == 1 {
		address = ":" + flag.Arg(0)
	}

	router := server.Start()

	l, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started successfully on http://%v/\n", l.Addr())

	err = http.Serve(l, router)
	if err != nil {
		log.Fatal(err)
	}
}
