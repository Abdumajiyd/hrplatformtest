package main

import (
	"hrplatform/api"
	"hrplatform/api/handlers"
	"hrplatform/config"
	"hrplatform/postgres"
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Konfiguratsiyani yuklash
	cfg := config.Load()

	// Malumotlar bazasiga ulani
	db, err := sqlx.Connect("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Ma'lumotlar bazasiga ulanishda xatolik: %v", err)
	}

	// Repositorylarni yaratish
	userRepo := &postgres.PostgresUserRepository{DB: db}
	resumeRepo := &postgres.PostgresResumeRepository{DB: db}
	recruiterRepo := &postgres.PostgresRecruiterRepository{DB: db}
	companyRepo := &postgres.PostgresCompanyRepository{DB: db}
	interviewRepo := &postgres.PostgresInterviewRepository{DB: db}
	vacancyRepo := &postgres.PostgresVacancyRepository{DB: db}

	// Handlerlarni yaratish
	userHandler := &handlers.UserHandler{UserRepository: userRepo}
	resumeHandler := &handlers.ResumeHandler{ResumeRepository: resumeRepo}
	recruiterHandler := &handlers.RecruiterHandler{RecruiterRepository: recruiterRepo}
	companyHandler := &handlers.CompanyHandler{CompanyRepository: companyRepo}
	interviewHandler := &handlers.InterviewHandler{InterviewRepository: interviewRepo}
	vacancyHandler := &handlers.VacancyHandler{VacancyRepository: vacancyRepo}

	// Gin routerni sozlash
	router := api.SetupRouter(userHandler, resumeHandler, recruiterHandler, companyHandler, interviewHandler, vacancyHandler)

	// Serverni ishga tushirish
	if err := router.Run(":" + cfg.HTTPPort); err != nil {
		log.Fatalf("Serverni ishga tushirishda xatolik: %v", err)
	}
}























































// package main

// import (
// 	"log"
// 	"net/http"

// 	"hrplatform/api"
// 	"hrplatform/api/handlers"
// 	"hrplatform/config"
// 	"hrplatform/postgres"
// )

// func main() {
// 	// Konfiguratsiyani yuklash
// 	cfg := config.Load()

// 	// Ma'lumotlar bazasiga ulanish
// 	db, err := cfg.ConnectDB()
// 	if err != nil {
// 		log.Fatalf("Ma'lumotlar bazasiga ulanishda xatolik yuz berdi: %v", err)
// 	}
// 	defer db.Close()
// 	// Repozitoriyalarini yaratish
// 	userRepo := &postgres.PostgresUserRepository{DB: db}
// 	resumeRepo := &postgres.PostgresResumeRepository{DB: db}
// 	recruiterRepo := &postgres.PostgresRecruiterRepository{DB: db}
// 	companyRepo := &postgres.PostgresCompanyRepository{DB: db}
// 	interviewRepo := &postgres.PostgresInterviewRepository{DB: db}
// 	vacancyRepo := &postgres.PostgresVacancyRepository{DB: db}

// 	// Handlerlarini yaratish
// 	userHandler := &handlers.UserHandler{UserRepository: userRepo}
// 	resumeHandler := &handlers.ResumeHandler{ResumeRepository: resumeRepo}
// 	recruiterHandler := &handlers.RecruiterHandler{RecruiterRepository: recruiterRepo}
// 	companyHandler := &handlers.CompanyHandler{CompanyRepository: companyRepo}
// 	interviewHandler := &handlers.InterviewHandler{InterviewRepository: interviewRepo}
// 	vacancyHandler := &handlers.VacancyHandler{VacancyRepository: vacancyRepo}

// 	// Serverni ishga tushirish
// 	router := api.SetupRouter(userHandler, resumeHandler, recruiterHandler, companyHandler, interviewHandler,vacancyHandler)
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }
