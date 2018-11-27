package pocket2rm

import (
	"fmt"
	"os"
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

	// fmt.Printf("PARAMS:%v+\n", data)
	req, err := http.NewRequest("POST", APIOrigin+url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return sendJSON(req, res)
}

// SaveJSONToFile dumps the supplied struct to the file in JSON format
func SaveJSONToFile(path string, v interface{}) error {
	w, err := os.Create(path)
	if err != nil {
		return err
	}

	defer w.Close()

	return json.NewEncoder(w).Encode(v)
}

// LoadJSONFromFile loads the JSON file into the supplied interface{}
func LoadJSONFromFile(path string, v interface{}) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}

	defer r.Close()

	return json.NewDecoder(r).Decode(v)
}

