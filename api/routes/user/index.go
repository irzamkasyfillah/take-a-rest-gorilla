package user

import (
	"github.com/gorilla/mux"
	"github.com/irzam/my-app/api/user/controller"
	"github.com/irzam/my-app/api/user/entity/repository"
	"github.com/irzam/my-app/api/user/service"
	"github.com/irzam/my-app/api/user/service/action"
	"github.com/irzam/my-app/lib/database"
)

func GenerateUserRoutes(r *mux.Router) {
	// User routes
	userRepository := repository.NewUserRepository()
	userHistoryRepository := repository.NewUserHistoryRepository()
	userAction := action.NewUserAction(userRepository, userHistoryRepository, database.GetDB())
	userService := service.NewUserService(userAction)
	userController := controller.NewUserController(userService)
	userRoutes := NewUserRoutes(userController)

	userRouter := r.PathPrefix("/users").Subrouter()
	userRouter.HandleFunc("/{id}", userRoutes.GetUser).Methods("GET")
	userRouter.HandleFunc("", userRoutes.GetAllUsers).Methods("GET")
	userRouter.HandleFunc("", userRoutes.CreateUser).Methods("POST")
	userRouter.HandleFunc("/{id}", userRoutes.UpdateUser).Methods("PUT")
	userRouter.HandleFunc("/{id}", userRoutes.DeleteUser).Methods("DELETE")
	userRouter.HandleFunc("/history", userRoutes.GetUserHistory).Methods("POST")
}
