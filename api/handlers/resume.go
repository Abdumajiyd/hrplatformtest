package handlers

import (
	"hrplatform/models"
	"hrplatform/postgres"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResumeHandler struct {
	ResumeRepository postgres.ResumeRepository
}

func (h *ResumeHandler) CreateResume(c *gin.Context) {
	var resumeCreate models.CreateResume
	if err := c.BindJSON(&resumeCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resume, err := h.ResumeRepository.CreateResume(resumeCreate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resume)
}

// http://localhost:8080/resumes
// {
// 	"position": "Software Engineer",
// 	"experience": 5,
// 	"description": "Experienced in backend development with Go and Python.",
// 	"user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c"
// }
// {
//     "id": "12b00aef-1db6-4779-96ff-9ec9db55c1b1",
//     "position": "Software Engineer",
//     "experience": 5,
//     "description": "Experienced in backend development with Go and Python.",
//     "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//     "created_at": "2024-05-21T05:58:16.542313Z",
//     "updated_at": "2024-05-21T05:58:16.542313Z",
//     "deleted_at": 0
// }

func (h *ResumeHandler) GetResumeByID(c *gin.Context) {
	id := c.Param("id")
	resumeID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid resume ID"})
		return
	}

	resume, err := h.ResumeRepository.GetResumeByID(resumeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume not found"})
		return
	}

	c.JSON(http.StatusOK, resume)
}

// http://localhost:8080/resumes/12b00aef-1db6-4779-96ff-9ec9db55c1b1
// {
//     "id": "12b00aef-1db6-4779-96ff-9ec9db55c1b1",
//     "position": "Software Engineer",
//     "experience": 5,
//     "description": "Experienced in backend development with Go and Python.",
//     "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//     "created_at": "2024-05-21T05:58:16.542313Z",
//     "updated_at": "2024-05-21T05:58:16.542313Z",
//     "deleted_at": 0
// }

func (h *ResumeHandler) GetResumesByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	resumes, err := h.ResumeRepository.GetResumesByUserID(parsedUserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No resumes found for the specified user"})
		return
	}

	c.JSON(http.StatusOK, resumes)
}

// http://localhost:8080/resumes/6c3fd2cc-a684-49a1-8de8-5761438a7655
// [
//
//	{
//	    "id": "12b00aef-1db6-4779-96ff-9ec9db55c1b1",
//	    "position": "Software Engineer",
//	    "experience": 5,
//	    "description": "Experienced in backend development with Go and Python.",
//	    "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//	    "created_at": "2024-05-21T05:58:16.542313Z",
//	    "updated_at": "2024-05-21T05:58:16.542313Z",
//	    "deleted_at": 0
//	}
//
// ]

func (h *ResumeHandler) GetAllResumes(c *gin.Context) {
	filter := make(map[string]interface{})

	// Position filter (case-insensitive search)
	if position := c.Query("position"); position != "" {
		filter["position"] = position
	}

	// Minimum experience filter
	if minExpStr := c.Query("min_exp"); minExpStr != "" {
		minExp, err := strconv.Atoi(minExpStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid min_exp value"})
			return
		}
		filter["min_exp"] = minExp
	}

	// Fetch resumes using the repository method
	resumes, err := h.ResumeRepository.GetAllResumes(filter)
	if err != nil {
		log.Printf("Error fetching resumes: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch resumes"})
		return
	}

	c.JSON(http.StatusOK, resumes)
}

// http://localhost:8080/resumes

// [
//     {
//         "id": "af5e8269-52be-4b1a-80fa-ea9c97fdc8b8",
//         "position": "Software Engineer",
//         "experience": 2,
//         "description": "Expert in python",
//         "user_id": "536e9741-1f09-4474-b6ff-f4039a598098",
//         "user_name": "User2",
//         "user_email": "user2@example.com",
//         "created_at": "2024-05-20T01:41:33.76103Z",
//         "updated_at": "2024-05-20T01:41:33.76103Z",
//         "deleted_at": 0
//     },
//     {
//         "id": "6c3fd2cc-a684-49a1-8de8-5761438a7655",
//         "position": "Software Engineer",
//         "experience": 2,
//         "description": "Expert in python",
//         "user_id": "536e9741-1f09-4474-b6ff-f4039a598098",
//         "user_name": "User2",
//         "user_email": "user2@example.com",
//         "created_at": "2024-05-20T01:38:55.610102Z",
//         "updated_at": "2024-05-20T02:37:10.98015Z",
//         "deleted_at": 0
//     },
//     {
//         "id": "12b00aef-1db6-4779-96ff-9ec9db55c1b1",
//         "position": "Software Engineer",
//         "experience": 5,
//         "description": "Experienced in backend development with Go and Python.",
//         "user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
//         "user_name": "NewUser",
//         "user_email": "NewUser@example.com",
//         "created_at": "2024-05-21T05:58:16.542313Z",
//         "updated_at": "2024-05-21T05:58:16.542313Z",
//         "deleted_at": 0
//     }
// ]

// http://localhost:8080/resumes?position=Software&min_exp=3[
	// [
	// 	{
	// 		"id": "12b00aef-1db6-4779-96ff-9ec9db55c1b1",
	// 		"position": "Software Engineer",
	// 		"experience": 5,
	// 		"description": "Experienced in backend development with Go and Python.",
	// 		"user_id": "ac7b5e32-9f18-445c-81cf-a2422457964c",
	// 		"user_name": "NewUser",
	// 		"user_email": "NewUser@example.com",
	// 		"created_at": "2024-05-21T05:58:16.542313Z",
	// 		"updated_at": "2024-05-21T05:58:16.542313Z",
	// 		"deleted_at": 0
	// 	}
	// ]
func (h *ResumeHandler) UpdateResume(c *gin.Context) {
	id := c.Param("id")

	var resumeUpdate models.UpdateResume
	if err := c.BindJSON(&resumeUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resumeUpdate.ID = id

	if err := h.ResumeRepository.UpdateResume(resumeUpdate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resume updated successfully"})
}
// http://localhost:8080/resumes/12b00aef-1db6-4779-96ff-9ec9db55c1b1
// {
//     "message": "Resume updated successfully"
// }

// http://localhost:8080/resumes/6c3fd2cc-a684-49a1-8de8-5761438a7655
// {
//     "position": "Software Engineer",
//     "experience": 2,
//     "description": "Expert in python",
//     "user_id": "536e9741-1f09-4474-b6ff-f4039a598098"
// }
// {
//     "message": "Resume updated successfully"
// }

func (h *ResumeHandler) DeleteResume(c *gin.Context) {
	id := c.Param("id")

	if err := h.ResumeRepository.DeleteResume(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Resume deleted successfully"})
}
