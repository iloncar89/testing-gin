package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"testing-gin/domain"
	"testing-gin/model"
	"testing-gin/payload"
	"testing-gin/service"
)

type TestController struct {
	testService service.TestService
}

func NewTestController(service service.TestService) *TestController {
	return &TestController{
		testService: service,
	}
}

func (controller *TestController) Case1(context *gin.Context) {
	var response payload.GreetingResponse

	response.Greeting = "hello"

	context.JSON(http.StatusOK, &response)
}

func (controller *TestController) Case2(context *gin.Context) {
	body := payload.GreetingRequest{}

	if err := context.BindJSON(&body); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var response payload.GreetingResponse

	response.Greeting = "hello " + body.Name

	context.JSON(http.StatusOK, &response)
}

func (controller *TestController) Case3(context *gin.Context) {
	numberString := context.Param("number")
	number, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		context.String(http.StatusInternalServerError, "error")
	}
	result := controller.testService.CalculateFibonacci(number)

	var response payload.CalculateFibonacciResponse

	response.Number = result

	context.JSON(http.StatusOK, &response)

}

func (controller *TestController) Case4(context *gin.Context) {
	body := payload.PersonRequestBody{}

	if err := context.BindJSON(&body); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var person model.Person

	person.FirstName = body.FirstName
	person.LastName = body.LastName
	person.YearOfBirth = body.YearOfBirth

	person, err := controller.testService.CreateGetDeletePersonTestCase(person)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var personResponse payload.PersonResponseBody

	personResponse.Id = person.Id
	personResponse.FirstName = person.FirstName
	personResponse.LastName = person.LastName
	personResponse.YearOfBirth = person.YearOfBirth

	context.JSON(http.StatusOK, &personResponse)
}

func (controller *TestController) Case5(context *gin.Context) {
	body := payload.PersonRequestBody{}

	if err := context.BindJSON(&body); err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var person domain.Person

	person.FirstName = body.FirstName
	person.LastName = body.LastName
	person.YearOfBirth = body.YearOfBirth

	person, err := controller.testService.CreateGetDeletePersonORMTestCase(person)
	if err != nil {
		context.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var personResponse payload.PersonResponseBody

	personResponse.Id = person.Id
	personResponse.FirstName = person.FirstName
	personResponse.LastName = person.LastName
	personResponse.YearOfBirth = person.YearOfBirth

	context.JSON(http.StatusOK, &personResponse)
}
