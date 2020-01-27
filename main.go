package main

import (
	"fmt"
	// "github.com/tenahubapi/entity"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/natnaelawel/tenahubapi/delivery/http/handler"
	"github.com/natnaelawel/tenahubapi/entity"
	"os"

	hcserviceRepository "github.com/natnaelawel/tenahubapi/service/repository"
	hcserviceService "github.com/natnaelawel/tenahubapi/service/service"

	commentRepository "github.com/natnaelawel/tenahubapi/comment/repository"
	commentService "github.com/natnaelawel/tenahubapi/comment/service"

	ratingRepository "github.com/natnaelawel/tenahubapi/rating/repository"
	ratingService "github.com/natnaelawel/tenahubapi/rating/service"

	sesRepository "github.com/natnaelawel/tenahubapi/session/repository"
	sesService "github.com/natnaelawel/tenahubapi/session/service"

	// serviceRepo "github.com/natnaelawel/tenahubapi/service/repository"
	// serviceServ "github.com/natnaelawel/tenahubapi/service/service"
	adminRepo "github.com/natnaelawel/tenahubapi/admin/repository"
	adminServ "github.com/natnaelawel/tenahubapi/admin/service"
	agentRepo "github.com/natnaelawel/tenahubapi/agent/repository"
	agentServ "github.com/natnaelawel/tenahubapi/agent/service"
	healthCenterRepo "github.com/natnaelawel/tenahubapi/healthcenter/repository"
	healthCenterServ "github.com/natnaelawel/tenahubapi/healthcenter/service"
	//userRepo "github.com/natnaelawel/tenahubapi/user/repository"
	//userServ "github.com/natnaelawel/tenahubapi/user/service"
	feedBackRepo "github.com/natnaelawel/tenahubapi/comment/repository"
	feedBackServ "github.com/natnaelawel/tenahubapi/comment/service"
	"github.com/natnaelawel/tenahubapi/user/repository"
	"github.com/julienschmidt/httprouter"
	"github.com/natnaelawel/tenahubapi/user/service"
	"net/http"
	"github.com/jinzhu/gorm"
)


