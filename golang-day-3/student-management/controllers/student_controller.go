package controllers

import (
	"log"
	"net/http"
	"student-management/config"
	"student-management/models"

	"github.com/gin-gonic/gin"
)

// Get all students
func GetStudents(c *gin.Context) {
	var students []models.Student

	if config.DB == nil {
		log.Fatal("Database connection is nil")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection failed"})
		return
	}

	result := config.DB.Find(&students)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, students)
}

// Create a new student
func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&student)
	c.JSON(http.StatusCreated, student)
}

// Update student by ID
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	// Check if the student exists
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Create a temporary struct to hold the new values
	var updatedStudent models.Student
	if err := c.ShouldBindJSON(&updatedStudent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update only the fields that are provided
	config.DB.Model(&student).Updates(updatedStudent)

	c.JSON(http.StatusOK, student)
}

// Get student by ID
func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	result := config.DB.First(&student, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// Delete student by ID
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Student{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}
