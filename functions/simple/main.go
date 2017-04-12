package main

// Importing dependencies
import (
	"encoding/json"
	"github.com/ChimeraCoder/anaconda"
	"github.com/apex/go-apex"
	"io/ioutil"
	"log"
	"net/url"
	"os"
)

// Creating credentials structure
type credentials struct {
	CtwitterKey         string `json:"twitterKey"`
	CtwitterSecret      string `json:"twitterSecret"`
	CtwitterToken       string `json:"twitterToken"`
	CtwitterTokenSecret string `json:"twitterTokenSecret"`
}

// Hooray Rakslice explained this and now I know why I'm doing it!
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
	anaconda.SetConsumerKey(cred.CtwitterKey)
	anaconda.SetConsumerSecret(cred.CtwitterSecret)
}

// It's the main function! Where it does... stuff! Right now it tweets 'Hello World'. I'm a professional.
func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		api := anaconda.NewTwitterApi(cred.CtwitterToken, cred.CtwitterTokenSecret)
		var tweet string
		tweet = "Hello World"
		log.Printf("Attempting to post tweet: %s", tweet)

		v := url.Values{}
		_, err := api.PostTweet(tweet, v)
		if err != nil {
			log.Printf("Error posting tweet: %s", tweet)
		}
		return tweet, nil
	})
}
