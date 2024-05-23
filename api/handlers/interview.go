package handlers

import (
	"hrplatform/models"
	"hrplatform/postgres"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InterviewHandler struct {
	InterviewRepository postgres.InterviewRepository
}

func (h *InterviewHandler) CreateInterview(c *gin.Context) {
	var interviewCreate models.CreateInterview
	if err := c.BindJSON(&interviewCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	interview, err := h.InterviewRepository.CreateInterview(interviewCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, interview)
}
// http://localhost:8080/interviews/
// {
//     "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//     "vacancy_id": "17db725a-6627-4226-b564-90a75f3a0f11",
//     "recruiter_id": "d83f27ea-c5fe-485b-b9f9-58e07be915ac",
//     "interview_date": "2024-12-21 10:00:00"
// }
// n10=> SELECT position FROM resumes WHERE user_id = 'ac7b5e32-9f18-445c-81cf-a2422457964c';
//      position      
// -------------------
//  Software Engineer
// (1 row)
// n10=> SELECT position FROM vacancies WHERE id = '17db725a-6627-4226-b564-90a75f3a0f11';
//      position     
// ------------------
//  Backend Engineer
// (1 row)
// UPDATE vacancies 
// SET position = 'Software Engineer'
// WHERE id = '17db725a-6627-4226-b564-90a75f3a0f11';
// UPDATE 1

// {
//     "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//     "vacancy_id": "17db725a-6627-4226-b564-90a75f3a0f11",
//     "recruiter_id": "d83f27ea-c5fe-485b-b9f9-58e07be915ac",
//     "interview_date": "2024-12-21 10:00:00"
// }
// {
//     "id": "30a81c19-b0c4-4d55-a2b9-a17a5af60ff2",
//     "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//     "vacancy_id": "17db725a-6627-4226-b564-90a75f3a0f11",
//     "recruiter_id": "d83f27ea-c5fe-485b-b9f9-58e07be915ac",
//     "interview_date": "2024-12-21T10:00:00Z",
//     "created_at": "2024-05-21T09:55:05.127182309+05:00",
//     "updated_at": "2024-05-21T09:55:05.127182506+05:00",
//     "deleted_at": 0
// }

func (h *InterviewHandler) GetInterviewByID(c *gin.Context) {
	id := c.Param("id")
	interviewID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interview ID"})
		return
	}

	interview, err := h.InterviewRepository.GetInterviewByID(interviewID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Interview not found"})
		return
	}

	c.JSON(http.StatusOK, interview)
}
// http://localhost:8080/interviews/30a81c19-b0c4-4d55-a2b9-a17a5af60ff2
// {
//     "id": "30a81c19-b0c4-4d55-a2b9-a17a5af60ff2",
//     "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//     "vacancy_id": "17db725a-6627-4226-b564-90a75f3a0f11",
//     "recruiter_id": "d83f27ea-c5fe-485b-b9f9-58e07be915ac",
//     "interview_date": "2024-12-21T10:00:00Z",
//     "created_at": "2024-05-21T09:55:05.127182Z",
//     "updated_at": "2024-05-21T09:55:05.127183Z",
//     "deleted_at": 0
// }

func (h *InterviewHandler) GetInterviewsByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	interviews, err := h.InterviewRepository.GetInterviewsByUserID(parsedUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No interviews found for the specified user"})
		return
	}

	c.JSON(http.StatusOK, interviews)
}
// http://localhost:8080/interviews/user/ac7b5e32-9f18-445c-81cf-a2422457964c
// {
//     "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//     "vacancy_id": "17db725a-6627-4226-b564-90a75f3a0f11",
//     "recruiter_id": "d83f27ea-c5fe-485b-b9f9-58e07be915ac",
//     "interview_date": "2024-12-21 10:00:00"
// }
// [
//     {
//         "id": "30a81c19-b0c4-4d55-a2b9-a17a5af60ff2",
//         "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//         "vacancy_id": "17db725a-6627-4226-b564-90a75f3a0f11",
//         "recruiter_id": "d83f27ea-c5fe-485b-b9f9-58e07be915ac",
//         "interview_date": "2024-12-21T10:00:00Z",
//         "created_at": "2024-05-21T09:55:05.127182Z",
//         "updated_at": "2024-05-21T09:55:05.127183Z",
//         "deleted_at": 0
//     }
// ]


func (h *InterviewHandler) GetAllInterviews(c *gin.Context) {
	filter := make(map[string]interface{})

	if companyIDStr := c.Query("company_id"); companyIDStr != "" {
		companyID, err := uuid.Parse(companyIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID format"})
			return
		}
		filter["company_id"] = companyID
	}

	if position := c.Query("position"); position != "" {
		filter["position"] = position
	}

	if experienceStr := c.Query("experience"); experienceStr != "" {
		minExp, err := strconv.Atoi(experienceStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experience value"})
			return
		}
		filter["experience"] = minExp
	}

	interviews, err := h.InterviewRepository.GetAllInterviews(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, interviews)
}
// http://localhost:8080/interviews/

func (h *InterviewHandler) UpdateInterview(c *gin.Context) {
	id := c.Param("id")
	interviewID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interview ID"})
		return
	}

	var interviewUpdate models.UpdateInterview
	if err := c.BindJSON(&interviewUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	interviewUpdate.ID = interviewID

	if err := h.InterviewRepository.UpdateInterview(interviewUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Interview updated successfully"})
}
// http://localhost:8080/interviews/17db725a-6627-4226-b564-90a75f3a0f11
// {
//     "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//     "vacancy_id": "17db725a-6627-4226-b564-90a75f3a0f11",
//     "recruiter_id": "d83f27ea-c5fe-485b-b9f9-58e07be915ac",
//     "interview_date": "2024-12-21 14:00:00"
// }

func (h *InterviewHandler) DeleteInterview(c *gin.Context) {
	id := c.Param("id")
	interviewID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interview ID"})
		return
	}

	if err := h.InterviewRepository.DeleteInterview(interviewID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Interview deleted successfully"})
}

