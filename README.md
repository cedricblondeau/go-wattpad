# go-wattpad [![Build Status](https://travis-ci.org/cedricblondeau/go-wattpad.svg)](https://travis-ci.org/cedricblondeau/go-wattpad)

> Simple and definitely not full-featured Wattpad API wrapper written in Go.

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

// Get a category by tag
c := wattpad.NewClient("YOUR_API_KEY_HERE")
fantasy, err := c.Category("fantasy")
// You should check err variable here
fmt.Printf("%d - %s\n", fantasy.ID, fantasy.Name)

// Get new stories
stories, err := c.NewStories()
// You should check err variable here
for _, story := range stories {
  fmt.Printf("%d - %s by %s\n", story.ID, story.Title, story.Author)
}
```

## Run tests

```go
go test
```
