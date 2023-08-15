package rcsdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// load and save id in txt
// var defaultPath = "dataApi/utilstore/last_id.txt"
var DefaultPath = "dataApi"

func SetPath(path string){
	DefaultPath = path
}

// create database json
func JSONdbCreate(path string, name string) error{
	if path == ""{
		path = DefaultPath
	}

	filePath := path+"/"+name+".json"
	_, err := os.ReadFile(filePath)
	
	if err == nil{
		return errors.New("The file name has already existed!")
	}
	
	// emptyMap := map[string]interface{}{}
	// data := make(map[string]interface{})

	data := []interface{}{}  // empty array or slice
	emptySlice, err := json.Marshal(data)
	if err != nil{
		fmt.Println(err)
		return errors.New("Cannot create JSON-encoded!!!")
	}

	// create nested dir with mode permission d755
	if err := os.MkdirAll(path,os.ModeDir|0755); err != nil{
		fmt.Println(err)
		return errors.New("Cannot create directory!!!")
	}

	if err := os.WriteFile(filePath,emptySlice,0744); err != nil{
		fmt.Println(err)
		return errors.New("Cannot create JSON file!!!")
	}

	if err := JSONcreateLastID(path); err != nil{
		return err
	}

	return nil
}

// create last_id.json
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

func LoadLastID() (int, error){
	data, err := os.ReadFile(DefaultPath)
	if err != nil{
		return 0, err
	}

	id, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil{
		return 0, err
	}

	return id, nil
}

func SaveLastID(id int) error{
	id++
	data := []byte(strconv.Itoa(id))
	res := os.WriteFile(DefaultPath, data, 0644)
	return res
}

// load and save id in json
var jsonPath = "dataApi/utilstore/last_id.json"

type LastData struct{
	Id int
	Timestamp time.Time
}

func JSONLoadID() (*LastData, error){
	var lastData LastData
	data, err := os.ReadFile(jsonPath)
	if err != nil{
		return nil, err
	}
	if err = json.Unmarshal(data, &lastData); err!=nil{
		return nil, err
	}
	return &lastData, nil
}

func JSONSaveID(lastData *LastData) error{
	lastData.Id++
	lastData.Timestamp = time.Now()

	data, err :=json.Marshal(lastData)
	if err != nil{
		return err
	}
	
	if err:=os.WriteFile(jsonPath, data, 0644); err != nil{
		return err
	}
	return nil
}

// func main(){
// 	err := JSONdbCreate("dataApi/source","user")
// 	fmt.Println(err)
// }

	// create nested dir with mode permission d777
	// if err := os.MkdirAll("/dataApi/source",os.ModePerm); err != nil{
	// 	fmt.Println(err)
	// 	return errors.New("Cannot create directory!!!")
	// }

	// create single dir with mode permission d755
	// if err := os.Mkdir("/dataApi",os.ModeDir|0755); err != nil{
	// 	fmt.Println(err)
	// 	return errors.New("Cannot create directory!!!")
	// }