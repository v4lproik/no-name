package main

import (
	"github.com/juju/loggo"
	"github.com/jessevdk/go-flags"
	"github.com/v4lproik/no-name/module"
	"github.com/v4lproik/no-name/util"
	"strconv"
	"strings"
	"github.com/v4lproik/no-name/client"
	"github.com/v4lproik/no-name/data"
	"os"
)

var logger = loggo.GetLogger("main")
var rootDir = ""


type Options struct{
	Ips string `short:"f" long:"filename" description:"File path containing the IPs to scan" required:"true"`
	Favicons string `short:"d" long:"database" description:"File path containing the md5 computation of the web interface's favicon" required:"true"`
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
	ips, channels, chains := setUp(opts.Favicons, opts.Ips, reportFormat, DEFAULT_PASSWORD,
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
			channel <- webInterface.ReportPath
		}(channel, idx, chain)
	}
}

func setUp(optsFavicon string, optsIps string, optsOutput data.ReportFormat, defaultPasswordPath string,
	passwordPath string, loginPath string, htmlTagsNamesPath string) ([]string, []chan string, []module.Module) {

	// parse favicons database
	favicons := getFavicons(optsFavicon)
	showFavicons(favicons)

	// parse ips to scan
	ips := getIps(optsIps)
	showIps(ips)

	// parse default password database
	credentials := data.NewCredentials(defaultPasswordPath, passwordPath, loginPath)

	//parse html tags' names
	htmlTagsNames := data.NewHtmlSearchValues(htmlTagsNamesPath)

	// create the chains findform -> findid -> bruteforce -> report
	channels := initChannels(len(ips))
	chains := initChains(ips, optsOutput, credentials, htmlTagsNames)

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

func getFavicons(filePath string) (map[string]string) {
	linesFav, _ := util.ReadLines(filePath)

	favicons := make(map[string]string)
	for line := range linesFav {
		tmp := strings.Split(linesFav[line], ":")

		if len(tmp) > 1 {
			favicons[tmp[0]] = tmp[1]
		}else{
			logger.Warningf("Can't process line : <" + linesFav[line] + ">")
		}
	}

	return favicons
}

func initChains(ips []string, reportFormat data.ReportFormat, credentials *data.Credentials, htmlTagsNames *data.HtmlSearchValues) ([]module.Module) {
	chains := make([]module.Module, len(ips))

	// init chains
	for key, _ := range ips  {
		firstModule := module.NewScrapModule()
		secondModule := module.NewFindFormModule(strconv.Itoa(key), htmlTagsNames)
		thirdModule := module.NewFaviconModule(credentials)
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

func showFavicons(favicons map[string]string) {
	for key, value := range favicons {
		logger.Infof("key " + key + " value " + value)
	}
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