package service

import (
	"github.com/jessevdk/go-flags"
	"fmt"
)

type ArgumentsService struct{
	Ips string `short:"f" long:"filename" description:"File path containing the IPs to scan" required:"true"`
	Favicons string `short:"d" long:"database" description:"File path containing the md5 computation of the web interface's favicon" required:"true"`
	Modules string `short:"m" long:"module" description:"Modules you want to see running against the list of Ips you provided " required:"true"`
}

func NewArgumentsService() *ArgumentsService{
	return &ArgumentsService{}
}

func (a* ArgumentsService) CheckArguments(args []string)  (bool, error) {
	//default behaviour is HelpFlag | PrintErrors | PassDoubleDash - we need to override the stderr output
	parser := flags.NewParser(a, flags.HelpFlag)

	//parse arguments
	_, err := parser.ParseArgs(args)

	if err != nil {
		return false, err
	}

	fmt.Printf("Ips: %s\n", a.Ips)
	fmt.Printf("Favicons: %s\n", a.Favicons)
	fmt.Printf("Modules: %s\n", a.Modules)

	return true, nil
}
