package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hey there! I will check the MX, SPF, and DMARC records for you to validate the domain name!")
	fmt.Print("Please enter the domain name (e.g., 'gmail.com'): ")

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
		color.Yellow("uh-oh, seems like you forgot the gTLD (.com) or ccTLD (.ca)!")
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
		color.Red("uh-oh, email address does not exist!")
		os.Exit(3)
	}

	for _, record := range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcString = record
			break
		}
	}

	fmt.Printf("Domain searched: %v\n", domain)
	fmt.Println("Results found:")
	if hasMX{
		fmt.Printf("MX Records: %v\n", color.GreenString("true"))
	} else {
		fmt.Printf("MX Records: %v\n", color.RedString("false"))
	}
	if hasSPF{
		fmt.Printf("SPF Records: %v\t\t%v\n", color.GreenString("true"), spfString)
	} else {
		fmt.Printf("SPF Records: %v\t\t%v\n", color.RedString("false"), spfString)
	}
	if hasDMARC{
		fmt.Printf("DMARC Records: %v\t\t%v\n", color.GreenString("true"), dmarcString)
	} else {
		fmt.Printf("DMARC Records: %v\t\t%v\n", color.RedString("false"), dmarcString)
	}
	fmt.Print("Want to validate another? If so, type it here: ")
}