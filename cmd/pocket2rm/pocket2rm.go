package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	
	p2rm "github.com/timmyw/pocket2rm"
)

func loadToken(p *p2rm.Pocket2RM) {
	var token = p.Token
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
	
	redirectURI := p2rm.GenerateAuthURL(p.Code)
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

	url := p2rm.GenerateAuthURL(ts.URL)

	<- ch

	
}

func main() {

	command := flag.String("command", "pull", "Command to execute")
	flag.Parse()
	fmt.Printf("%s\n", *command)

	var p *p2rm.Pocket2RM = new(p2rm.Pocket2RM)
	p.Init()

	loadToken(p)
	
	switch *command {
	case "url":
		p.Authorise()
		fmt.Println(p2rm.GenerateAuthURL(p.Code))
	case "pull":
	case "auth":
		p.Authorise()
		p.GetAccessToken()
	}
}
