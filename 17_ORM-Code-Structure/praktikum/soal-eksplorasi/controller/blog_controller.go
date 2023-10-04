package controller

import (
	"appsv/config"
	"appsv/model"
	"appsv/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
    var blogs []model.Blog

    err := config.DB.Preload("User").Find(&blogs).Error
    if err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve blog data"))
    }

    if len(blogs) == 0 {
        return c.JSON(http.StatusNotFound, utils.ErrorResponse("Empty data"))
    }

    return c.JSON(http.StatusOK, utils.SuccessResponse("Blog data successfully retrieved", blogs))
}

func Show(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
    }

    var blog model.Blog
    result := config.DB.First(&blog, id)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve blog"))
    }

    return c.JSON(http.StatusOK, utils.SuccessResponse("Blog data successfully retrieved", blog))
}

func Store(c echo.Context) error {
    var blog model.Blog

    if err := c.Bind(&blog); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    if err := config.DB.Create(&blog).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to store blog data"))
    }

    return c.JSON(http.StatusCreated, utils.SuccessResponse("Success Created Data", blog))
}

func Update(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
    }

    var updatedBlog model.Blog

    if err := c.Bind(&updatedBlog); err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid request body"))
    }

    var existingBlog model.Blog
    result := config.DB.First(&existingBlog, id)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve blog"))
    }

    config.DB.Model(&existingBlog).Updates(updatedBlog)

    return c.JSON(http.StatusOK, utils.SuccessResponse("Blog data successfully updated", existingBlog))
}

func Delete(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, utils.ErrorResponse("Invalid ID"))
    }

    var existingBlog model.Blog
    result := config.DB.First(&existingBlog, id)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to retrieve blog"))
    }

    config.DB.Delete(&existingBlog)

    return c.JSON(http.StatusOK, utils.SuccessResponse("Blog data successfully deleted", nil))
}


