package handlers 

import ( 
    // Kerakli kutubxonalarni import qilish
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
    "time"

    "hrplatform/models"      // Loyiha ichidagi modellar
    "hrplatform/postgres"   // Ma'lumotlar bazasi bilan ishlash uchun obyektlar

    "github.com/gin-gonic/gin"   // HTTP so'rovlarni ishlash uchun framework
    "github.com/google/uuid"     // Yagona IDlar generatsiya qilish uchun kutubxona
)

type RecruiterHandler struct {
    RecruiterRepository postgres.RecruiterRepository // Ma'lumotlar bazasi bilan ishlash uchun repository
}

func (h *RecruiterHandler) GetRecruiterByID(c *gin.Context) {
    id := c.Param("id") // URLdan IDni olish

    // Repository orqali yollanma xodimni topish
    recruiter, err := h.RecruiterRepository.GetRecruiterByID(id)
    if err != nil {
        // Agar topilmasa, xatolik qaytarish
        c.JSON(http.StatusNotFound, gin.H{"error": "Yollanma xodim topilmadi"})
        return
    }

    // Topilgan yollanma xodimni JSON formatida qaytarish
    c.JSON(http.StatusOK, recruiter)
}
// http://localhost:8080/recruiters
// {
//     "name": "Husan MUsa",
//     "email": "Husan99@example.com",
//     "phone_number": "123-456-7890",
//     "birthday": "1985-08-15",
//     "gender": "m",
//     "company_id": "ad277609-f698-489a-a744-ec3cb9e812ce"
// }
// {
//     "id": "3815c071-bdab-4c7c-ac39-c672b715030c",
//     "name": "Husan MUsa",
//     "email": "Husan99@example.com",
//     "phone_number": "123-456-7890",
//     "birthday": "1985-08-15T00:00:00Z",
//     "gender": "m",
//     "company_id": "ad277609-f698-489a-a744-ec3cb9e812ce",
//     "created_at": "2024-05-21T07:25:57.208529805+05:00",
//     "updated_at": "2024-05-21T07:25:57.208529895+05:00",
//     "deleted_at": 0
// }

func (h *RecruiterHandler) GetAllRecruiters(c *gin.Context) {
    // So'rov parametrlaridan filtrlarni olish (yoshi, jinsi, kompaniya IDsi)
    age, _ := strconv.Atoi(c.DefaultQuery("age", "0")) 
    gender := c.DefaultQuery("gender", "")
    companyID := c.DefaultQuery("company_id", "")

    // Repository orqali filtrlangan yollanma xodimlarni olish
    recruiters, err := h.RecruiterRepository.GetAllRecruiters(age, gender, companyID)
    if err != nil {
        // Agar xatolik bo'lsa, xato qaytarish
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Topilgan yollanma xodimlarni JSON formatida qaytarish
    c.JSON(http.StatusOK, recruiters)
}
func (h *RecruiterHandler) CreateRecruiter(c *gin.Context) {
    var recruiterCreate models.CreateRecruiter // Yangi yollanma xodim ma'lumotlari uchun model

    // JSON so'rovdan ma'lumotlarni o'qish va modelga yuklash
    if err := c.BindJSON(&recruiterCreate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Noto'g'ri JSON ma'lumot: " + err.Error()})
        return
    }

    // Majburiy maydonlarni tekshirish
    if recruiterCreate.Name == "" || recruiterCreate.Email == "" || recruiterCreate.PhoneNumber == "" || recruiterCreate.Gender == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Majburiy maydonlar to'ldirilmagan"})
        return
    }

    // Tug'ilgan sanani tekshirish va to'g'ri formatlash
    birthday, err := time.Parse("2006-01-02", recruiterCreate.Birthday)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Noto'g'ri sana formati. YYYY-MM-DD shaklida kiriting"})
        return
    }
    recruiterCreate.Birthday = birthday.Format(time.RFC3339)

    // Kompaniya mavjudligini tekshirish
    exists, err := h.RecruiterRepository.CheckCompanyExists(recruiterCreate.CompanyID)
    if err != nil || !exists {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Berilgan IDga ega kompaniya mavjud emas"})
        return
    }

    // Repository orqali yangi yollanma xodimni yaratish
    recruiter, err := h.RecruiterRepository.CreateRecruiter(recruiterCreate)
    if err != nil {
        // Xatoliklarni tekshirish va tegishli xabarlarni qaytarish
        if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
            c.JSON(http.StatusConflict, gin.H{"error": fmt.Sprintf("'%s' elektron pochtali yollanma xodim allaqachon mavjud", recruiterCreate.Email)})
        } else if strings.Contains(err.Error(), "violates foreign key constraint") {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Berilgan IDga ega kompaniya mavjud emas"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Yollanma xodimni yaratishda xatolik"})
            log.Println(err)
        }
        return
    }

    // Yaratilgan yollanma xodimni JSON formatida qaytarish
    c.JSON(http.StatusCreated, recruiter)
}
// http://localhost:8080/recruiters
// {
//     "name": "Husan MUsayev",
//     "email": "Husan99@example.com",
//     "phone_number": "123-456-7890",
//     "birthday": "2000-08-15",
//     "gender": "m",
//     "company_id": "ad277609-f698-489a-a744-ec3cb9e812ce"
// }
// {
//     "id": "d83f27ea-c5fe-485b-b9f9-58e07be915ac",
//     "name": "Husan MUsayev",
//     "email": "Husan99@example.com",
//     "phone_number": "123-456-7890",
//     "birthday": "2000-08-15T00:00:00Z",
//     "gender": "m",
//     "company_id": "ad277609-f698-489a-a744-ec3cb9e812ce",
//     "created_at": "2024-05-21T08:59:34.277727001+05:00",
//     "updated_at": "2024-05-21T08:59:34.277727065+05:00",
//     "deleted_at": 0
// }

