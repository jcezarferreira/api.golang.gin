package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jcezarferreira/api.gin/controllers"
	"github.com/jcezarferreira/api.gin/database"
	"github.com/jcezarferreira/api.gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRoutesTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() models.Student {
	database.ConnectDB()

	student := models.Student{Name: "Julio Ferreira", CPF: "08297000486", RG: "7791479"}

	database.DB.Create(&student)

	ID = int(student.ID)

	return student
}

func DeleteStudentMock() {
	var student models.Student

	database.DB.Delete(&student, ID)
}

func TestCheckStatusCodeGreetingsWithParams(t *testing.T) {
	r := SetupRoutesTest()

	r.GET("/:name", controllers.Greetings)

	req, _ := http.NewRequest("GET", "/julio", nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	mockResponse := `{"API says":"Hi julio, whats up?"}`

	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, mockResponse, string(resBody))

}

func TestFindAllStudents(t *testing.T) {
	database.ConnectDB()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupRoutesTest()

	r.GET("/students", controllers.FindAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestFindByCPF(t *testing.T) {
	database.ConnectDB()
	studentMock := CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupRoutesTest()

	r.GET("/students/cpf/:cpf", controllers.FindStudentByCPF)

	req, _ := http.NewRequest("GET", "/students/cpf/"+studentMock.CPF, nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var studentResponse models.Student

	json.Unmarshal(res.Body.Bytes(), &studentResponse)

	assert.Equal(t, http.StatusOK, res.Code)

	assert.Equal(t, studentMock.CPF, studentResponse.CPF)
}

func TestFindByID(t *testing.T) {
	database.ConnectDB()
	studentMock := CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupRoutesTest()

	r.GET("/students/:id", controllers.FindStudentById)

	req, _ := http.NewRequest("GET", "/students/"+strconv.FormatUint(uint64(studentMock.ID), 10), nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	var studentResponse models.Student

	json.Unmarshal(res.Body.Bytes(), &studentResponse)

	assert.Equal(t, http.StatusOK, res.Code)

	assert.Equal(t, studentMock.ID, studentResponse.ID)
	assert.Equal(t, studentMock.Name, studentResponse.Name)
	assert.Equal(t, studentMock.RG, studentResponse.RG)
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDB()
	studentMock := CreateStudentMock()

	r := SetupRoutesTest()

	r.DELETE("/students/:id", controllers.RemoveStudentById)

	req, _ := http.NewRequest("DELETE", "/students/"+strconv.FormatUint(uint64(studentMock.ID), 10), nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	mockResponse := `{"message":"Student removed successfull"}`

	assert.Equal(t, http.StatusOK, res.Code)

	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, mockResponse, string(resBody))

}