func main()  {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:0912345678@localhost/tenahub?sslmode=disable")
	//dbconn, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	errs := dbconn.CreateTable(&entity.Admin{}, &entity.Agent{}, &entity.Comment{}, &entity.Hcrating{}, &entity.HealthCenter{}, &entity.Rating{},&entity.Service{}, &entity.Session{}, &entity.User{},&entity.UserComment{}).GetErrors()
	fmt.Println(errs)
	
	if len(errs) > 0 {
		panic(errs)
	}

	userRepo := repository.NewUserGormRepo(dbconn)
	userServ := service.NewUserService(userRepo)
	userHandl := handler.NewUserHander(userServ)

	comRepo := commentRepository.NewCommentGormRepo(dbconn)
	comServ := commentService.NewCommentService(comRepo)
	cmtHandl := handler.NewCommentHandler(comServ)

	ratingRepo := ratingRepository.NewGormRatingRepository(dbconn)
	ratingServ := ratingService.NewHcRatingService(ratingRepo)
	ratingHandl := handler.NewRatingHandler(ratingServ)

	sessionRepo := sesRepository.NewSessionGormRepo(dbconn)
	sessionService := sesService.NewSessionService(sessionRepo)
	sesHandl := handler.NewSessionHandler(sessionService)

	adminRespository := adminRepo.NewAdminGormRepo(dbconn)
	adminService := adminServ.NewAdminService(adminRespository)
	adminHandler := handler.NewAdminHandler(adminService)

	agentRespository := agentRepo.NewAgentGormRepo(dbconn)
	agentService := agentServ.NewAgentService(agentRespository)
	agentHandler := handler.NewAgentHandler(agentService)


	healthCenterRespository := healthCenterRepo.NewHealthCenterGormRepo(dbconn)
	healthCenterService := healthCenterServ.NewHealthCenterService(healthCenterRespository)
	healthCenterHandler := handler.NewHealthCenterHandler(healthCenterService)

	feedBackRepository := feedBackRepo.NewCommentGormRepo(dbconn)
	feedBackService := feedBackServ.NewCommentService(feedBackRepository)
	feedBackHandler := handler.NewCommentHandler(feedBackService)

	serviceRepo := hcserviceRepository.NewServiceGormRepo(dbconn)
	serviceServ := hcserviceService.NewServiceService(serviceRepo)
	serviceHandler := handler.NewServiceHandler(serviceServ)

	////////////////

	// defer dbconn.Close()

	//errs := dbconn.CreateTable(&entity.HealthCenter{}).GetErrors()
	//if len(errs)> 0 {
	//	panic(errs)
	//}else {
	//	fmt.Println("something is occurred")
	//}

	router := httprouter.New()

	router.GET("/v1/admin/:id", adminHandler.GetSingleAdmin)
	router.POST("/v1/admin", adminHandler.GetAdmin)
	router.PUT("/v1/admin/:id", adminHandler.PutAdmin)
	router.POST("/v1/admins",adminHandler.PostAdmin)
	router.DELETE("/v1/admin",adminHandler.DeleteAdmin)
	router.GET("/v1/agent/:id", agentHandler.GetSingleAgent)

	router.GET("/v1/agent", agentHandler.GetAgents)
	router.PUT("/v1/agents/:id", agentHandler.PutAgent)
	router.POST("/v1/agents", agentHandler.PostAgent)
	router.POST("/v1/agent", agentHandler.GetAgent)
	router.OPTIONS("/v1/agent", agentHandler.PostAgent)
	router.DELETE("/v1/agent/:id", agentHandler.DeleteAgent)

	// router.GET("/v1/healthcenter/:id", healthCenterHandler.GetSingleHealthCenter)
	// router.POST("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	// router.PUT("/v1/healthcenter/:id", healthCenterHandler.PutHealthCenter)
	// router.GET("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	// router.GET("/v1/healthcenters", healthCenterHandler.GetHealthCenters)
	// router.DELETE("/v1/healthcenter/:id", healthCenterHandler.DeleteHealthCenter)

	//router.GET("/v1/user/:id", userHandler.GetSingleUser)
	//router.GET("/v1/user", userHandler.GetUsers)
	//router.DELETE("/v1/user/:id", userHandler.DeleteUser)

	//router.GET("/v1/service/:id", serviceHandler.GetSingleService)
	router.GET("/v1/services/:id", serviceHandler.GetServices)
	router.GET("/v1/pending/services/:id", serviceHandler.GetPendingServices)
	router.PUT("/v1/service/:id", serviceHandler.PutService)
	router.POST("/v1/service", serviceHandler.PostService)
	router.OPTIONS("/v1/service", serviceHandler.PostService)
	router.DELETE("/v1/service/:id", serviceHandler.DeleteService)
	router.GET("/v1/service/:id", serviceHandler.GetSingleService)

	//router.GET("/v1/feedback/:id", feedBackHandler.GetComment)
	router.GET("/v1/feedback/:id", feedBackHandler.GetComments)
	router.PUT("/v1/feedback/:id", feedBackHandler.PutComment)
	router.POST("/v1/feedback", feedBackHandler.PostComment)
	router.OPTIONS("/v1/feedback", feedBackHandler.PostComment)
	router.DELETE("/v1/feedback/:id", feedBackHandler.DeleteComment)

	router.GET("/v1/users", userHandl.GetUsers)
	router.GET("/v1/users/:id", userHandl.GetSingleUser)
	router.POST("/v1/user", userHandl.GetUser)
	router.PUT("/v1/users/:id", userHandl.PutUser)
	router.POST("/v1/users", userHandl.PostUser)
	router.DELETE("/v1/users/:id", userHandl.DeleteUser)

	//router.GET("/v1/services/:id", hcservHandl.GetServices)
	//router.PUT("/v1/services/:id", hcservHandl.PutService)
	//router.DELETE("/v1/services/:id", hcservHandl.DeleteService)
	//router.POST("/v1/services", hcservHandl.PostService)

	router.GET("/v1/comments/:id", cmtHandl.GetComments)
	router.GET("/v1/comment/:id", cmtHandl.GetComment)
	router.PUT("/v1/comments/:id", cmtHandl.PutComment)
	router.DELETE("/v1/comments/:id", cmtHandl.DeleteComment)
	router.POST("/v1/comments", cmtHandl.PostComment)
	router.POST("/v1/comments/check", cmtHandl.Check)

	router.GET("/v1/healthcenter/:id", healthCenterHandler.GetSingleHealthCenter)
	router.POST("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	router.GET("/v1/healthcenter/:id/agent", healthCenterHandler.GetHealthCentersByAgentId)
	router.POST("/v1/healthcenter/addhealthcenter", healthCenterHandler.PostHealthCenter)
	router.PUT("/v1/healthcenter/:id", healthCenterHandler.PutHealthCenter)
	router.GET("/v1/healthcenter", healthCenterHandler.GetHealthCenter)
	router.GET("/v1/healthcenters", healthCenterHandler.GetHealthCenters)
	router.DELETE("/v1/healthcenter/:id", healthCenterHandler.DeleteHealthCenter)

	router.GET("/v1/healthcenters/search", healthCenterHandler.SearchHealthcenters)
	//router.GET("/v1/healthcenter/:id", healthCenterHandler.)
	router.GET("/v1/healthcenters/top/:amount", healthCenterHandler.GetTop)

	router.GET("/v1/rating/:id", ratingHandl.GetRating)
	router.POST("/v1/rating", ratingHandl.PostRating)

	router.GET("/v1/session", sesHandl.GetSession)
	router.POST("/v1/session", sesHandl.PostSession)
	router.DELETE("/v1/session/:uuid", sesHandl.DeleteSession)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), router)
	if err != nil {
		panic(err)
	}
	//http.ListenAndServe(":8181", router)
}

