package rcsdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

// create last_id.json
func JSONcreateLastID(path string, name string) error{
	if path == ""{
		path = defaultPath
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

	filePath := path+"/"+name+"_lastid.json"

	// file existing guard clause
	if _,err := os.Stat(filePath); !os.IsNotExist(err){
		return errors.New(name+"_lastid.json"+" already exists")
	}
	// create file ____lastid.json
	if err := os.WriteFile(filePath,jsonByte,0644); err != nil{
		fmt.Println(err)
		return errors.New("Cannot create file "+name+"_lastid.json")
	}

	return nil
}

func CreatePath(path string,name string) error{
	
	filePath := dirDB+"/path/"+name+"path.txt"

	if _, err := os.Stat(filePath); os.IsExist(err){
		return errors.New(name+"path.txt has already existed")
	}

	if err := os.MkdirAll(dirDB+"/path",0755); err != nil{
		fmt.Println(err)
		return errors.New("Cannot create path directory!!!")
	}
	
	datapath := "data"+" "+path+"/"+name+".json"
	lastid := "lastid"+" "+path+"/lastid/"+name+"_lastid.json"

	// file existing guard clause
	if _,err := os.Stat(filePath); !os.IsNotExist(err){
		return errors.New(name+"path.txt has already existed")
	}

	data := []string{datapath,lastid}

	file, err := os.Create(filePath)
	if err != nil {
		return errors.New("Cannot create file")
	}
	defer file.Close()

	for _,line := range data{
		_,err := file.WriteString(line+ "\n")
		if err !=nil{
			return err
		}
	}
	return nil
}
