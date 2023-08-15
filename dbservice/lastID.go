package rcsdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

func JSONcreateLastID(path string) error{
	if path == ""{
		path = DefaultPath
	}
	path +="/lastid"

	if err := os.MkdirAll(path,os.ModeDir|0755); err != nil{
		fmt.Println(err)
		return errors.New("Cannot create lastid directory!!!")
	}

	type Data struct{
		Id int `json:"id"`;
		Timestamp time.Time `json:"timestamp"`;
	}

	data := Data{Id:0,Timestamp:time.Now()}
	jsonByte, err := json.Marshal(data)
	if err != nil{
		fmt.Println(err)
		return errors.New("Cannot convert struct instance to JSON-encoded slice byte")
	}

	filePath := path+"/last_id.json"
	if err := os.WriteFile(filePath,jsonByte,0644); err != nil{
		fmt.Println(err)
		return errors.New("Cannot create last_id.json file!!!")
	}

	return nil
}