func (h *RecruiterHandler) UpdateRecruiter(c *gin.Context) {
    id := c.Param("id") // URLdan IDni olish

    var recruiterUpdate models.UpdateRecruiter // Yangilanish uchun ma'lumotlar modeli

    // JSON so'rovdan ma'lumotlarni o'qish
    if err := c.BindJSON(&recruiterUpdate); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    recruiterUpdate.ID = uuid.MustParse(id) // IDni UUID formatiga o'tkazish

    // Tug'ilgan sanani tekshirish va formatlash (agar o'zgartirilgan bo'lsa)
    if recruiterUpdate.Birthday != nil {
        birthday, err := time.Parse("2006-01-02", *recruiterUpdate.Birthday)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Noto'g'ri sana formati"})
            return
        }
        formattedBirthday := birthday.Format(time.RFC3339)
        recruiterUpdate.Birthday = &formattedBirthday
    }

    // Kompaniya IDsi o'zgartirilgan bo'lsa, uning mavjudligini tekshirish
    if recruiterUpdate.CompanyID != nil {
        exists, err := h.RecruiterRepository.CheckCompanyExists(*recruiterUpdate.CompanyID)
        if err != nil || !exists {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Berilgan IDga ega kompaniya mavjud emas"})
            return
        }
    }

    // Repository orqali yollanma xodimni yangilash
    if err := h.RecruiterRepository.UpdateRecruiter(recruiterUpdate); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Muvaffaqiyatli yangilanganligi haqida xabar qaytarish
    c.JSON(http.StatusOK, gin.H{"message": "Yollanma xodim muvaffaqiyatli yangilandi"})
}
// http://localhost:8080/recruiters/d83f27ea-c5fe-485b-b9f9-58e07be915ac
// {
//     "name": "Husan MUsayev",
//     "email": "Husan99@example.com"
// }
// {
//     "message": "Recruiter updated successfully"
// }

func (h *RecruiterHandler) DeleteRecruiter(c *gin.Context) {
    id := c.Param("id") // URLdan IDni olish

    // Repository orqali yollanma xodimni o'chirish
    if err := h.RecruiterRepository.DeleteRecruiter(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Muvaffaqiyatli o'chirilganligi haqida xabar qaytarish
    c.JSON(http.StatusOK, gin.H{"message": "Recruiter deleted successfully"})
}

