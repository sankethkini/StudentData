package file

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/sankethkini/StudentData/constants"
	"github.com/sankethkini/StudentData/domain/user"
)

// struct for file Adapter.
type Adapter struct{}

// constructor.
func NewFileAdapter() (*Adapter, error) {
	return &Adapter{}, nil
}

// checking if file exists.
func checkFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// removeAndCreateFile removes exisiting file and creates new file.
func removeAndCreateFile() (*os.File, error) {
	// removing file.
	exi := checkFileExists(constants.Filelocation)
	if exi {
		err := os.Remove(constants.Filelocation)
		if err != nil {
			return nil, errors.Wrap(err, "cannot remove file")
		}
	}

	// creating file.json
	file, err := os.Create(constants.Filelocation)
	if err != nil {
		return nil, errors.Wrap(err, "cannot create the file")
	}

	return file, nil
}

// saving into disk.
// convering into json and saving into file.json.
func (file *Adapter) Save(items []user.User) error {
	// marshleing record into json.

	jsonData, err := json.Marshal(items)
	if err != nil {
		return errors.Wrap(err, "error in json marshal")
	}

	// take file pointer.
	fileInfo, err := removeAndCreateFile()
	if err != nil {
		return err
	}
	defer fileInfo.Close()
	// write into file.
	_, err = fileInfo.Write(jsonData)
	if err != nil {
		return errors.Wrap(err, "cannot write on files")
	}
	return nil
}

func openFile(filename string) (*os.File, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "cannot open the file")
	}
	return file, nil
}

// retrieving from file.json.
// unmarsheling into user struct.
func (file *Adapter) RetriveAll() ([]user.User, error) {
	// check for file existence
	if !checkFileExists(constants.Filelocation) {
		// else create file.
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
		return nil, errors.Wrap(err, "cannot read file")
	}

	var dataret []user.User
	err = json.Unmarshal(data, &dataret)
	if err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal file")
	}

	return dataret, nil
}
