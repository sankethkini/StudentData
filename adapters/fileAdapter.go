package adapters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sankethkini/StudentData/domain/user"
)

type FileAdapter struct {
}

// func checkFileExists(filename string) bool{
// 	fileInfo, err := os.Stat("data.json")
// 	if err != nil {
// 		if !os.IsNotExist(err) {
// 			fmt.Println(fileInfo)
// 			return false
// 		}
// 	}
// }
func (file *FileAdapter) Save(items interface{}) error {
	curItems := items.([]user.User)
	jsondata, err := json.Marshal(curItems)
	if err != nil {
		return err
	}

	os.Remove("data.json")

	fileinfo, err := os.Create("data.json")
	if err != nil {
		return err
	}
	defer fileinfo.Close()
	n, err := fileinfo.Write(jsondata)
	if err != nil {
		fmt.Println(n, err)
		return err
	}
	return nil
}

func (file *FileAdapter) RetriveAll(field string) ([]user.User, error) {
	fileinfo, err := os.Open("data.json")
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(fileinfo)
	if err != nil {
		return nil, err
	}
	var dataret []user.User
	err = json.Unmarshal(data, &dataret)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return dataret, nil
}
