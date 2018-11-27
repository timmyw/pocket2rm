package pocket2rm

import (
	"os"
	"fmt"
	"database/sql"
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

	db               *sql.DB
}

// Init performs any initialisation (such as loading API keys etc).
func (p *Pocket2RM) Init() {
	if p.init {
		return
	}
	p.Config = confer.NewConfig()
	p.Config.ReadPaths(ConfigFile)
	p.ConsumerKey = p.Config.GetString("consumer_key")

	p.openDatabase()
	
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
func (p *Pocket2RM) PullFromPocket(count int) {
	pocketArticles, err := PullArticles(p.ConsumerKey, p.AccessToken, count)
	if err != nil {
		panic(err)
	}

	var (
		newArticles []string
		removedArticles []string
	)

	// For each article from Pocket, check to see if it is new
	for k := range pocketArticles.List {
		//fmt.Printf("key:%s\n", k)
		itemTime := p.isArticleKnown(k)
		//fmt.Printf("%v:%v\n", itemTime, NullTime)
		if itemTime == NullTime {
			// Need to add this one
			newArticles = append(newArticles, k)
			//fmt.Printf("%v\n", newArticles)
		}
	}

	// For each article we already know about, check to see if it
	// is still in the pocket list (otherwise it should be
	// removed)
	existingArticles, _ := p.listAllArticles()
	for _,v := range existingArticles {
		var found = false
		for j := range pocketArticles.List {
			if j == v {
				found = true
				break
			}
		}

		if !found {
			removedArticles = append(removedArticles, v)
		}
	}

	for _, itemID := range newArticles {
		p.AddArticle(itemID, pocketArticles.List)
	}
}

// AddArticle will retrieve the HTML, generate a PDF, and upload it to
// RM
func (p *Pocket2RM) AddArticle(itemID string, pocketArticles map[string]Item) {

	item := pocketArticles[itemID]
	url := item.ResolvedURL

	fmt.Printf("%s\n", url)

	var ad = &ArticleDetails{}
	
	err := p.GetArticleDetails(url, ad)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Author:%s\n", ad.Author)
}
