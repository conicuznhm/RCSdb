package servicefn

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// "fmt"

type Path struct{
	Data string `json:"data"`
	Lastid string `json:"lastid"`
}

// get path
func GetPath(name string) (*Path, error){

	path := &Path{}
	pathDir := "database/path/"+name+"path.txt"
	readFile, err := os.Open(pathDir)
	if err != nil{
		if os.IsNotExist(err){
			return path, errors.New("The file to open dose not found")
		}else{
			return path, errors.New("Cannot open the file")
		}
	}
	defer readFile.Close()

	fileScan := bufio.NewScanner(readFile)
	fileScan.Split(bufio.ScanLines)
	
	keys := []string{}
	values := make([]string,0)
	
	for fileScan.Scan(){
		item := strings.Split(fileScan.Text()," ")
		keys = append(keys,item[0])
		values = append(values,item[1])
		
		if len(item) >=2{
			k := item[0]
			v := item[1]
			switch k{
			case "data":
				path.Data = v
			case "lastid":
				path.Lastid = v
			default:
				return path, errors.New("Cannot create path struct")
			}
		}
	}
	
	fmt.Println(keys)
	fmt.Println(values)
	return path, nil
}

// create path
func CreatePath(path string,dirDB string,name string) error{
	
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

