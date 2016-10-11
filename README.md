# go-wattpad [![Build Status](https://travis-ci.org/cedricblondeau/go-wattpad.svg)](https://travis-ci.org/cedricblondeau/go-wattpad)

> Simple Wattpad API wrapper written in Go. Not meant to be full-featured.

## Usage

Get the package:

```bash
$ go get github.com/cedricblondeau/go-wattpad
```

Example:

```go
import (
	"fmt"

	"github.com/cedricblondeau/go-wattpad"
)

// Build HTTP Client
c := wattpad.NewClient("YOUR_API_KEY_HERE")

// Get a category by tag
fantasy, err := c.Category("fantasy")
// You should check err variable here

// Print category
fmt.Printf("%d - %s\n", fantasy.ID, fantasy.Name)

// Get new stories
stories, err := c.NewStories()
// You should check err variable here

// Print stories
for _, story := range stories {
  fmt.Printf("%d - %s by %s\n", story.ID, story.Title, story.Author)
}
```

## Run tests

```go
go test
```
