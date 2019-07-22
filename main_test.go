package main

import (
	"fmt"
	"testing"
	"time"
	"app_sample/pkg/util"
	"net/http"
	"net/http/httptest"
)

//TestGetVersion base test case sample
func TestGetVersion(t *testing.T)  {

	version, err := GetVersion()

	if err != nil{
		t.Error(err)
	}

	if version != "v1"{
		t.Errorf("app version not match: %s, expect: %s.", version, "v1")
	}

	fmt.Println(version)
}

//TestLongTimeTestCase skip test
func TestLongTimeTestCase(t *testing.T)  {

	if testing.Short(){
		t.Skip("Skip test when test in short mode.")
	}

	time.Sleep(3 * time.Second)

}

//TestParallel1
func TestParallel1(t *testing.T)  {

	t.Parallel()
	t.Log("test parallel 1")
	time.Sleep(1 * time.Second)

}

//TestParallel2
func TestParallel2(t *testing.T)  {

	t.Parallel()
	t.Log("test parallel 2")
	time.Sleep(1 * time.Second)

}

//BenchmarkGetVersion benchmarking the get version
func BenchmarkGetVersion(b *testing.B) {
	for i := 0; i < b.N; i++{
		GetVersion()
	}
}

//BenchmarkUnmarshalJson bench test unmarshal json file
func BenchmarkUnmarshalJson(b *testing.B) {
	for i :=0; i < b.N; i++{
		util.UnmarshalJson("pkg/util/post_temp.json")
	}
}

//BenchmarkUnmarshalJson bench test decode json file
func BenchmarkDecodeJson(b *testing.B) {
	for i :=0; i < b.N; i++{
		util.DecodeJson("pkg/util/post_temp.json")
	}
}

//TestHandleGetVersion test handle func
func TestHandleGetVersion(t *testing.T) {
	sv := ServerVersion{Version:"v1", IP:"127.0.0.1", Port:8080}

	mux := http.NewServeMux()
	mux.HandleFunc("/version", sv.handGetVersion)

	writer := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/version", nil)
	mux.ServeHTTP(writer, req)

	fmt.Println(writer.Body.String())
}


