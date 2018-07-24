package sparkle

import (
	"log"
	"net/http"
	"encoding/xml"
)

type HTTPServer struct {
	mux *http.ServeMux
	srv *http.Server
}

type appCastXMLHandler struct {
	Items
}

type appCastAssetHandler struct {
}

func (acah *appCastAssetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("dl", r.URL.Path)
	http.Redirect(w, r, "https://github.com/xor-gate/syncthing-macosx/releases/download/v0.14.46-1/Syncthing-0.14.46-1.dmg", http.StatusMovedPermanently)
}

func (ach *appCastXMLHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("load appcast.xml")
	s := &Sparkle{
		Version: "2.0",
		XMLNSSparkle: "http://www.andymatuschak.org/xml-namespaces/sparkle",
		XMLNSDC: "http://purl.org/dc/elements/1.1/",
		Channels: []Channel {
			Channel{
				Title: "Synthing for Mac OS X Changelog",
				Link: "https://xor-gate.github.io/syncthing-macosx/appcast.xml",
				Description: "Most recent changes with links to updates.",
				Language: "en",
				Items: ach.Items,
			},
		},
	}

	w.Write([]byte(xml.Header))
	xw := xml.NewEncoder(w)
	log.Println(xw.Encode(s))
}

func NewHTTPServer(addr string, items Items) (*HTTPServer, error) {
	mux := http.NewServeMux()

	mux.Handle("/appcast.xml", &appCastXMLHandler{Items: items})
	mux.Handle("/dl/", &appCastAssetHandler{})

	srv := &http.Server {
		Addr: addr,
		Handler: mux,
	}

	return &HTTPServer{srv: srv}, nil
}

func (s *HTTPServer) Serve() error {
	return s.srv.ListenAndServe()
}
