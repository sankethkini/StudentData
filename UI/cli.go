package UI

import (
	"fmt"
	"os"

	"github.com/sankethkini/StudentData/application"
	"github.com/sankethkini/StudentData/constants"
	"github.com/sankethkini/StudentData/domain/course"
)

//function to print menu for user
func printMenu() {
	fmt.Println("Choose options:")
	fmt.Println("1. Add user details")
	fmt.Println("2. Display user details")
	fmt.Println("3. Delete user details")
	fmt.Println("4. Save user details")
	fmt.Println("5. Exit")
}

//cli func for taking course input from user
func takeCourseInput() []course.Course {

	var selectedCourse []course.Course
	fmt.Println("enter", constants.NoOfCourses, "course")
	for i, val := range constants.AllCourses {
		fmt.Printf(" | %d.  course code: %s    coursename: %s | ", i, val.Code, val.Name)
	}

	//taking input till all courses are entered properly
	for i := 1; i <= constants.NoOfCourses; {
		fmt.Printf("\nchoose course number %d", i)
		var code int
		fmt.Scanf("%d", &code)

		//check for proper code
		if code < 0 || code >= len(constants.AllCourses) {
			fmt.Println("Enter the proper code")
			continue
		} else {

			//checking if course is already seleceted
			var selected bool = false
			for j := 0; j < len(selectedCourse); j++ {
				if constants.AllCourses[code].Code == selectedCourse[j].Code {
					fmt.Println("course already selected")
					selected = true
				}
			}
			if selected {
				continue
			}

			//appending course to the list
			selectedCourse = append(selectedCourse, constants.AllCourses[code])
		}
		i++
	}
	return selectedCourse
}

//cli function for user to add new user
func addUser() {
	var fname, address, rollnum string
	var age int

	fmt.Println("Enter the full name")
	fmt.Scanf("%s", &fname)
	fmt.Println("Enter the rollnum")
	fmt.Scanf("%s", &rollnum)
	fmt.Println("Enter the Address")
	fmt.Scanf("%s", &address)
	fmt.Println("Enter the age")
	fmt.Scanf("%d", &age)

	//taking course input from user
	selectedCourse := takeCourseInput()

	userdata := make(map[string]interface{})
	userdata["fname"] = fname
	userdata["rollnum"] = rollnum
	userdata["address"] = address
	userdata["age"] = age
	userdata["courses"] = selectedCourse

	//adding user
	res, err := application.Add(userdata)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		//display msg from application
		for _, val := range res {
			fmt.Println(val["message"])
		}
	}
}

//cli func for display users
func displayUser() {

	//options to user
	fmt.Println("Enter the field you want to sort by")
	fmt.Println("1. Name")
	fmt.Println("2. RollNumber")
	fmt.Println("3. Age")
	fmt.Println("4. Address")

	var option1, option2 int
	fmt.Scanf("%d", &option1)

	fmt.Println("Enter the order 1. Ascending 2. Descending")
	fmt.Scanf("%d", &option2)

	data := make(map[string]interface{})
	data["order"] = option2

	switch option1 {
	case 1:
		data["field"] = "name"
	case 2:
		data["field"] = "rollnum"
	case 3:
		data["field"] = "age"
	case 4:
		data["field"] = "address"
	default:
		data["field"] = "name"
		fmt.Println("Enter the valid number")
	}

	//displaying data from application
	users, err := application.Display(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} else {
		for _, val := range users {
			fmt.Printf("Name: %v Rollnum: %v Address: %v Age:%v\n", val.Fname, val.RollNo, val.Adress, val.Age)
			fmt.Printf(" | Course 1: %s Course 2: %s Course 3: %s Course 3:%s | \n", val.Courses[0].Code, val.Courses[1].Code, val.Courses[2].Code, val.Courses[3].Code)
		}
	}
}

//cli func for delete users
func deleteUser() {
	//taking input from user
	fmt.Println("Enter user rollnumber ")
	var roll string
	fmt.Scanf("%s", &roll)

	data := make(map[string]interface{})
	data["rollnum"] = roll

	//calling delete from app
	msg, err := application.Delete(data)

	//displaying message
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range msg {
			fmt.Println(val["message"])
		}
	}
}

//cli for save
func saveUser() {
	//calling save functiom
	msg, err := application.Save(nil)

	//displaying message
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range msg {
			fmt.Println(val["message"])
		}
	}
}

//cli for exit
func exit() {
	fmt.Println("exiting ..... ")
	msg, err := application.Exit(nil)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, val := range msg {
			fmt.Println(val["message"])
		}
	}
}

//taking menu option from user
func Start() error {
	for {
		printMenu()
		var option int
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			addUser()
		case 2:
			displayUser()
		case 3:
			deleteUser()
		case 4:
			saveUser()
		case 5:
			exit()
		}

	}

}
