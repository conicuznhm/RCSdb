package servicefn

import (
	"encoding/json"
	"os"
	"time"
)

// struct blueprint
type LastData struct{
	Id int
	Timestamp time.Time
}

// save last id to path
func JSONSaveID(lastid_path string, lastData *LastData) error{
	lastData.Id++
	lastData.Timestamp = time.Now()

	data, err :=json.Marshal(lastData)
	if err != nil{
		return err
	}
	
	if err:=os.WriteFile(lastid_path, data, 0644); err != nil{
		return err
	}
	return nil
}

// load last id from path version 1
func JSONLoadID(lastid_path string) (*LastData, error){
	// var lastData *LastData  //bad declare because pointer not initialize can hold nil value
	lastData := &LastData{}
	data, err := os.ReadFile(lastid_path)
	if err != nil{
		return nil, err
	}
	if err = json.Unmarshal(data, lastData); err!=nil{
		return nil, err
	}
	return lastData, nil
}

// load last id from path version 2
func JSONLoadID2(lastid_path string) (*LastData, error){
	var lastData LastData
	data, err := os.ReadFile(lastid_path)
	if err != nil{
		return nil, err
	}
	if err = json.Unmarshal(data, &lastData); err!=nil{
		return nil, err
	}
	return &lastData, nil
}

