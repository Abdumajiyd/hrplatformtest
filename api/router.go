package api

import (
	"hrplatform/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler,
    resumeHandler *handlers.ResumeHandler, 
    recruiterHandler *handlers.RecruiterHandler, 
    companyHandler *handlers.CompanyHandler, 
    interviewHandler *handlers.InterviewHandler,
    vacancyHandler *handlers.VacancyHandler) *gin.Engine {
	router := gin.Default()

	userGroup := router.Group("/users")
	{
		userGroup.POST("/", userHandler.CreateUser)
		userGroup.GET("/:id", userHandler.GetUserByID)
		userGroup.GET("/", userHandler.GetAllUsers)
		userGroup.PUT("/:id", userHandler.UpdateUser)
		userGroup.DELETE("/:id", userHandler.DeleteUser)
		userGroup.GET("/:id/myInterview", userHandler.GetUserInterviews)
		userGroup.GET("/:id/myresume", userHandler.GetUserResume)

	}

	resumeGroup := router.Group("/resumes")
	{
		resumeGroup.POST("/", resumeHandler.CreateResume)
		resumeGroup.GET("/:id", resumeHandler.GetResumeByID)
		resumeGroup.GET("/", resumeHandler.GetAllResumes)
		resumeGroup.PUT("/:id", resumeHandler.UpdateResume)
		resumeGroup.DELETE("/:id", resumeHandler.DeleteResume)
		// resumeGroup.GET("/user/:user_id", resumeHandler.GetResumesByUserID)
	}

	companyGroup := router.Group("/companies")
	{
		companyGroup.POST("/", companyHandler.CreateCompany)
		companyGroup.GET("/:id", companyHandler.GetCompanyByID)
		companyGroup.GET("/", companyHandler.GetAllCompanies)
		companyGroup.PUT("/:id", companyHandler.UpdateCompany)
		companyGroup.DELETE("/:id", companyHandler.DeleteCompany)
	}
	recruiterGroup := router.Group("/recruiters")
	{
		recruiterGroup.POST("/", recruiterHandler.CreateRecruiter)
		recruiterGroup.GET("/:id", recruiterHandler.GetRecruiterByID)
		recruiterGroup.GET("/", recruiterHandler.GetAllRecruiters)
		recruiterGroup.PUT("/:id", recruiterHandler.UpdateRecruiter)
		recruiterGroup.DELETE("/:id", recruiterHandler.DeleteRecruiter)
	}

	vacancyGroup := router.Group("/vacancies") 
	{
		vacancyGroup.POST("/", vacancyHandler.CreateVacancy)
		vacancyGroup.GET("/:id", vacancyHandler.GetVacancyByID)
		vacancyGroup.GET("/", vacancyHandler.GetAllVacancies)
		vacancyGroup.PUT("/:id", vacancyHandler.UpdateVacancy)
		vacancyGroup.DELETE("/:id", vacancyHandler.DeleteVacancy)
	}

	interviewGroup := router.Group("/interviews")
	{
		interviewGroup.POST("/", interviewHandler.CreateInterview)
		interviewGroup.GET("/:id", interviewHandler.GetInterviewByID)
		interviewGroup.GET("/", interviewHandler.GetAllInterviews)
		interviewGroup.PUT("/:id", interviewHandler.UpdateInterview)
		interviewGroup.DELETE("/:id", interviewHandler.DeleteInterview)
		// interviewGroup.GET("/user/:user_id", interviewHandler.GetInterviewsByUserID)
	}
	
	return router
}
