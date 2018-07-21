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

	// FAKE newest item
	item := sparkle.Item {
		Title: "Version 0.14.48-1",
		SparkleReleaseNotesLink: "https://xor-gate.github.io/syncthing-macosx",
		PubDate: time.Now().Format(time.RFC1123),
		Enclosure: sparkle.Enclosure {
			SparkleShortVersionString: "0.14.48-1",
			SparkleVersion: "0144801",
			Type: "application/octet-stream",
			URL: "https://github.com/xor-gate/syncthing-macosx/releases/download/v0.14.48-1/Syncthing-0.14.46-1.dmg",
		},
	}
	items = append(items, item)

	for i, release := range a.Releases {
		fmt.Println(fmt.Sprintf("Release #%d:", i+1), release.Version, release.Title, release.PublishedDateTime, release.IsPrerelease)

		item = sparkle.Item {
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
