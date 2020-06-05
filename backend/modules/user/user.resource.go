package usermodule

import (
	"gousers/core"
)

// BootstrapRoutes ... to bootstrap routes of User module
func BootstrapRoutes(ir *core.GousersRouter) {
	ir.POST("/createUser", CreateUserRoute)
	ir.GET("/getAllUsers", GetAllUserRoute)
	ir.PUT("/updateUser/:id", UpdateUserRoute)
	ir.GET("/getUser/:id", GetUserRoute)
	ir.DELETE("/deleteUser/:id", DeleteUserRoute)
}
