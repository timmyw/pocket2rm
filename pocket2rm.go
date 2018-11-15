package pocket2rm

import (
	// "fmt"
	"os"
	// "bytes"
	// "encoding/json"
	"github.com/jacobstr/confer"
	// "net/http"
	// "io/ioutil"
	//"net/url"
)

const version = "0.0.1"
var configFile = os.ExpandEnv("$HOME/.config/pocket2rm.yaml")

// Pocket2RM contains the interface to the Pocket2RM API
type Pocket2RM struct {
	version string
	config *confer.Config
	Code string
}

// Init performs any initialisation (such as loading API keys etc).
func (p *Pocket2RM) Init() {
	p.config = confer.NewConfig()
	p.config.ReadPaths(configFile)
}

// PocketCode contains the code returned from the first Pocket auth
// call.
type PocketCode struct {
	Code string `json:"code"`
}

// Authorise carries out an OAUTH call to Pocket to get a token
func (p *Pocket2RM) Authorise() {
	// First step is to get a token from the Pocket servers
	url := "https://getpocket.com/v3/oauth/request"
	params := map[string]interface{}{
		"consumer_key": p.config.GetString("consumer_key"),
		"redirect_uri": "http://localhost",
	}

	result := &PocketCode{}
	err := PostJSON(url, params, result)

	if err != nil {
		panic(err)
	}

	p.Code = result.Code
	
}
