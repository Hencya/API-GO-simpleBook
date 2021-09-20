package Controller

import (
	Database "Book_CRUD/lib/database"
	Models "Book_CRUD/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateBookController(ctx echo.Context)error{
	books,err := Database.PostBook(ctx)
	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"Message": "success create new book",
		"Data": books,
	})
}

func GetAllBooksController(ctx echo.Context)error{
	books,err := Database.GetBooks()

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound,err.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success get all books",
		"data": books,
	})
}

func GetBookByIdController(ctx echo.Context)error{
	id,err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound,err.Error())
	}

	book,errGet := Database.GetBookById(id)

	if errGet != nil{
		return echo.NewHTTPError(http.StatusInternalServerError,errGet.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success get all books",
		"data": book,
	})
}

func EditBookByIdController(ctx echo.Context)error{
	id,err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound,err.Error())
	}

	var book Models.Books
	ctx.Bind(&book)
	result,errUpdate := Database.PutBookById(id,book)

	if errUpdate != nil{
		return echo.NewHTTPError(http.StatusInternalServerError,errUpdate.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success get all books",
		"data": result,
	})
}

func DeleteBookByIdController(ctx echo.Context)error{
	id,err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound,err.Error())
	}

	_,errDelete := Database.DeleteBookById(id)

	if errDelete != nil{
		return echo.NewHTTPError(http.StatusInternalServerError,err.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success delete book",
	})

}
