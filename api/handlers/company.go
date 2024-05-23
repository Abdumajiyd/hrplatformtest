package handlers

import (
	"net/http"

	"hrplatform/models"
	"hrplatform/postgres"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CompanyHandler struct {
	CompanyRepository postgres.CompanyRepository
}

func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var companyCreate models.CreateCompany
	if err := c.BindJSON(&companyCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company, err := h.CompanyRepository.CreateCompany(companyCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, company)
}

// http://localhost:8080/companies
// {
//     "name": "Google",
//     "location": "Tashkent",
//     "workers": 500
// }
// {
//     "id": "ad277609-f698-489a-a744-ec3cb9e812ce",
//     "name": "Google",
//     "location": "Tashkent",
//     "workers": 500,
//     "created_at": "2024-05-21T07:15:28.387133Z",
//     "updated_at": "2024-05-21T07:15:28.387133Z",
//     "deleted_at": 0
// }

func (h *CompanyHandler) GetCompanyByID(c *gin.Context) {
	id := c.Param("id")
	companyID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	company, err := h.CompanyRepository.GetCompanyByID(companyID.String())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}
// http://localhost:8080/companies/8568841e-de0a-4ff6-8ff6-c0d298799b03
// {
//     "id": "8568841e-de0a-4ff6-8ff6-c0d298799b03",
//     "name": "Najot Ta'lim",
//     "location": "Tashkent",
//     "workers": 500,
//     "created_at": "2024-05-20T05:52:59.996994Z",
//     "updated_at": "2024-05-20T05:52:59.996994Z",
//     "deleted_at": 0
// }
// GetAllCompanies handles GET requests to retrieve all companies.
func (h *CompanyHandler) GetAllCompanies(c *gin.Context) {
	companies, err := h.CompanyRepository.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companies)
}
// http://localhost:8080/companies
// [{
// "id": "dd8134d4-ba10-4064-b7da-fa9fc0da7347",
// "name": "Acme Corporation",
// "location": "Silicon Valley",
// "workers": 500,
// "created_at": "2024-05-20T05:52:16.91036Z",
// "updated_at": "2024-05-20T05:52:16.91036Z",
// "deleted_at": 0},
// {"id": "8568841e-de0a-4ff6-8ff6-c0d298799b03",
// "name": "Najot Ta'lim",
// "location": "Tashkent",
// "workers": 500,
// "created_at": "2024-05-20T05:52:59.996994Z",
// "updated_at": "2024-05-20T05:52:59.996994Z",
// "deleted_at": 0},
// {"id": "d6c95d8b-c36c-4265-bc53-6b05b5cab4ab",
// "name": "RealSoft",
// "location": "Tashkent",
// "workers": 120,
// "created_at": "2024-05-20T05:53:31.496829Z",
// "updated_at": "2024-05-20T05:53:31.496829Z",
// "deleted_at": 0},
// {"id": "106870ef-d698-415d-a476-73dc8edbc171",
// "name": "Abu Tech",
// "location": "Tashkent",
// "workers": 500,
// "created_at": "2024-05-20T05:56:09.182291Z",
// "updated_at": "2024-05-20T05:56:09.182291Z",
// "deleted_at": },
// {"id": "73f9a53c-eb74-477f-b0aa-62e1485c3b8c",
// "name": "Abu Tech",
// "location": "Tashkent",
// "workers": 500,
// "created_at": "2024-05-20T05:57:22.4134Z",
// "updated_at": "2024-05-20T05:57:22.4134Z",
// "deleted_at": 0}
// ]
// UpdateCompany handles PUT requests to update a company's information.
// UpdateCompany handles PUT requests to update a company's information.
func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	id := c.Param("id")

	var companyUpdate models.UpdateCompany
	if err := c.BindJSON(&companyUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyUpdate.ID = uuid.MustParse(id)

	if err := h.CompanyRepository.UpdateCompany(companyUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company updated successfully"})
}
// http://localhost:8080/companies/dd8134d4-ba10-4064-b7da-fa9fc0da7347
// {"name": "Udevs",
// "location": "Tashkent"}
func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	if err := h.CompanyRepository.DeleteCompany(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
