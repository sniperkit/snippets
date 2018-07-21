package main

import (
	"fmt"
	"time"

	"github.com/victorpopkov/go-appcast"
	"github.com/xor-gate/snippets/golang/appcast/sparkle"
)

func main() {
	a := appcast.New()
	a.LoadFromURL("https://github.com/xor-gate/syncthing-macosx/releases.atom")
	a.GenerateChecksum(appcast.SHA256)
	a.ExtractReleases()

	fmt.Println("Checksum:", a.GetChecksum())
	fmt.Println("Provider:", a.GetProvider())

	var items sparkle.Items

	for i, release := range a.Releases {
		fmt.Println(fmt.Sprintf("Release #%d:", i+1), release.Version, release.Title, release.PublishedDateTime, release.IsPrerelease)

		item := sparkle.Item {
			Title: release.Title,
			PubDate: release.PublishedDateTime.Format(time.RFC1123),
			Enclosure: sparkle.Enclosure{
				SparkleShortVersionString: release.Version.String(),
				Type: "application/octet-stream",
			},
		}
		items = append(items, item)
	}

	srv, _ := sparkle.NewHTTPServer("127.0.0.1:8080", items)
	srv.Serve()
}
