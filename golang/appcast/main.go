package main

import (
	//"fmt"

	//"github.com/victorpopkov/go-appcast"
	"github.com/xor-gate/snippets/golang/appcast/sparkle"
)

func main() {
/*
	a := appcast.New()
	a.LoadFromURL("https://github.com/xor-gate/syncthing-macosx/releases.atom")
	a.GenerateChecksum(appcast.SHA256)
	a.ExtractReleases()

	fmt.Println("Checksum:", a.GetChecksum())
	fmt.Println("Provider:", a.GetProvider())

	for i, release := range a.Releases {
		fmt.Println(fmt.Sprintf("Release #%d:", i+1), release.Version, release.Title, release.PublishedDateTime, release.IsPrerelease)
	}

	fmt.Println("Release #1 description:", a.Releases[0].Description)
*/

	srv, _ := sparkle.NewHTTPServer("127.0.0.1:8080")
	srv.Serve()
}
