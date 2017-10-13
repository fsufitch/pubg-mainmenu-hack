package resources

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/NYTimes/gziphandler"
	"github.com/fsufitch/pubg-mainmenu-hack/server/handlers"
	"github.com/gorilla/mux"
)

// RegisterResourcePaths registers paths for static resources
func RegisterResourcePaths(router *mux.Router, staticResourceDir string) error {
	paths, err := getStaticFiles(staticResourceDir)
	if err != nil {
		return err
	}

	for _, path := range paths {
		if path == handlers.EvilIndexFileName {
			router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Location", "/")
				w.WriteHeader(http.StatusPermanentRedirect)
			})
			continue
		}
		registerStaticFile(router, path, staticResourceDir)
	}

	return nil
}

func getStaticFiles(basePath string) ([]string, error) {
	fds, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, fmt.Errorf("error listing static files: %v", err)
	}
	files := []string{}
	for _, fd := range fds {
		if fd.IsDir() {
			continue // TODO: implement recursive static files?
		}
		files = append(files, fd.Name())
	}
	return files, nil
}

func registerStaticFile(router *mux.Router, fileName string, basePath string) error {
	staticData, err := ioutil.ReadFile(path.Join(basePath, fileName))
	if err != nil {
		return fmt.Errorf("error reading static file %s: %v", fileName, err)
	}
	h := handlers.NewStaticHandler(staticData)
	log.Printf("Registering static file: %s", fileName)
	router.Handle("/"+fileName, gziphandler.GzipHandler(h))
	return nil
}
