package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/apichitlakorn/SA-BRAINBOOM/entity"
	"github.com/apichitlakorn/SA-BRAINBOOM/config"
	
)

// Assume DB is the global database connection
// GET /payments


// GET /Course
func ListCourse(c *gin.Context) {
	var Course []entity.Course
	if err := config.DB().Find(&Course).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch Course"})
			return
	}

	c.JSON(http.StatusOK, Course)
}


// GET /course-price/:id
func GetCoursePrice(c *gin.Context) {
	ID := c.Param("id")
	var price float64

	db := config.DB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	results := db.Model(&entity.Course{}).Select("price").Where("id = ?", ID).Scan(&price)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course_id": ID,
		"price": price,
	})
}

// GET /course-title/:id
func GetCourseName(c *gin.Context) {
	ID := c.Param("id")
	var title string

	db := config.DB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	results := db.Model(&entity.Course{}).Select("title").Where("id = ?", ID).Scan(&title)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"course_id": ID,
		"title": title,
	})
}

func CountCourses(c *gin.Context) {
	var count int64

	// ตรวจสอบการเชื่อมต่อกับฐานข้อมูล
	db := config.DB()
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}

	// ใช้ GORM เพื่อทำการนับจำนวน Course โดยอ้างอิงจาก id
	if err := db.Model(&entity.Course{}).Count(&count).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งจำนวน Course กลับไปในรูปแบบ JSON
	c.JSON(http.StatusOK, gin.H{
		"total_courses": count,
	})
}



// POST /payment
