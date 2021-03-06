package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"

	p2rm "github.com/timmyw/pocket2rm"
)

func loadToken(p *p2rm.Pocket2RM) {
	var token = p.AccessToken
	if token == nil {
		// The user needs to (re)auth us with pocket
		authWithPocket(p)
	}
}

func authWithPocket(p *p2rm.Pocket2RM) {
	// Load the consumer code
	p.Init()
	// Get an initial request token
	p.GetRequestToken()

	ch := make(chan struct{})
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/favicon.ico" {
				http.Error(w, "Not Found", 404)
				return
			}

			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintln(w, "Authorized.")
			ch <- struct{}{}
		}))
	defer ts.Close()

	url := p2rm.GenerateAuthURL(p.RequestToken.Code, ts.URL)
	fmt.Println(url)

	<-ch

	p.GetAccessToken()
}

func main() {

	command := flag.String("command", "pull", "Command to execute")
	count := flag.Int("count", 10, "Maximum number of articles to pull")
	
	flag.Parse()
	fmt.Printf("%s\n", *command)

	var p = new(p2rm.Pocket2RM)
	p.Init()

	loadToken(p)

	switch *command {
	case "pull":
		p.PullFromPocket(*count)
	case "rmls":
		p.ListFromRemarkable()
		
	case "auth":
		//p.Authorise()
		//p.GetAccessToken()
	}
}
