package main

import (
	"os"
	"fmt"
	"time"
	"encoding/xml"
	"strings"
	"strconv"
	"context"

	"github.com/hashicorp/go-version"
	"github.com/google/go-github/github"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/xor-gate/snippets/golang/appcast/sparkle"
)

func MarkdownToHTML(md string) string {
	return string(blackfriday.Run([]byte(md)))
}

// LastVersion : Check last version of package
func Releases() ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	ctx := context.Background()
	releases, _, err := client.Repositories.ListReleases(ctx, "xor-gate", "syncthing-macosx", nil)
	return releases, err
}

func main() {
	var items sparkle.Items

	// FAKE newest item
	item := sparkle.Item {
		Title: "Version 0.14.48-1",
		PubDate: time.Now().Format(time.RFC1123),
		Description: sparkle.CdataString{Value: "Die shit is los"},
		Enclosure: sparkle.Enclosure {
			SparkleShortVersionString: "0.14.48-1",
			SparkleVersion: "144801",
			Type: "application/octet-stream",
			URL: "http://localhost:8080/xor-gate/syncthing-macosx/releases/download/v0.14.46-1/Syncthing-0.14.46-1.dmg",
		},
	}
	items = append(items, item)

	releases, _ := Releases()
	for _, release := range releases {
		// Decode git tag into sparkleVersion for CFBundleVersion check
		// "v0.14.48-1" -> "144801"
		rTag := release.GetTagName()
		rVersion, _ := version.NewVersion(rTag)
		rSegments := rVersion.Segments()
		if len(rSegments) != 3 {
			continue
		}

		if rTag[0] != 'v' {
			continue
		}

		dmgVersion := rTag[1:]

		distVersion, err := strconv.ParseUint(rVersion.Prerelease(), 10, 8)
		if err != nil {
			continue
		}
		sparkleVersion := fmt.Sprintf("%02d%02d%02d", rSegments[1], rSegments[2], distVersion)

		var dmgAssetURL string

		for _, asset := range release.Assets {
			url := asset.GetBrowserDownloadURL()
			if !strings.HasSuffix(url, ".dmg") {
				continue
			}
			dmgAssetURL = fmt.Sprintf("https://github.com/xor-gate/syncthing-macosx/releases/download/%s/Syncthing-%s.dmg", rTag, dmgVersion)
		}

		if dmgAssetURL == "" {
			continue
		}

		htmlDescription := MarkdownToHTML(release.GetBody())

		item := sparkle.Item {
			Title: release.GetName(),
			PubDate: release.PublishedAt.Format(time.RFC1123),
			Description: sparkle.CdataString{Value: htmlDescription},
			Enclosure: sparkle.Enclosure{
				SparkleShortVersionString: rTag,
				SparkleVersion: sparkleVersion,
				URL: dmgAssetURL,
				Type: "application/octet-stream",
			},
		}
		items = append(items, item)
	}

	s := &sparkle.Sparkle{
		Version: "2.0",
		XMLNSSparkle: "http://www.andymatuschak.org/xml-namespaces/sparkle",
		XMLNSDC: "http://purl.org/dc/elements/1.1/",
		Channels: []sparkle.Channel {
			sparkle.Channel{
				Title: "Synthing for Mac OS X Changelog",
				Link: "https://xor-gate.github.io/syncthing-macosx/appcast.xml",
				Description: "Most recent changes with links to updates.",
				Language: "en",
				Items: items,
			},
		},
	}

	os.Stdout.Write([]byte(xml.Header))
	xw := xml.NewEncoder(os.Stdout)
	xw.Encode(s)
}
