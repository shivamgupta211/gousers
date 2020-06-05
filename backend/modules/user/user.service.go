package usermodule

import (
	"errors"
	"gousers/db"
	"gousers/logger"
	"gousers/utils"
)

// CreateUser ... to create User
func CreateUser(user User) (User, error) {
	user.CreatedAt = utils.GetNowUTC()
	user.UpdatedAt = utils.GetNowUTC()

	connection := db.GetPgConnection()
	defer connection.Close()

	err := connection.Insert(&user)
	if err != nil {
		logger.CaptureError(err, 400, "Error while Inserting User")
		return user, errors.New("Error while Inserting User")
	}
	connection.Model(&user).Returning("id").Update()

	return user, nil
}

// GetUser ... get the User
func GetUser(UserID int) (User, error) {
	user := User{
		ID: UserID,
	}

	connection := db.GetPgConnection()
	defer connection.Close()

	err := connection.Select(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// GetAllUsers ... get the User
func GetAllUsers() ([]User, error) {
	connection := db.GetPgConnection()
	defer connection.Close()
	var users []User
	erro := connection.Model(&users).Select()

	if erro != nil {
		return users, erro
	}

	return users, nil
}

// UpdateUser ... UpdateUser user
func UpdateUser(id int, user User) (User, error) {
	user.UpdatedAt = utils.GetNowUTC()

	connection := db.GetPgConnection()
	defer connection.Close()

	err := connection.Update(&user)

	if err != nil {
		logger.CaptureError(err, 500, "Error while updating data")
		return user, errors.New("Error while updating data")
	}

	return user, nil
}

// DeleteUser ... DeleteUser user
func DeleteUser(id int) (User, error) {
	user := User{
		ID: id,
	}
	connection := db.GetPgConnection()
	defer connection.Close()

	err := connection.Delete(&user)

	if err != nil {
		logger.CaptureError(err, 500, "Error while deleting data")
		return user, errors.New("Error while deleting data")
	}

	return user, nil
}
