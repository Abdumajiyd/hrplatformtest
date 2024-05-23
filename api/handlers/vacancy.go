package handlers

import (
	"hrplatform/models"
	"hrplatform/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VacancyHandler struct {
	VacancyRepository postgres.VacancyRepository
}

func (h *VacancyHandler) CreateVacancy(c *gin.Context) {
	var vacancyCreate models.CreateVacancy
	if err := c.BindJSON(&vacancyCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	if vacancyCreate.Name == "" || vacancyCreate.Position == "" || vacancyCreate.Description == "" || vacancyCreate.CompanyID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Majburiy maydonlar yetishmayapti"})
		return
	}

	if vacancyCreate.MinExp < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Minimal tajriba noldan kam bo'lishi mumkin emas"})
		return
	}

	exists, err := h.VacancyRepository.CheckCompanyExists(vacancyCreate.CompanyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kompaniya mavjudligini tekshirib bolmadi"})
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Berilgan ID ga ega kompaniya mavjud emas"})
		return
	}

	vacancy, err := h.VacancyRepository.CreateVacancy(vacancyCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vacancy"})
		return
	}

	c.JSON(http.StatusCreated, vacancy)
}

// http://localhost:8080/vacancies
//
//	{
//	    "name": "Senior Backend Engineer",
//	    "position": "Backend Engineer",
//	    "min_exp": 5,
//	    "company_id": "8568841e-de0a-4ff6-8ff6-c0d298799b03",
//	    "description": "We're looking for a talented backend engineer..."
//	}
//
//	{
//	    "id": "9a95ebb5-4ef1-422b-b2a0-a8336d611f8a",
//	    "name": "Senior Backend Engineer",
//	    "position": "Backend Engineer",
//	    "min_exp": 5,
//	    "company_id": "8568841e-de0a-4ff6-8ff6-c0d298799b03",
//	    "description": "We're looking for a talented backend engineer...",
//	    "created_at": "2024-05-20T11:07:04.768021206+05:00",
//	    "updated_at": "2024-05-20T11:07:04.7680213+05:00",
//	    "deleted_at": 0
//	}
func (h *VacancyHandler) GetVacancyByID(c *gin.Context) {
	id := c.Param("id")
	vacancyID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vacancy ID"})
		return
	}

	vacancy, err := h.VacancyRepository.GetVacancyByID(vacancyID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vacancy not found"})
		return
	}

	c.JSON(http.StatusOK, vacancy)
}

// http://localhost:8080/vacancies/9a95ebb5-4ef1-422b-b2a0-a8336d611f8a
// {
//     "id": "9a95ebb5-4ef1-422b-b2a0-a8336d611f8a",
//     "name": "Senior Backend Engineer",
//     "position": "Backend Engineer",
//     "min_exp": 5,
//     "company_id": "8568841e-de0a-4ff6-8ff6-c0d298799b03",
//     "description": "We're looking for a talented backend engineer...",
//     "created_at": "2024-05-20T11:07:04.768021Z",
//     "updated_at": "2024-05-20T11:07:04.768021Z",
//     "deleted_at": 0
// }

func (h *VacancyHandler) GetAllVacancies(c *gin.Context) {
	filter := make(map[string]interface{})

	if position := c.Query("position"); position != "" {
		filter["position"] = position
	}
	if minExp := c.Query("min_exp"); minExp != "" {
		filter["min_exp"] = minExp
	}
	if companyID := c.Query("company_id"); companyID != "" {
		parsedCompanyID, err := uuid.Parse(companyID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
			return
		}
		filter["company_id"] = parsedCompanyID
	}

	vacancies, err := h.VacancyRepository.GetAllVacancies(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vacancies)
}

// http://localhost:8080/vacancies/
// [
//     {
//         "id": "9a95ebb5-4ef1-422b-b2a0-a8336d611f8a",
//         "name": "Senior Backend Engineer",
//         "position": "Backend Engineer",
//         "min_exp": 5,
//         "company_id": "8568841e-de0a-4ff6-8ff6-c0d298799b03",
//         "description": "We're looking for a talented backend engineer...",
//         "created_at": "2024-05-20T11:07:04.768021Z",
//         "updated_at": "2024-05-20T11:07:04.768021Z",
//         "deleted_at": 0
//     }
// ]

func (h *VacancyHandler) UpdateVacancy(c *gin.Context) {
	id := c.Param("id")

	var vacancyUpdate models.UpdateVacancy
	if err := c.BindJSON(&vacancyUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vacancyUpdate.ID, _ = uuid.Parse(id)

	if err := h.VacancyRepository.UpdateVacancy(vacancyUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vacancy updated successfully"})
}

// http://localhost:8080/vacancies/9a95ebb5-4ef1-422b-b2a0-a8336d611f8a
/// {
//     "position": "junior engineer",
//     "min_exp": 6
// }
// {
//     "id": "9a95ebb5-4ef1-422b-b2a0-a8336d611f8a",
//     "name": "Senior Backend Engineer",
//     "position": "junior engineer",
//     "min_exp": 6,
//     "company_id": "8568841e-de0a-4ff6-8ff6-c0d298799b03",
//     "description": "We're looking for a talented backend engineer...",
//     "created_at": "2024-05-20T11:07:04.768021Z",
//     "updated_at": "2024-05-20T11:19:14.47031Z",
//     "deleted_at": 0
// }

func (h *VacancyHandler) DeleteVacancy(c *gin.Context) {
	id := c.Param("id")
	vacancyID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vacancy ID"})
		return
	}

	if err := h.VacancyRepository.DeleteVacancy(vacancyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vacancy deleted successfully"})
}
