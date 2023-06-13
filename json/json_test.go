package json_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stu struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Other string
}

func TestJson(t *testing.T) {
	s := stu{
		Name:  "zhangsan",
		Other: "other",
	}
	file, err := ioutil.ReadFile("./data.json")
	assert.Nil(t, err)
	json.Unmarshal(file, &s)
	t.Log(s)
}
