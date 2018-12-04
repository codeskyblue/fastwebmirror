package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/levigross/grequests"
)

const mirrorURL = "https://testerhome.com"

func hashURL(url string) string {
	m := md5.New()
	m.Write([]byte(url))
	return fmt.Sprintf("%x", m.Sum(nil))
}

func initHandlers() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// io.WriteString(w, r.RequestURI)
		u := url.URL{}
		u.Scheme = "https"
		u.Host = "testerhome.com"
		u.Path = r.URL.Path
		u.RawQuery = r.URL.Query().Encode()
		_ = u
		resp, err := grequests.Get(u.String(), nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Close()
		contentType := resp.Header.Get("Content-Type")
		log.Println("contentType", contentType)
		log.Println(u.String())
		log.Println(r.URL.Path)
		w.Header().Set("Content-Type", contentType)
		if resp.StatusCode == 404 {
			io.Copy(w, resp)
			return
		}

		f, err := os.Create(filepath.Join("cache", hashURL(u.String())))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		mw := io.MultiWriter(w, f)
		io.Copy(mw, resp)
	})
}
func main() {
	// resp, err := grequests.Get("http://httpbin.org/get", nil)
	// if err != nil {
	// 	log.Fatalln("Unable to make requesat: ", err)
	// }
	// log.Println(resp.String())

	os.MkdirAll("cache", 0755)

	initHandlers()
	http.ListenAndServe(":8000", nil)
}
