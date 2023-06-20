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

	// checking mxRecords
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}

}