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

func ServeAppCastXML(w http.ResponseWriter, r *http.Request) {
	log.Println("load appcast.xml")
	s := &Sparkle{
		Version: "2.0",
		XMLNSSparkle: "http://www.andymatuschak.org/xml-namespaces/sparkle",
		XMLNSDC: "http://purl.org/dc/elements/1.1/",
		Channels: []Channel {
			Channel{
				Items: []Item {
					Item{
						Title: "Version 0.14.46-1",
						SparkleReleaseNotesLink: "https://xor-gate.github.io/syncthing-macosx/releases/0.14.46-1.html",
						PubDate: "Thu, 19 Apr 2018 21:36:00 GMT+2",
						Enclosure: Enclosure {
							SparkleShortVersionString: "0.14.46-1",
							SparkleVersion: "0144601",
							Type: "application/octet-stream",
							URL: "https://github.com/xor-gate/syncthing-macosx/releases/download/v0.14.46-1/Syncthing-0.14.46-1.dmg",
						},
					},
				},
			},
		},
	}

	w.Write([]byte(xml.Header))
	xw := xml.NewEncoder(w)
	xw.Encode(s)
}

func NewHTTPServer(addr string) (*HTTPServer, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/appcast.xml", ServeAppCastXML)

	srv := &http.Server {
		Addr: addr,
		Handler: mux,
	}

	return &HTTPServer{srv: srv}, nil
}

func (s *HTTPServer) Serve() error {
	return s.srv.ListenAndServe()
}
