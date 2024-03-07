package main

import (
	"fmt"
	"strings"

	"github.com/cert-manager/cert-manager/pkg/issuer/acme/dns/util"
)

func main() {
	domain := "www.example.com"
	issuers := []string{"letsencrypt.org"}

	err := util.ValidateCAA(domain, issuers, false, util.RecursiveNameservers);
	if err != nil {
		panic(err)
	}
	fmt.Printf("Non-wildcard certificate for %v can be issued by %v.\n", domain, strings.Join(issuers, ","))
}
