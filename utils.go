package pocket2rm

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
)

func sendJSON(req *http.Request, res interface{}) error {
	req.Header.Add("X-Accept", "application/json")
	req.Header.Add("Content-Type", "application/json; charset=UTF-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("got response %d; X-Error=[%s]", resp.StatusCode, resp.Header.Get("X-Error"))
	}

	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(res)	
}

// PostJSON creates a POST request with the supplied data, and sends
// it using sendJSON.
func PostJSON(url string, data, res interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return sendJSON(req, res)
}

