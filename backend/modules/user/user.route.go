package usermodule

import (
	"errors"
	"gousers/core"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

// CreateUserRoute ... to create new User
func CreateUserRoute(c *gin.Context) (core.ReturnType, error) {
	var data User
	c.BindJSON(&data)

	result, err := CreateUserController(data)

	if err != nil {
		response := core.ReturnType{
			Status:   400,
			Response: err.Error(),
		}
		return response, nil
	}
	response := core.ReturnType{
		Status:   200,
		Response: result,
	}
	return response, nil
}

// GetUserRoute ... to get User based on id
func GetUserRoute(c *gin.Context) (core.ReturnType, error) {
	id := c.Param("id")

	if id == "" {
		return core.ReturnType{
			Status:   400,
			Response: "Please provide the UserId",
		}, nil
	}

	UserID, err := strconv.Atoi(id)
	if err != nil {
		return core.ReturnType{
			Status:   400,
			Response: "Please provide valid UserId",
		}, nil
	}

	result, err := GetUser(UserID)
	if err != nil {
		if err == pg.ErrNoRows {
			response := core.ReturnType{
				Status:   400,
				Response: "No User found for this ID",
			}
			return response, nil
		}
		response := core.ReturnType{
			Status:   400,
			Response: err.Error(),
		}
		return response, nil
	}
	response := core.ReturnType{
		Status:   200,
		Response: result,
	}
	return response, nil
}

// GetAllUserRoute ... to get all Users
func GetAllUserRoute(c *gin.Context) (core.ReturnType, error) {
	result, err := GetAllUsers()
	if err != nil {
		if err == pg.ErrNoRows {
			response := core.ReturnType{
				Status:   400,
				Response: "No User found for this ID",
			}
			return response, nil
		}
		response := core.ReturnType{
			Status:   400,
			Response: err.Error(),
		}
		return response, nil
	}
	response := core.ReturnType{
		Status:   200,
		Response: result,
	}
	return response, nil
}

// UpdateUserRoute ... to check user's acceptance eligibility
func UpdateUserRoute(c *gin.Context) (core.ReturnType, error) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return core.ReturnType{
			Status:   400,
			Response: "Wrong Input",
		}, nil
	}

	var data User
	c.BindJSON(&data)
	data.ID = userID

	result, err := UpdateUser(data.ID, data)
	if err != nil {
		if err == pg.ErrNoRows {
			response := core.ReturnType{
				Status:   400,
				Response: errors.New("No User found for this ID"),
			}
			return response, nil
		}
		response := core.ReturnType{
			Status:   400,
			Response: err.Error(),
		}
		return response, nil
	}
	response := core.ReturnType{
		Status:   200,
		Response: result,
	}
	return response, nil
}

// DeleteUserRoute ... to check user's acceptance eligibility
func DeleteUserRoute(c *gin.Context) (core.ReturnType, error) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		return core.ReturnType{
			Status:   400,
			Response: "Wrong Input",
		}, nil
	}

	result, err := DeleteUser(userID)

	if err != nil {
		if err == pg.ErrNoRows {
			response := core.ReturnType{
				Status:   400,
				Response: errors.New("No User found for this ID"),
			}
			return response, nil
		}
		response := core.ReturnType{
			Status:   400,
			Response: err.Error(),
		}
		return response, nil
	}
	response := core.ReturnType{
		Status:   200,
		Response: result,
	}
	return response, nil
}
