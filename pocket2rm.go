package pocket2rm

import (
	"os"
	"fmt"
	"github.com/jacobstr/confer"
)

const version = "0.0.1"

// ConfigFile contains the default configuration file
var ConfigFile = os.ExpandEnv("$HOME/.config/pocket2rm.yaml")

// AccessFile contains the access token file
var AccessFile = os.ExpandEnv("$HOME/.config/pocket2rm.access.json")

// DatastoreFile contains the path to the SQLite3 database
var DatastoreFile = os.ExpandEnv("$HOME/.config/pocket2rm.db")

// Pocket2RM contains the interface to the Pocket2RM API
type Pocket2RM struct {
	version		 string
	Config		 *confer.Config
	ConsumerKey	 string
	RequestToken	 *RequestToken
	AccessToken      *AccessToken
	init             bool
}

// Init performs any initialisation (such as loading API keys etc).
func (p *Pocket2RM) Init() {
	if p.init {
		return
	}
	p.Config = confer.NewConfig()
	p.Config.ReadPaths(ConfigFile)
	p.ConsumerKey = p.Config.GetString("consumer_key")

	p.AccessToken = nil
	accessToken := &AccessToken{}
	err := LoadJSONFromFile(AccessFile, accessToken)
	if err == nil {
		p.AccessToken = accessToken
	}
	
	p.init = true
}

// GetRequestToken calls into the API to get an initial request token
func (p *Pocket2RM) GetRequestToken() {
	var err error
	p.RequestToken, err = GetRequestToken(p.ConsumerKey)
	if (err != nil) {
		panic(err)
	}

	fmt.Printf("%+v\n", *p.RequestToken)
}

// GetAccessToken carries out an OAUTH call to Pocket to get a token
func (p *Pocket2RM) GetAccessToken() {

	var err error
	p.AccessToken, err = GetAccessToken(p.Config.GetString("consumer_key"),
		p.RequestToken)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", *p.AccessToken)

	// Store the token for next time
	SaveJSONToFile(AccessFile, p.AccessToken)
}

// PullFromPocket retrieves a list of the articles from pocket and
// compares against local state.
func (p *Pocket2RM) PullFromPocket() {
	result, err := PullArticles(p.ConsumerKey, p.AccessToken, 1)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", result)
}
