package ui

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sankethkini/StudentData/application"
	"github.com/sankethkini/StudentData/constants"
	"github.com/sankethkini/StudentData/domain/course"
	"github.com/sankethkini/StudentData/domain/user"
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

//userDataInput takes user realted info
func userDataInput() (fname string, address string, rollnum string, age int) {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the full name")
	scanner.Scan()
	fname = scanner.Text()
	fmt.Println("Enter the rollnum")
	fmt.Scanf("%s", &rollnum)
	fmt.Println("Enter the Address")
	scanner.Scan()
	address = scanner.Text()
	fmt.Println("Enter the age")
	fmt.Scanf("%d", &age)
	return
}

//cli function for user to add new user
func addUser() {

	//user data
	fname, address, rollnum, age := userDataInput()

	//taking course input from user
	selectedCourse := takeCourseInput()

	user := user.NewUser(fname, age, address, rollnum, selectedCourse)

	//adding user
	res, err := application.Add(user)
	if err != nil {
		fmt.Println(err)

	} else {
		//display msg from application
		for _, val := range res {
			fmt.Println(val["message"])
		}
	}
}

//takeDisplayInput function takes input from user required to display all users according to preference
func takeDisplayInput() map[string]interface{} {
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

	}

	return data
}

//cli func for display users
func displayUser() {

	data := takeDisplayInput()

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
	msg, err := application.Save()

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
	fmt.Println("do you want to save the users [y/n]")
	var ch string
	fmt.Scanf("%s", &ch)
	if ch == "y" {
		saveUser()
	} else {
		fmt.Println("exiting program without saving user")
	}
	fmt.Println("exiting .... ")
	application.Exit()
}

//taking menu option from user
func Start() {
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
