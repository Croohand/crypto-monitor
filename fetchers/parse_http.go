package fetchers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// JSON format
func parseHttp(r *http.Response, dst interface{}) error {
	defer r.Body.Close()
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("Parse http response: status code %d", r.StatusCode)
	}
	return json.Unmarshal(bytes, dst)
}
