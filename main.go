package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("hasMX\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}
}

func checkDomain(domain string){
	var hasMX, hasSPF bool

	// checking MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	// deadass why am i so stupid
	for _, record := range txtRecords{
		if strings.HasPrefix(record, "v=spf1"){
			hasSPF = true
			break
		}
	}
}