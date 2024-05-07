package handlers

import (
	"net/http"
	services "scalexmedia-assignment/pkg/usecase/interface"
	"scalexmedia-assignment/pkg/utils/models"
	"scalexmedia-assignment/pkg/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdminHandler struct {
	adminUseCase services.AdminUseCase
	adminMap     map[string]models.AdminStruct
}

func NewAdminHandler(usecase services.AdminUseCase, adminMap map[string]models.AdminStruct) *AdminHandler {
	return &AdminHandler{
		adminUseCase: usecase,
		adminMap:     adminMap,
	}
}

func (ad *AdminHandler) AdminLogin(c *gin.Context) {
	var adminLoginDetail models.AdminLoginRequest
	if err := c.ShouldBindJSON(&adminLoginDetail); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Details not in correct format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := validator.New().Struct(adminLoginDetail)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Constraints not satisfied", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}

	adminData, ok := ad.adminMap[adminLoginDetail.Email]
	if !ok || adminData.Password != adminLoginDetail.Password {
		errs := response.ClientResponse(http.StatusUnauthorized, "Invalid credentials", nil, "invalid")
		c.JSON(http.StatusUnauthorized, errs)
		return
	}

	loginDetails, err := ad.adminUseCase.AdminLogin(adminLoginDetail.Email, adminData.Name)
	if err != nil {
		errs := response.ClientResponse(http.StatusInternalServerError, "Failed to admin login", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errs)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Admin successfully logged in", loginDetails, nil)
	c.JSON(http.StatusOK, success)
}

func (ad *AdminHandler) ShowBooks(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "page number not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	countStr := c.DefaultQuery("count", "10")
	pageSize, err := strconv.Atoi(countStr)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "count in a page not in right format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errorRes)
		return
	}
	booksDetails, err := ad.adminUseCase.ShowBooks(page, pageSize)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Failed to get books details", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Successfully get books details", booksDetails, nil)
	c.JSON(http.StatusCreated, success)
}

func (ad *AdminHandler) AddBook(c *gin.Context) {
	var addBook models.AddBookRequest
	if err := c.ShouldBindJSON(&addBook); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := ad.adminUseCase.AddBook(addBook)
	if err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "Could not add book", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	success := response.ClientResponse(http.StatusCreated, "Successfully added book", nil, nil)
	c.JSON(http.StatusCreated, success)
}

func (ad *AdminHandler) DeleteBook(c *gin.Context) {
	var deleteBook models.BookDetails
	if err := c.ShouldBindJSON(&deleteBook); err != nil {
		errs := response.ClientResponse(http.StatusBadRequest, "fields provided are in wrong format", nil, err.Error())
		c.JSON(http.StatusBadRequest, errs)
		return
	}
	err := ad.adminUseCase.DeleteBook(deleteBook)
	if err != nil {
		errorRes := response.ClientResponse(http.StatusBadRequest, "could not delete the book", nil, err.Error())
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}
	success := response.ClientResponse(http.StatusOK, "Successfully deleted the book", nil, nil)
	c.JSON(http.StatusOK, success)

}
