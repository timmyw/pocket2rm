package pocket2rm

import (
	"os"
	"fmt"
	"github.com/jacobstr/confer"
)

const version = "0.0.1"

// Contains the default configuration file
var configFile = os.ExpandEnv("$HOME/.config/pocket2rm.yaml")

// Pocket2RM contains the interface to the Pocket2RM API
type Pocket2RM struct {
	version	 string
	config	 *confer.Config
	Code	 string
	token	 *PocketToken
}

// Init performs any initialisation (such as loading API keys etc).
func (p *Pocket2RM) Init() {
	p.config = confer.NewConfig()
	p.config.ReadPaths(configFile)
}

// Authorise carries out an OAUTH call to Pocket to get a token
func (p *Pocket2RM) Authorise() {

	code, err := Authorise(p.config.GetString("consumer_key"))

	if err != nil {
		panic(err)
	}

	p.Code = code
}

// GetAccessToken carries out an OAUTH call to Pocket to get a token
func (p *Pocket2RM) GetAccessToken() {

	var err error
	p.token, err = GetToken(p.config.GetString("consumer_key"),
		p.Code)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", *p.token)
}

func PullFromPocket() {

}
