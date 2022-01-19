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
	_, err := os.Stat(filename)
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
	file, err := os.Create(constants.Filelocation)
	if err != nil {
		return nil, fmt.Errorf("cannot create file %w", err)
	}

	return file, nil

}

//saving into disk
//convering into json and saving into file.json
func (file *adapter) Save(items []user.User) error {
	//marshleing record into json

	jsonData, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("error in json marshal %w", err)
	}

	//take file pointer
	fileInfo, err := removeAndCreateFile()
	if err != nil {
		return err
	}
	defer fileInfo.Close()
	//write into file
	_, err = fileInfo.Write(jsonData)
	if err != nil {
		return fmt.Errorf("cannot write on file %w", err)
	}
	return nil
}

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(constants.Filelocation)
	if err != nil {
		return nil, fmt.Errorf("cannot open the file %w", err)
	}
	return file, nil

}

//retriving from file.json
//unmarsheling into user struct
func (file *adapter) RetriveAll() ([]user.User, error) {
	//check for file existence
	if !checkFileExists(constants.Filelocation) {
		//else create file
		_, err := removeAndCreateFile()
		if err != nil {
			return nil, err
		}

	}

	fi, err := openFile(constants.Filelocation)
	if err != nil {
		return nil, err
	}
	defer fi.Close()

	data, err := ioutil.ReadAll(fi)
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
