package main

import "fmt"
import "net/http"
import "io/ioutil"
import "log"
import "encoding/json"
import "strings"
import "sort"
import "flag"

type ImageInfos struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func main() {
	var registry, imagePath, tagPrefix string
	var latestOnly bool
	flag.StringVar(&registry, "r", "docker-reg.example.org:5000", "host + port (optional) to docker-registry")
	flag.StringVar(&imagePath, "i", "project/app", "image-path on registry")
	flag.StringVar(&tagPrefix, "p", "SNAPSHOT", "tag-prefix")
	flag.BoolVar(&latestOnly, "l", false, "print only latest tag-value")

	flag.Parse()

	resp, err := http.Get(fmt.Sprintf("https://%v/v2/%v/tags/list", registry, imagePath))
	if err != nil {
		log.Fatal(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var imageInfos ImageInfos
		json.Unmarshal([]byte(body), &imageInfos)

		var snapshotTags []string
		if !latestOnly {
			fmt.Printf("Name: %v\n", imageInfos.Name)
		}
		for _, tag := range imageInfos.Tags {
			if !latestOnly {
				fmt.Printf("Tag: %v\n", tag)
			}
			if strings.HasPrefix(tag, tagPrefix) {
				snapshotTags = append(snapshotTags, tag)
			}
		}
		sort.Strings(snapshotTags)
		latest := snapshotTags[len(snapshotTags)-1]
		if !latestOnly {
			fmt.Println("----")
			fmt.Printf("latest: %v\n", latest)
		} else {
			fmt.Println(latest)
		}
	}
}
