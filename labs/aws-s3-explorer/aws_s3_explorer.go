package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Data ...
type Data struct {
	BucketName string
	Objects    map[string]bool
	Dirs       map[string]bool
	Extentions map[string]int
}

//Content ...
type Content struct {
	Key string `xml:"Key"`
}

//ListBucket ...
type ListBucket struct {
	XMLName  xml.Name  `xml:"ListBucketResult"`
	Contents []Content `xml:"Contents"`
}

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		panic("Not enough arguments")
	}

	if args[0] != "--bucket" {
		panic("Not the right argument!")
	}

	bucketName := args[1]
	projectData := Data{BucketName: bucketName, Dirs: make(map[string]bool), Objects: make(map[string]bool), Extentions: make(map[string]int)}

	/* Get XML from URL */
	resp, err := http.Get("https://" + bucketName + ".s3.amazonaws.com")

	if err != nil {
		panic("Get: " + err.Error())
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("Read body: " + err.Error())
	}

	/* Process XML Response into an object */

	var ListBucket ListBucket
	xml.Unmarshal(data, &ListBucket)

	/* Process Object Data */

	for _, Content := range ListBucket.Contents {
		key := Content.Key

		if strings.Contains(key, ".") {
			if _, exists := projectData.Objects[key]; !exists {
				projectData.Objects[key] = true
			}

			container := strings.Split(key, ".")

			ext := container[len(container)-1]
			_, exists := projectData.Extentions[ext]

			if !exists {
				projectData.Extentions[ext] = 0
				projectData.Extentions[ext]++
			} else {
				projectData.Extentions[ext]++
			}

		}

		if strings.HasSuffix(key, "/") {

			if !projectData.Dirs[key] {
				projectData.Dirs[key] = true
			}
		}

	}

	print(projectData)
}

func print(data Data) {
	fmt.Println("AWS S3 Explorer")
	fmt.Println("Bucket Name\t\t: " + data.BucketName)
	fmt.Println("Number of objects\t: " + strconv.Itoa(len(data.Objects)))
	fmt.Println("Number of directories\t: " + strconv.Itoa(len(data.Dirs)))
	fmt.Print("Extensions\t\t: ")
	for key, value := range data.Extentions {
		fmt.Print(key + "(" + strconv.Itoa(value) + ")")
		fmt.Print(", ")
	}
}
