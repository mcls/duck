package main

import (
	"flag"
	"fmt"
	"github.com/skratchdot/open-golang/open"
	"strings"
)

func main() {
	flag.Parse()
	query := strings.Join(flag.Args(), " ")
	url := fmt.Sprintf("https://duckduckgo.com/?q=%s", query)
	open.Run(url)
}
