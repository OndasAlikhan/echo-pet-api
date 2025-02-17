package routes

import (
	"auth-service/src/application"
	"auth-service/src/controller"
	"auth-service/src/middleware"
	"os"
)

func RegisterRoute(urlPrefix string) {
	routeGroup := application.GlobalApp.Group(urlPrefix)
	authM := middleware.NewJwtMiddleware([]byte(os.Getenv("JWT_SECRET")), "Bearer ")
	permissionM := middleware.NewPermissionMiddleware()

	authController := controller.NewAuthController()
	routeGroup.POST("/login", authController.Login)
	routeGroup.GET("/me", authController.Me, authM.Run())

	userController := controller.NewUserController()
	routeGroup.GET("/users", userController.Index, authM.Run(), permissionM.Run("users-index"))
	routeGroup.POST("/users", userController.Store)
	routeGroup.GET("/users/:id", userController.Show, authM.Run())
	routeGroup.PATCH("/users/:id", userController.UpdatePassword, authM.Run())
	routeGroup.DELETE("/users/:id", userController.Delete, authM.Run())

	roleController := controller.NewRoleController()
	routeGroup.GET("/roles", roleController.Index, authM.Run(), permissionM.Run("roles-index"))
	routeGroup.POST("/roles", roleController.Store, authM.Run(), permissionM.Run("roles-store"))
	routeGroup.GET("/roles/:id", roleController.Show, authM.Run(), permissionM.Run("roles-show"))
	routeGroup.PATCH("/roles/:id", roleController.Update, authM.Run(), permissionM.Run("roles-update"))
	routeGroup.DELETE("/roles/:id", roleController.Delete, authM.Run(), permissionM.Run("roles-delete"))
	routeGroup.POST("/roles/:id/permissions", roleController.AddPermissions, authM.Run(), permissionM.Run("roles-add-permissions"))
	routeGroup.DELETE("/roles/:id/permissions", roleController.RemovePermissions, authM.Run(), permissionM.Run("roles-remove-permissions"))

	permissionController := controller.NewPermissionController()
	routeGroup.GET("/permissions", permissionController.Index, authM.Run(), permissionM.Run("permissions-index"))
	routeGroup.POST("/permissions", permissionController.Store, authM.Run(), permissionM.Run("permissions-store"))
	routeGroup.GET("/permissions/:id", permissionController.Show, authM.Run(), permissionM.Run("permissions-show"))
	routeGroup.PATCH("/permissions/:id", permissionController.Update, authM.Run(), permissionM.Run("permissions-update"))
	routeGroup.DELETE("/permissions/:id", permissionController.Delete, authM.Run(), permissionM.Run("permissions-delete"))
}
