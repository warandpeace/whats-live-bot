package main

import (
  "encoding/json"
  "github.com/apex/go-apex"
  // Commenting out anaconda until I'm actually using the Twitter library
  // "github.com/ChimeraCoder/anaconda"
)

type message struct {
  Hello string `json:"hello"`
}

func main() {
  apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
    var m message

    if err := json.Unmarshal(event, &m); err != nil {
      return nil, err
    }

    return m, nil
  })
}

