package util

import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
)

type Post struct {
	Name string `json:"name"`
	Id string `json:"id"`

}


func UnmarshalJson(filename string) (post Post, err error)  {

	jsonFile, err := os.Open(filename)
	if err != nil{
		log.Println("Open json file failed.")
		return
	}
	defer jsonFile.Close()

	dataJson, err := ioutil.ReadAll(jsonFile)
	if err != nil{
		log.Println("Read json file failed.")
		return
	}
	json.Unmarshal(dataJson, &post)

	return

}

func DecodeJson(filename string) (post Post, err error) {

	jsonFile, err := os.Open(filename)
	if err != nil{
		log.Println("Open json file failed.")
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&post)
	if err != nil{
		log.Println("Json decode failed.")
		return
	}

	return


}
