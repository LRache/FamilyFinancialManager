package repository

func CreateUser(username string, password string) (int, error) {
	return 1, nil
}

func UserExists(username string) (bool, error) {
	if username == "exist" {
		return true, nil
	}

	return false, nil
}

func UserLogin(username string, password string) (int, error) {
	if username == "test" && password == "test" {
		return 1, nil
	}

	return 0, nil
}
