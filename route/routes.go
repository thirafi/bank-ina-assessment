package route

import (
	"bank-ina-assessment/controller"
	"bank-ina-assessment/middleware"
	"bank-ina-assessment/repository/database"
	"bank-ina-assessment/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoute(e *gin.Engine, db *gorm.DB) {
	// e.Validator = &CustomValidator{validator: validator.New()}
	userRepository := database.NewUserRepository(db)
	taskRepository := database.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository)
	taskUsecase := usecase.NewTaskUsecase(taskRepository)

	authController := controller.NewAuthController(userUsecase)
	userController := controller.NewUserController(userUsecase, userRepository)
	taskController := controller.NewTaskController(taskUsecase, taskRepository)

	middleware := middleware.NewMiddleware(userRepository)

	e.POST("/login", authController.LoginUserController)
	e.GET("/callback", authController.LoginUserController)
	e.POST("/register", userController.CreateUserController)

	// user collection
	user := e.Group("/users")
	user.GET("", userController.GetAllUsersController)
	user.GET("/:id", userController.GetUserController)
	user.POST("", userController.CreateUserController)
	user.PUT("/:id", userController.UpdateUserController)
	user.DELETE("/:id", userController.DeleteUserController)

	// task collection
	task := e.Group("/tasks", middleware.ValidateToken())
	task.POST("", taskController.CreateTaskController)
	task.GET("", taskController.GetTasksController)
	task.GET("/:id", taskController.GetTaskController)
	task.PUT("/:id", taskController.UpdateTaskController)
	task.DELETE("/:id", taskController.DeleteTaskController)

}
