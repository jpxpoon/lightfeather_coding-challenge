package libs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string, result interface{}) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("%w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("%w", err)
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("%w", err)
	}
}
