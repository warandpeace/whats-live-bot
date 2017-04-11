package main

// Importing dependencies
import (
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
	"github.com/apex/go-apex"
	"io/ioutil"
	"log"
	"os"
)

// Creating credentials structure
type credentials struct {
	CtwitterKey         string `json:"twitterKey"`
	CtwitterSecret      string `json:"twitterSecret"`
	CtwitterToken       string `json:"twitterToken"`
	CtwitterTokenSecret string `json:"twitterTokenSecret"`
}

// Creating message structure (Note: From Apex Golang Example)
// TODO: Remove this at some point
type message struct {
	Hello string `json:"hello"`
}

// Hooray rakslice explained this and now I know why I'm doing it!
// It's because I've defined credentials above but I still need to _declare_ an instance of it
var cred credentials

// Using the init function to read in my totally insecure credentials json file
// and set all my credentials for API calls in main()
func init() {
	file, err := ioutil.ReadFile("./credentials.json")
	if err != nil {
		log.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &cred)
	// Set Anaconda credentials
	// TODO: Verify these are being set properly with the values from the local json file
	anaconda.SetConsumerKey(cred.CtwitterKey)
	anaconda.SetConsumerSecret(cred.CtwitterSecret)
	// api := anaconda.NewTwitterApi(cred.CtwitterToken, cred.CtwitterTokenSecret)
}

// It's the main function! Where it does... stuff! (Note: From Apex Golang Example)
// TODO: Fucking everything. Replace this at bare minimum with a hello world tweet to get started.
func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		return m, nil
	})
}
