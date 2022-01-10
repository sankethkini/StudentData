package application

//alisaing map
type data map[string]interface{}

//validators of input
func InputValidator(userdata data) error {

	curname := userdata["fname"].(string)
	if len(curname) < 2 {
		return NoNameErr
	}
	curroll := userdata["rollnum"].(string)
	if len(curroll) < 2 {
		return NoRollNum
	}
	curadd := userdata["address"].(string)
	if len(curadd) < 2 {
		return NoAddress
	}
	curage := userdata["age"].(int)
	if curage <= 0 || curage >= 120 {
		return AgeErr
	}

	return nil
}

//function to add user
func Add(data) ([]data, error) {
	return nil, nil
}

//function to display
func Display(data) ([]data, error) {
	return nil, nil
}

func Delete(data) ([]data, error) {
	return nil, nil
}

func Save(data) ([]data, error) {
	return nil, nil
}

func Exit(data) ([]data, error) {
	return nil, nil
}
