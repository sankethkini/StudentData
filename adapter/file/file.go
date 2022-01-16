package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sankethkini/StudentData/constants"
	"github.com/sankethkini/StudentData/domain/user"
)

//struct for file adapter
type adapter struct {
}

//checking if file exists
func checkFileExists(filename string) bool {
	_, err := os.Stat(constants.Filelocation)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//removeAndCreateFile removes exisiting file and creates new file
func removeAndCreateFile() (*os.File, error) {
	//removing file
	exi := checkFileExists(constants.Filelocation)
	if exi {
		err := os.Remove(constants.Filelocation)
		if err != nil {
			return nil, fmt.Errorf("cannot remove file %w", err)
		}
	}

	//creating file.json
	fileinfo, err := os.Create(constants.Filelocation)
	if err != nil {
		return nil, fmt.Errorf("cannot create file %w", err)
	}
	return fileinfo, nil
}

//saving into disk
//convering into json and saving into file.json
func (file *adapter) Save(items interface{}) error {
	//marshleing record into json
	curItems := items.([]user.User)
	jsondata, err := json.Marshal(curItems)
	if err != nil {
		return fmt.Errorf("error in json marshal %w", err)
	}

	//take file pointer
	fileinfo, err := removeAndCreateFile()
	if err != nil {
		return err
	}

	defer fileinfo.Close()

	//write into file
	_, err = fileinfo.Write(jsondata)
	if err != nil {
		return fmt.Errorf("cannot write on file %w", err)
	}
	return nil
}

//retriving from file.json
//unmarsheling into user struct
func (file *adapter) RetriveAll() (interface{}, error) {
	//check for file existence
	if !checkFileExists(constants.Filelocation) {
		//else create file
		fileinfo, err := removeAndCreateFile()
		if err != nil {
			return nil, err
		}
		fileinfo.Close()
	}

	fileinfo, err := os.Open(constants.Filelocation)
	if err != nil {
		return nil, fmt.Errorf("cannot open the file %w", err)
	}
	defer fileinfo.Close()

	data, err := ioutil.ReadAll(fileinfo)
	if len(data) == 0 {
		var dataret []user.User
		return dataret, nil
	}
	if err != nil {
		return nil, fmt.Errorf("cannot read the file %w", err)
	}

	var dataret []user.User
	err = json.Unmarshal(data, &dataret)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarsh the contents of the file %w", err)
	}

	return dataret, nil
}

//constructor
func NewFileAdapter() (*adapter, error) {
	return &adapter{}, nil
}
