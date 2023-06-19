package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("hasMX\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}
}

func checkDomain(domain string){
	var hasMX bool

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		panic(err)
	}
}