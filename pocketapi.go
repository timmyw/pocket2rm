package pocket2rm

import (
	"fmt"
//	"time"
	"net/url"
)

// PocketCode contains the code returned from the first Pocket auth
// call.
type PocketCode struct {
	Code string `json:"code"`
}

// ItemStatus is a typedef wrapping the status of an item
type ItemStatus int

// Item contains all the data returned 
type Item struct {
	ItemID        int        `json:"item_id,string"`
	ResolvedID    int        `json:"resolved_id,string"`
	GivenURL      string     `json:"given_url"`
	ResolvedURL   string     `json:"resolved_url"`
	GivenTitle    string     `json:"given_title"`
	ResolvedTitle string     `json:"resolved_title"`
	Favorite      int        `json:",string"`
	Status        ItemStatus `json:",string"`
	Excerpt       string
	IsArticle     int                 `json:"is_article,string"`
	// HasImage      ItemMediaAttachment `json:"has_image,string"`
	// HasVideo      ItemMediaAttachment `json:"has_video,string"`
	WordCount     int                 `json:"word_count,string"`

	// Fields for detailed response
	Tags    map[string]map[string]interface{}
	Authors map[string]map[string]interface{}
	Images  map[string]map[string]interface{}
	Videos  map[string]map[string]interface{}

	// Fields that are not documented but exist
	// SortID        int       `json:"sort_id"`
	// TimeAdded     time.Time `json:"time_added"`
	// TimeUpdated   time.Time `json:"time_updated"`
	// TimeRead      time.Time `json:"time_read"`
	// TimeFavorited time.Time `json:"time_favorited"`
}

// RetrieveResult contains the response from a GET request
type RetrieveResult struct {
	List     map[string]Item
	Status   int
	Complete int
	Since    int
}

// APIOrigin contains the destination
const APIOrigin = "https://getpocket.com"

// Authorise carries out an auth call to Pocket to get a token
func Authorise(consumerKey string) (string, error) {
	// First step is to get a token from the Pocket servers
	url := "/v3/oauth/request"
	params := map[string]interface{}{
		"consumer_key": consumerKey,
		"redirect_uri": "http://localhost",
	}

	result := &PocketCode{}
	err := PostJSON(url, params, result)

	if err != nil {
		return "", err
	}

	return result.Code, nil
	
}

// AccessToken stores the returned access token and user name from an
// access call
type AccessToken struct {
	Token		 string `json:"access_token"`
	Username	 string `json:"username"`
}

// RequestToken stores the initial request token
type RequestToken struct {
	Code string `json:"code"`
}

// GetRequestToken will retrieve a RequestToken from Pocket
func GetRequestToken(consumerKey string) (*RequestToken, error) {
	result := &RequestToken{}	
	err := PostJSON("/v3/oauth/request",
		map[string]string {
			"consumer_key"	: consumerKey,
			"redirect_uri"  : "http://localhost",
		},
		result)

	if err != nil {
		return nil, err
	}

	return result, nil

}

// GetAccessToken will retrieve an access token from Pocket
func GetAccessToken(consumerKey string, requestToken *RequestToken) (*AccessToken, error) {
	result := &AccessToken{}
	err := PostJSON("/v3/oauth/authorize",
		map[string]string {
			"consumer_key"	: consumerKey,
			"code"		: requestToken.Code,
		},
		result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GenerateAuthURL will return the URL to redirect the user to in
// order to authorise the app with Pocket
func GenerateAuthURL(code string, redirect string) string {
	params := url.Values{ "request_token": {code}, "redirect_uri": {redirect}}
	return fmt.Sprintf("%s/auth/authorize?%s", APIOrigin, params.Encode())
}

// PullArticles retrieves a list of articles
func
PullArticles(consumerKey string, accessToken *AccessToken, count int) (*RetrieveResult, error) {
	result := &RetrieveResult{}
	err := PostJSON("/v3/get",
		map[string]string {
			"consumer_key"	: consumerKey,
			"access_token"	: accessToken.Token,
			"detailType"    : "simple",
			"count"         : fmt.Sprintf("%d", count),
		},
		result)

	if err != nil {
	 	return nil, err
	}

	return result, nil
}
