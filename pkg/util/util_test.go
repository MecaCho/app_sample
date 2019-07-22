package util

import (
	"testing"
)

//TestUnmarshalJson test unmarshal
func TestUnmarshalJson(t *testing.T) {
	post, err := UnmarshalJson("post_temp.json")

	if err != nil{
		t.Error(err)
	}
	t.Logf("Get json file data, name: %s, id: %s", post.Name, post.Id)
}

//BenchmarkUnmarshalJson bench test unmarshal json file
func BenchmarkUnmarshalJson(b *testing.B) {
	for i :=0; i < b.N; i++{
		UnmarshalJson("post_temp.json")
	}
}
