package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/renanfvcunha/huncoding-go-first-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getAllUsers", userController.FindAllUsers)
	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	r.POST("/login", userController.LoginUser)
}
