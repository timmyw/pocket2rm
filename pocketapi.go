package pocket2rm

import (
	"fmt"
	"net/url"
)

// PocketCode contains the code returned from the first Pocket auth
// call.
type PocketCode struct {
	Code string `json:"code"`
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
