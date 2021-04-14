package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Blogs ...
type Blogs struct {
	Blogs []Blog `json:"blogs"`
}

// Blog ...
type Blog struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

// GetBlogs ...
func GetBlogs() Blogs {
	jsonFile, err := os.Open("./client/blogs.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var blogs Blogs
	json.Unmarshal(byteValue, &blogs)
	defer jsonFile.Close()
	return blogs
}
