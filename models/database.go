package models

func SaveUser(username string, name string) error {
	db := GetDB()
	_, err := db.Exec("INSERT INTO users (username, name) VALUES($1, $2)", username, name)
	return err
}

func GetUsers() (*[]User, error) {
	var result *[]User
	db := GetDB()
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&result)
	}
	return result, err
}
