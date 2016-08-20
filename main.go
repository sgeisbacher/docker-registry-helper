package main

import "fmt"
import "net/http"
import "io/ioutil"
import "log"
import "encoding/json"
import "strings"
import "sort"

type ImageInfos struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func main() {
	resp, err := http.Get("https://docker-reg-01.local.netconomy.net:5000/v2/xis/hybris/tags/list")
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
		fmt.Printf("Name: %v\n", imageInfos.Name)
		for _, tag := range imageInfos.Tags {
			fmt.Printf("Tag: %v\n", tag)
			if strings.HasPrefix(tag, "SNAPSHOT") {
				snapshotTags = append(snapshotTags, tag)
			}
		}
		sort.Strings(snapshotTags)
		fmt.Println("----")
		fmt.Printf("latest: %v\n", snapshotTags[len(snapshotTags)-1])
	}
}
