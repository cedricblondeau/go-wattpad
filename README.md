# go-wattpad

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
fantasy, _ := c.Category("fantasy")
// TODO: Deal with err
fmt.Printf("%d - %s\n", fantasy.ID, fantasy.Name)

// Get new stories
stories, _ := c.NewStories()
// TODO: Deal with err
for _, story := range stories {
  fmt.Printf("%d - %s by %s\n", story.ID, story.Title, story.Author)
}
```

## Run tests

```go
go test
```