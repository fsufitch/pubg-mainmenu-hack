package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"regexp"
)

// EvilIndexFileName is the filename that the "evil" index is found at
const EvilIndexFileName = "index.html"

var loadScriptRegex = regexp.MustCompile("(?s)<script>(.*)</script>")

type evilIndexHandler struct {
	Wrapped http.Handler
}

func (h evilIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if HSTSRedirect(w, r) {
		return
	}
	h.Wrapped.ServeHTTP(w, r)
}

type evilPUBGLoadData struct {
	RealLoadJS string `json:"real_load_js"`
	APIHost    string `json:"api_host"`
}

// NewEvilIndexHandler returns a handler for returning the evil index.html
func NewEvilIndexHandler(realPUBGURL string, staticDirPath string, apiHost string) (http.Handler, error) {
	log.Printf("Creating fake game loading handler; loading real JS (%s)", realPUBGURL)

	response, err := http.Get(realPUBGURL)
	if err != nil {
		return nil, fmt.Errorf("Error querying real PUBG index file %s: %v", realPUBGURL, err)
	}
	realBodyData, _ := ioutil.ReadAll(response.Body)

	foundScript := loadScriptRegex.FindAllSubmatch(realBodyData, -1)
	if len(foundScript) < 1 || len(foundScript[0]) < 2 {
		return nil, fmt.Errorf("Could not extract real PUBG load data from: %s", string(realBodyData))
	}
	realLoadJS := string(foundScript[0][1])

	log.Printf("Loading evil index: %s", path.Join(staticDirPath, EvilIndexFileName))
	evilIndexData, err := ioutil.ReadFile(path.Join(staticDirPath, EvilIndexFileName))
	if err != nil {
		return nil, fmt.Errorf("Error reading evil index file %s: %v", EvilIndexFileName, err)
	}

	loadData := evilPUBGLoadData{
		RealLoadJS: realLoadJS,
		APIHost:    apiHost,
	}
	loadDataJSON, _ := json.Marshal(loadData)

	evilIndexData = bytes.Replace(evilIndexData, []byte("__LOAD_DATA_JSON__"), []byte(loadDataJSON), -1)

	return evilIndexHandler{NewStaticHandler(evilIndexData)}, nil
}
