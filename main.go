package main

import (
	"github.com/juju/loggo"
	"github.com/jessevdk/go-flags"
	"github.com/v4lproik/no-name/module"
	"github.com/v4lproik/no-name/util"
	"strconv"
	"github.com/v4lproik/no-name/client"
	"github.com/v4lproik/no-name/data"
	"os"
)

var logger = loggo.GetLogger("main")
var rootDir = ""


type Options struct{
	Ips string `short:"f" long:"filename" description:"File path containing the IPs to scan" required:"true"`
	Selenium string `short:"s" long:"selenium" description:"Url of your standalone/master selenium server"`
	Output string `short:"o" long:"output" description:"Format of the report" required:"true"`
}

func init()  {
	loggo.ConfigureLoggers("debug")

	rootDir, _ = os.Getwd()
}

const STOP_AT_FIRST = true
const LOGIN = "conf/login.txt"
const PASSWORD = "conf/password.txt"
const DEFAULT_PASSWORD = "conf/default-password-web-interface.txt"
const HTML_TAGS_NAMES = "conf/html-detection-tags.txt"

func banner() {
	var banner = `
	|----------------------------------------------------------|
	|              Web Interface Auto Submit 1.3               |
	|                         v4lproik                         |
	|----------------------------------------------------------|
	`

	logger.Infof(banner)
}

func main() {
	// var
	opts := Options{}

	// display banner
	banner()

	// init parser : Pass struct pointer so the init parser can change the data inside the struct
	parser := initParser(&opts)

	// parse cli arguments
	_, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	// parse optsOutput
	reportFormat := data.ReportFormat(0)
	switch opts.Output {
	case "grep":
		reportFormat = data.GREPABLE
	case "html":
		reportFormat = data.HTML
	default:
		panic(parser.Usage)
	}

	// setting up the different objects
	ips, channels, chains := setUp(opts.Ips, opts.Selenium, reportFormat, DEFAULT_PASSWORD,
		PASSWORD, LOGIN, HTML_TAGS_NAMES)

	// launch the chains
	launchChains(ips, channels, chains)

	// wait for all the chains to be finished
	waitForResponse(channels)
}

func launchChains(ips []string, channels []chan string, chains []module.Module) {
	//launch all the first module of all the chains in //
	for idx, chain := range chains {
		channel := channels[idx]
		go func(channel chan string, idx int, chain module.Module) {
			webInterface := data.NewWebInterface(client.NewSimpleWebClient(ips[idx]))

			chain.Request(true, webInterface)
			channel <- webInterface.ClientWeb.GetDomain().String() + " => " + webInterface.ReportPath
		}(channel, idx, chain)
	}
}

func setUp(optsIps string, seleniumServerUrl string, optsOutput data.ReportFormat, defaultPasswordPath string,
	passwordPath string, loginPath string, htmlTagsNamesPath string) ([]string, []chan string, []module.Module) {

	// parse ips to scan
	ips := getIps(optsIps)
	showIps(ips)

	// parse list of credentials + known web interfaces
	credentials := data.NewCredentials(defaultPasswordPath, passwordPath, loginPath)

	//parse html tags' names
	htmlTagsNames := data.NewHtmlSearchValues(htmlTagsNamesPath)

	// create the chains findform -> findid -> bruteforce -> report
	channels := initChannels(len(ips))
	chains := initChains(ips, optsOutput, credentials, seleniumServerUrl, htmlTagsNames)

	return ips, channels, chains
}

func getIps(path string) (ips []string){
	// get
	lines, _ := util.ReadLines(path)

	// filter
	for line := range lines {
		if len(lines[line]) < 1 && lines[line] != "\n" {
			lines = append(lines[:line], lines[line+1:]...)
		}
	}

	return lines
}

func initChains(ips []string, reportFormat data.ReportFormat, credentials *data.Credentials, seleniumServerUrl string,
	htmlTagsNames *data.HtmlSearchValues) ([]module.Module) {
	chains := make([]module.Module, len(ips))

	// init chains
	for key, _ := range ips  {
		firstModule := module.NewScrapModule(seleniumServerUrl)
		secondModule := module.NewFindFormModule(strconv.Itoa(key), htmlTagsNames)
		thirdModule := module.NewFaviconModule(credentials.DefaultWebInterfaces)
		fourthModule := module.NewBruteforceModule(credentials, STOP_AT_FIRST, htmlTagsNames.LoginPatterns)
		fifthModule := module.NewReportModule(rootDir, reportFormat)

		firstModule.SetNextModule(secondModule)
		secondModule.SetNextModule(thirdModule)
		thirdModule.SetNextModule(fourthModule)
		fourthModule.SetNextModule(fifthModule)

		chains[key] = firstModule
	}

	return chains
}

func waitForResponse(channels []chan string) ([]string) {
	reports := make([]string, len(channels))
	for i := 0; i <= len(channels)-1; i++ {
		select {
		case msg := <-channels[i]:
			logger.Debugf(msg)
			reports[i] = msg
		}
	}

	return reports
}

func initChannels(nb int) ([]chan string){
	channels := make([]chan string, nb)

	for i := range channels {
		channels[i] = make(chan string)
	}

	return channels
}

func showIps(ips []string) {
	for key, _ := range ips {
		logger.Infof("- " + ips[key])
	}
}

func initParser(opts *Options) (parser *flags.Parser){
	//default behaviour is HelpFlag | PrintErrors | PassDoubleDash - we need to override the stderr output
	return flags.NewParser(opts, flags.HelpFlag)
}