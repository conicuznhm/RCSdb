package rcsdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/conicuznhm/rcsdb/dbservice/servicefn"
	// "time"
)

// load and save id in txt
// var defaultDir = "dataApi/utilstore/last_id.txt"
var defaultDir = ""

// set database dir
var dirDB = "database"

func SetPath(path string){
	defaultDir = path
}

// create database json
func JSONdbCreate(path string, name string) error{

	if name == ""{
		return errors.New("The name of the database model is required")
	}
	
	if path == ""{
		if defaultDir == ""{
			path = dirDB
		}else{
			path = dirDB+"/"+defaultDir
		}
	}else{
		path = dirDB+"/"+path
	}

	filePath := path+"/"+name+".json"

	// check whether a file has already existed or not
	if _,err := os.Stat(filePath); !os.IsNotExist(err){
		return errors.New(name+".json has already existed")
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
	if err := os.MkdirAll(path,0755); err != nil{
		fmt.Println(err)
		return errors.New("Cannot create directory!!!")
	}

	if err := os.WriteFile(filePath,emptySlice,0744); err != nil{
		fmt.Println(err)
		return errors.New("Cannot create JSON file!!!")
	}

	// create last_id.json
	if err := JSONcreateLastID(path, name); err != nil{
		return err
	}

	// create path dir and file by using CreatePath from servicefn module
	if err := servicefn.CreatePath(path,dirDB,name); err != nil{
		return err
	}

	return nil
}

func LoadLastID() (int, error){
	data, err := os.ReadFile(defaultDir)
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
	res := os.WriteFile(defaultDir, data, 0644)
	return res
}

// load and save id in json
// var jsonPath = "dataApi/utilstore/last_id.json"

// type LastData struct{
// 	Id int
// 	Timestamp time.Time
// }

// func JSONLoadID() (*LastData, error){
// 	var lastData LastData
// 	data, err := os.ReadFile(jsonPath)
// 	if err != nil{
// 		return nil, err
// 	}
// 	if err = json.Unmarshal(data, &lastData); err!=nil{
// 		return nil, err
// 	}
// 	return &lastData, nil
// }

// func JSONSaveID(lastData *LastData) error{
// 	lastData.Id++
// 	lastData.Timestamp = time.Now()

// 	data, err :=json.Marshal(lastData)
// 	if err != nil{
// 		return err
// 	}
	
// 	if err:=os.WriteFile(jsonPath, data, 0644); err != nil{
// 		return err
// 	}
// 	return nil
// }

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