package application

import (
	"fmt"
	"os"

	"github.com/sankethkini/StudentData/adapter"
	"github.com/sankethkini/StudentData/adapter/file"
	"github.com/sankethkini/StudentData/adapter/memory"
	"github.com/sankethkini/StudentData/domain/user"
)

// alisaing map.
type Data map[string]interface{}

type IApp interface {
	Add(user.User) ([]Data, error)
	Display(Data) ([]Data, error)
	Delete(Data) ([]Data, error)
	Save() ([]Data, error)
	Exit()
}

type App struct {
	adpt *adapter.Adapter
}

func NewApp() (*App, error) {
	// intializing file Adapter.
	fileAdapter, err := file.NewFileAdapter()
	if err != nil {
		return nil, err
	}

	// intializing memory Adapter.
	memoryAdapter, err := memory.NewMemory()
	if err != nil {
		return nil, err
	}

	// intializing Adapter.
	adpt, err := adapter.NewAdapter(fileAdapter, memoryAdapter)
	if err != nil {
		return nil, err
	}

	app := App{adpt: adpt}
	app.fileToMemory()
	return &app, nil
}

func (app *App) fileToMemory() {
	users, err := app.adpt.FileAdapter.RetriveAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = app.adpt.MemoryAdapter.Save(users...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// function to add user.
func (app *App) Add(user user.User) ([]Data, error) {
	// checking validity of user input.
	validationErr := user.Validate()
	if validationErr != nil {
		return nil, validationErr
	}

	// checking for rollnum existence.
	isRollexists := app.checkForRoll(user.RollNo)
	if isRollexists != nil {
		return nil, ErrRollExists
	}

	// adding user.
	err := app.adpt.MemoryAdapter.Save(user)
	if err != nil {
		return nil, err
	}

	// message.
	msg := createMessage("user added successfully")
	return msg, err
}

// validations for rollnumber exists.
func (app *App) checkForRoll(rollnum string) error {
	// checking existence of simialr rollnum.
	isExists := app.adpt.MemoryAdapter.Retrieve("rollnum", rollnum)
	if isExists {
		return ErrRollExists
	}
	return nil
}

// function to display.
func (app *App) Display(userData Data) ([]user.User, error) {
	field := userData["field"].(string)
	order := userData["order"].(int)

	// retrieving all data.
	items, err := app.adpt.MemoryAdapter.RetriveAll(field, order)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (app *App) Delete(userData Data) ([]Data, error) {
	roll := userData["rollnum"].(string)

	// deleting user.
	err := app.adpt.MemoryAdapter.Delete("rollnum", roll)
	if err != nil {
		return nil, err
	}

	msg := createMessage("user deleted successfully")
	return msg, err
}

func (app *App) Save() ([]Data, error) {
	// retirving all by name.
	allData, err := app.adpt.MemoryAdapter.RetriveAll("name", 1)
	if err != nil {
		return nil, err
	}

	// saving on disk.
	err = app.adpt.FileAdapter.Save(allData)
	if err != nil {
		return nil, err
	}
	msg := createMessage("users saved to disk successfully")
	return msg, err
}

// exit function.
func (app *App) Exit() {
	os.Exit(1)
}

// function to return message.
func createMessage(m string) []Data {
	var msg []Data
	mp := make(map[string]interface{})
	mp["message"] = m
	msg = append(msg, mp)
	return msg
}
