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
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var blogs Blogs

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &blogs)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	return blogs
}
