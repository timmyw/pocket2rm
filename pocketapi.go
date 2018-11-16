package pocket2rm

import (
	"fmt"
)

// PocketCode contains the code returned from the first Pocket auth
// call.
type PocketCode struct {
	Code string `json:"code"`
}

// Authorise carries out an auth call to Pocket to get a token
func Authorise(consumerKey string) (string, error) {
	// First step is to get a token from the Pocket servers
	url := "https://getpocket.com/v3/oauth/request"
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

// PocketToken stores the returned access token and user name from an
// access call
type PocketToken struct {
	Token		 string `json:"access_token"`
	Username	 string `json:"username"`
}

// GetToken will retrieve an access token from Pocket
func GetToken(consumerKey string, code string) (*PocketToken, error) {
	result := &PocketToken{}
	err := PostJSON("/v3/oauth/authorize",
		map[string]string {
			"consumer_key"	: consumerKey,
			"code"		: code,
		},
		result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

// GenerateAuthURL will return the URL to redirect the user to in
// order to authorise the app with Pocket
func GenerateAuthURL(code string) string {
	return fmt.Sprintf("https://getpocket.com/auth/authorize?redirect_uri=http://localhost&request_token=%s", code)
}
