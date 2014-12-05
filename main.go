package main

import (
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"net/url"
	"os"
	"strings"
)

var helpText = `Open a DuckDuckGo search from the command line.

Usage:

  duck [query]

Examples:

  # Default DuckDuckGo search
  duck nikola tesla

  # The '!bang' commands need to be escaped
  duck nikola tesla \!g
  duck duckduckgo \!gh
  duck synonym \!t
  duck chunky bacon \!wiki
  duck pie town \!gmap
  duck space cats \!i
  duck hasselhof antivirus \!yt

For all possible "!bang" commands see: https://duckduckgo.com/bang.html`

func main() {
	flag.Parse()
	if len(flag.Args()) >= 1 {
		query := url.QueryEscape(strings.Join(flag.Args(), " "))
		url := fmt.Sprintf("https://duckduckgo.com/?q=%s", query)
		open.Run(url)
	} else {
		fmt.Fprintf(os.Stderr, "%s\n\n", helpText)
	}
}
