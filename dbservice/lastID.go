package RCSdb

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"
)

// load and save id in txt
var defaultPath = "../laststore/last_id.txt"

func SetPath(path string){
	defaultPath = path
}

func LoadLastID() (int, error){
	data, err := os.ReadFile(defaultPath)
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
	res := os.WriteFile(defaultPath, data, 0644)
	return res
}

// load and save id in json
var jsonPath = "../laststore/last_id.json"

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


