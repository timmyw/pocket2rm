package pocket2rm

import (
	"fmt"
	"os"
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
//	"io/ioutil"
)

func sendJSON(req *http.Request, contentType string, res interface{}) error {
	req.Header.Add("X-Accept", "application/json")
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}
	
	// fmt.Printf("REQ\nMETHOD: %s\n", req.Method)
	// fmt.Printf("URL: %s\n", req.URL.String())
	// for h,v := range req.Header {
	// 	fmt.Printf("HDR: %s=%s\n", h, v)
	// }
	
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

// GetJSON sends a GET request to the specified URL and receives JSON
// in response
func GetJSON(url string, origin string, headers map[string]string, res interface{}) error {
	req, err := http.NewRequest("GET", APIOrigin[origin]+url, nil)
	fmt.Printf("URL:%s\n", APIOrigin[origin]+url)
	if err != nil {
		return err
	}

	for h,v := range headers {
		req.Header.Add(h, v)
	}
	
	//return sendJSON(req, "text/html; charset=UTF-8", res)
	return sendJSON(req, "text/plain", res)
}

// PostJSON creates a POST request with the supplied data, and sends
// it using sendJSON.
func PostJSON(url string, origin string, data, res interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// fmt.Printf("PARAMS:%v+\n", data)
	req, err := http.NewRequest("POST", APIOrigin[origin]+url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	return sendJSON(req, "application/json; charset=UTF-8", res)
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

// FixForFileName replaces any illegal characters in the supplied
// string to allow it to be used as a filename.
func FixForFileName(inp string) string {
	var illegals = []string{ "!", "\"", "'", "$", "^", "*" }
	out := inp
	for _, v := range illegals {
		out = strings.Replace(out, v, "_", -1)
	}

	return out
}

