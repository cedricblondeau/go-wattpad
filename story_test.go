package wattpad

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshallStories(t *testing.T) {
	data := `{
  "stories": [
    {
      "id": 1,
      "title": "Hello World!",
      "user": "Cédric Blondeau",
      "description": "Lorem ipsum",
      "url": "https://www.cedricblondeau.com"
    },
    {
      "id": 2,
      "title": "The Tale of Scrotie McBoogerballs",
      "user": "Leopold Stotch",
      "description": "Dolor sit amet",
      "url": "http://southpark.cc.com"
    }
  ]}`
	var envelope StoriesEnvelope
	err := json.Unmarshal([]byte(data), &envelope)
	if err != nil {
		t.Errorf("json.Unmarshal StoriesEnvelope returned an error %+v", err)
	}

	stories := envelope.Stories
	assert.Equal(t, 2, len(stories))

	want := []Story{
		Story{
			ID:          1,
			Title:       "Hello World!",
			Author:      "Cédric Blondeau",
			Description: "Lorem ipsum",
			URL:         "https://www.cedricblondeau.com",
		},
		Story{
			ID:          2,
			Title:       "The Tale of Scrotie McBoogerballs",
			Author:      "Leopold Stotch",
			Description: "Dolor sit amet",
			URL:         "http://southpark.cc.com",
		}}
	if !reflect.DeepEqual(stories, want) {
		t.Errorf("StoriesEnvelope returned %+v, want %+v", stories, want)
	}
}
