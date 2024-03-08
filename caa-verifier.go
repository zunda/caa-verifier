package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/cert-manager/cert-manager/pkg/issuer/acme/dns/util"
	"k8s.io/klog/v2"
)

func usage() string {
	return fmt.Sprintf("usage: %v <domain>\nverifies CAA recrods.\n", os.Args[0])
}

func init() {
	flag.Usage = func() {
		fmt.Fprint(flag.CommandLine.Output(), usage())
		fmt.Fprint(flag.CommandLine.Output(), "\nCommand line options:\n")
		flag.PrintDefaults()
	}
}

func main() {
	klog.InitFlags(nil)
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Fprint(os.Stderr, usage())
		fmt.Fprint(os.Stderr, "Supply -help option for command line options\n")
		os.Exit(0)
	}
	domain := flag.Arg(0)
	issuers := []string{"letsencrypt.org"}

	err := util.ValidateCAA(domain, issuers, false, util.RecursiveNameservers)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Non-wildcard certificate for %v can be issued by %v.\n", domain, strings.Join(issuers, ","))
}
