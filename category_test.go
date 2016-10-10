package wattpad

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshallCategories(t *testing.T) {
	data := `{
  "categories": [
    {
      "id": 4,
      "name": "Romance"
    },
    {
      "id": 5,
      "name": "Science Fiction"
    }
  ]}`
	var envelope CategoriesEnvelope
	err := json.Unmarshal([]byte(data), &envelope)
	if err != nil {
		t.Errorf("json.Unmarshal CategoriesEnvelope returned an error %+v", err)
	}

	categories := envelope.Categories
	assert.Equal(t, 2, len(categories))

	want := []Category{{ID: 4, Name: "Romance"}, Category{ID: 5, Name: "Science Fiction"}}
	if !reflect.DeepEqual(categories, want) {
		t.Errorf("CategoriesEnvelope returned %+v, want %+v", categories, want)
	}
}
