package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cert-manager/cert-manager/pkg/issuer/acme/dns/util"
)

func main() {
	if len(os.Args) < 2{
		fmt.Fprintf(os.Stderr, "usage: %v <domain>\nverifies CAA recrods\n", os.Args[0])
		os.Exit(1)
	}
	domain := os.Args[1]
	issuers := []string{"letsencrypt.org"}

	err := util.ValidateCAA(domain, issuers, false, util.RecursiveNameservers);
	if err != nil {
		panic(err)
	}
	fmt.Printf("Non-wildcard certificate for %v can be issued by %v.\n", domain, strings.Join(issuers, ","))
}
