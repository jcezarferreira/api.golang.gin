package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jcezarferreira/api.gin/database"
	"github.com/jcezarferreira/api.gin/models"
)

func FindAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(200, students)
}

func FindStudentById(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"ResourceNotFound": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func RemoveStudentById(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student removed successfull",
	})
}

func EditStudent(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidateData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)

}

func FindStudentByCPF(c *gin.Context) {
	var student models.Student

	cpf := c.Param("cpf")

	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"ResourceNotFound": "Student not found",
		})
		return
	}

	c.JSON(http.StatusOK, student)
}

func AddStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := models.ValidateData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)

}

func Greetings(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API says": "Hi " + name + ", whats up?",
	})
}
