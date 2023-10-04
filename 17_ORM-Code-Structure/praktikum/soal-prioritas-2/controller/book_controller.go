package controller

import (
	"apps/config"
	"apps/model"
	"apps/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	var books []model.Book

	err := config.DB.Find(&books).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve book"))
	}

	if len(books) == 0 {
		return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Book data successfully retrieved", books))
}

func Show(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var book model.Book
	result := config.DB.First(&book, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve book"))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("Book data successfully retrieved", book))
}


func Store(c echo.Context) error {
	var book model.Book

	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	if err := config.DB.Create(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store book data"))
	}

	return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", book))
}

func Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var updatedBook model.Book

	if err := c.Bind(&updatedBook); err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
	}

	var existingBook model.Book
	result := config.DB.First(&existingBook, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve book"))
	}

	config.DB.Model(&existingBook).Updates(updatedBook)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Book data successfully updated", existingBook))
}

func Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
	}

	var existingBook model.Book
	result := config.DB.First(&existingBook, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve book"))
	}

	config.DB.Delete(&existingBook)

	return c.JSON(http.StatusOK, utils.SuccessResponse("Book data successfully deleted", nil))
}


