package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain, hasMX?, hasSPF?, sprRecord?, hasdmarc?, dmarcRecord?, isEmailValid?\n")
	for scanner.Scan() {
		EmailCheck(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("error in input %v \n", err)
	}
}

func EmailCheck(domain string) {
	// GOLang has many standard libraries to check email validity
	// The package used is "net"
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error in MX lookup %v \n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error : %v \n", err)
	}
	for _, record := range txtRecords {
		// we are looking for the records of type "spf1"
		// strings is a package in GOLang which has many functions to manipulate strings
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error in DMARC lookup %v \n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%s, %t, %t, %s, %t, %s, %t\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord, hasMX && hasSPF && hasDMARC)
}
