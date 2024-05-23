package handlers

import (
	"net/http"
	"strconv"
	"time"

	"hrplatform/models"
	"hrplatform/postgres"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserRepository postgres.UserRepository 
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userCreate models.UserCreate
	if err := c.BindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	birthday, err := time.Parse("2006-01-02", userCreate.Birthday)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}

	userCreate.Birthday = birthday.Format("2006-01-02")

	user, err := h.UserRepository.CreateUser(userCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// http://localhost:8080/users/
// {
//     "name": "Test10",
//     "email": "test@example.com",
//     "phone_number": "+1234567890",
//     "birthday": "2000-01-01",
//     "gender": "m"
// }

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := h.UserRepository.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// http://localhost:8080/users/41cf99a7-16f9-4256-98fc-9bb495a455e8

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	filters := make(map[string]interface{})

	if ageStr := c.Query("age"); ageStr != "" {
		age, err := strconv.Atoi(ageStr)
		if err == nil {
			filters["age"] = age
		}
	}

	if gender := c.Query("gender"); gender != "" {
		filters["gender"] = gender
	}

	users, err := h.UserRepository.GetAllUsers(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// http://localhost:8080/users
// http://localhost:8080/users?age=30&gender=m

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var userUpdate models.UserUpdate
	if err := c.BindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if userUpdate.Birthday != "" {
		_, err := time.Parse("2006-01-02", userUpdate.Birthday)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}
	}

	userUpdate.ID = uuid.MustParse(id)

	if err := h.UserRepository.UpdateUser(userUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// http://localhost:8080/users/41cf99a7-16f9-4256-98fc-9bb495a455e8
// {
// 	"name": "Update name",
// 	"email": "updatedemail@example.com",
// 	"phone_number": "+998901234567",
// 	"birthday": "1995-05-15",
// 	"gender": "f"
// }

// yoki faqat 1 ta field uchun
// {
// 	"name": "Update  Name"
// }

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	if err := h.UserRepository.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *UserHandler) GetUserInterviews(c *gin.Context) {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	interviews, err := h.UserRepository.GetUserInterviews(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Interviews not found"})
		return
	}

	c.JSON(http.StatusOK, interviews)
}

// http://localhost:8080/users/41cf99a7-16f9-4256-98fc-9bb495a455e8/myInterview

func (h *UserHandler) GetUserResume(c *gin.Context) {
	id := c.Param("id")
	userID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	resumes, err := h.UserRepository.GetUserResume(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resumes not found"})
		return
	}

	c.JSON(http.StatusOK, resumes)
}

// http://localhost:8080/users/41cf99a7-16f9-4256-98fc-9bb495a455e8/myresume
