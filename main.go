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
	fmt.Println("domain, hasMX, hasSPF, spfString, hasDMARC, dmarcString")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}
}

func checkDomain(domain string){
	var hasMX, hasSPF, hasDMARC bool
	var dmarcString, spfString string

	// checking MX records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if len(mxRecords) > 0 {
		hasMX = true
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
			spfString = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}

	for _, record := range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcString = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, spfString, hasDMARC, dmarcString)
}