package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func loadToken(tokenFile string) string {
	return string(loadFile(tokenFile))
}

func loadTargets(targetFile string) map[string]string {
	var objmap map[string]string //*json.RawMessage
	fileContent := loadFile(targetFile)
	err := json.Unmarshal(fileContent, &objmap)
	if err != nil {
		// fmt.Println("ug")
	}
	return objmap
}

func scrapeAppTransforms() map[string]string {
	result := map[string]string{}

	for _, target := range Configuration.Targets {
		if target.SourceURL != "" {
			if strings.HasPrefix(strings.ToLower(target.SourceURL), "http") {
				var client = &http.Client{}

				req, _ := http.NewRequest("GET", target.SourceURL, nil)

				req.Header.Set("Authorization", "Basic "+Configuration.AccessToken)
				req.Header.Set("Accept", "text/plain")

				r, err := client.Do(req)

				// if unable to connect, mark failed and move on
				if err != nil {
					//return
				}
				// fmt.Println(target.FriendlyName)
				body, err := ioutil.ReadAll(r.Body)
				result[target.FriendlyName] = string(body)
				r.Body.Close()
			} else {
				result[target.FriendlyName] = string(loadFile(target.SourceURL))
			}
		}
	}
	return result
}
