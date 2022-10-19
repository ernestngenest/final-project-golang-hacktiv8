package server

import (
	"final_project_hacktiv8/middleware"
	user_ctrl "final_project_hacktiv8/modules/users/controller"
	user_repo "final_project_hacktiv8/modules/users/repository"
	user_serv "final_project_hacktiv8/modules/users/service"

	photo_ctrl "final_project_hacktiv8/modules/photos/controller"
	photo_repo "final_project_hacktiv8/modules/photos/repository"
	photo_serv "final_project_hacktiv8/modules/photos/service"

	commect_ctrl "final_project_hacktiv8/modules/comments/controller"
	commect_repo "final_project_hacktiv8/modules/comments/repository"
	commect_serv "final_project_hacktiv8/modules/comments/service"

	sosmed_ctrl "final_project_hacktiv8/modules/socialMedias/controller"
	sosmed_repo "final_project_hacktiv8/modules/socialMedias/repository"
	sosmed_serv "final_project_hacktiv8/modules/socialMedias/service"

	_ "final_project_hacktiv8/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func NewRouter(r *gin.Engine, db *gorm.DB) {

	// route user
	ctrlUser := user_ctrl.New(user_serv.New(user_repo.New(db)))
	routeUser := r.Group("/users")
	routeUser.POST("/register", ctrlUser.Create)
	routeUser.POST("/login", ctrlUser.Login)
	routeUser.PUT("", middleware.Authorization, ctrlUser.Update)
	routeUser.DELETE("", middleware.Authorization, ctrlUser.DeleteByID)

	// route photos
	ctrlPhoto := photo_ctrl.New(photo_serv.New(photo_repo.New(db)))
	r.POST("photos", middleware.Authorization, ctrlPhoto.Create)
	r.GET("photos", middleware.Authorization, ctrlPhoto.GetPhotos)
	r.PUT("photos/:photoID", middleware.Authorization, ctrlPhoto.Update)
	r.DELETE("photos/:photoID", middleware.Authorization, ctrlPhoto.Delete)

	// route comment
	ctrlComment := commect_ctrl.New(commect_serv.New(commect_repo.New(db)))
	r.POST("comments", middleware.Authorization, ctrlComment.Create)
	r.GET("comments", middleware.Authorization, ctrlComment.Get)
	r.PUT("comments/:commentID", middleware.Authorization, ctrlComment.Update)
	r.DELETE("comments/:commentID", middleware.Authorization, ctrlComment.Delete)

	// route social media
	repoSocialmedia := sosmed_repo.New(db)
	repoPhoto := photo_repo.New(db)
	ctrlSocialmedia := sosmed_ctrl.New(sosmed_serv.New(repoSocialmedia, repoPhoto))
	routerSocialmedia := r.Group("/socialmedias")
	routerSocialmedia.POST("", middleware.Authorization, ctrlSocialmedia.Create)
	routerSocialmedia.GET("", middleware.Authorization, ctrlSocialmedia.GetList)
	routerSocialmedia.PUT("/:socialmediaid", middleware.Authorization, ctrlSocialmedia.UpdateByID)
	routerSocialmedia.DELETE("/:socialmediaid", middleware.Authorization, ctrlSocialmedia.DeleteByID)

	// routing docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
