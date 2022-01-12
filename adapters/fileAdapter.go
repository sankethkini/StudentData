package adapters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sankethkini/StudentData/constants"
	"github.com/sankethkini/StudentData/domain/user"
)

//struct for file adapter
type FileAdapter struct {
}

//checking if file exists
func checkFileExists(filename string) bool {
	fileInfo, err := os.Stat(constants.Filelocation)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(fileInfo)
			return false
		}
	}
	return true

}

//saving into disk
//convering into json and saving into file.json
func (file *FileAdapter) Save(items interface{}) error {
	//marshleing record into json
	curItems := items.([]user.User)
	jsondata, err := json.Marshal(curItems)
	if err != nil {
		return err
	}

	//removing file
	os.Remove(constants.Filelocation)

	//creating file.json
	fileinfo, err := os.Create(constants.Filelocation)
	if err != nil {
		return err
	}
	defer fileinfo.Close()

	//write into file
	n, err := fileinfo.Write(jsondata)
	if err != nil {
		fmt.Println(n, err)
		return err
	}
	return nil
}

//retriving from file.json
//unmarsheling into user struct
func (file *FileAdapter) RetriveAll(field string, order int) (interface{}, error) {
	//check for file existence
	if !checkFileExists(constants.Filelocation) {
		//else create file
		fileinfo, err := os.Create(constants.Filelocation)
		if err != nil {
			return nil, err
		}
		fileinfo.Close()
	}

	fileinfo, err := os.Open(constants.Filelocation)

	if err != nil {
		return nil, err
	}
	defer fileinfo.Close()

	data, err := ioutil.ReadAll(fileinfo)
	if len(data) == 0 {
		var dataret []user.User
		return dataret, nil
	}
	if err != nil {
		return nil, err
	}
	var dataret []user.User
	err = json.Unmarshal(data, &dataret)
	if err != nil {

		return nil, err
	}

	return dataret, nil
}

func (adapter *FileAdapter) Retrive(feild string, value string) bool {
	return false
}

func (adapter *FileAdapter) Delete(field string, value string) error {
	return nil
}
