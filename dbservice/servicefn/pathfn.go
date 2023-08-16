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