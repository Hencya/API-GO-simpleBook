package Controller

import (
	Database "Book_CRUD/lib/database"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func CreateBookController(ctx echo.Context)(err error){
	newBook := new(PostBookRequest)

	if err = ctx.Bind(newBook); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	result,errPost := Database.PostBook(PostBookRequestToModelRequest(newBook))

	if errPost != nil{
		fmt.Println("Masuk disini")
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"Message": "success create new book",
		"Data": ModelToBookDetailResponse(result),
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

	if err != nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	book,errGet := Database.GetBookById(id)

	if errGet != nil {
		return echo.NewHTTPError(http.StatusNotFound,errors.New("Not Found Book"))
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success get all books",
		"data": ModelToBookDetailResponse(book),
	})
}

func EditBookByIdController(ctx echo.Context) (err error){
	id,errConvert := strconv.Atoi(ctx.Param("id"))

	if errConvert != nil {
		return echo.NewHTTPError(http.StatusBadRequest,errConvert.Error())
	}

	updateBook := new(PutBookRequest)
	if err = ctx.Bind(updateBook); err!= nil{
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	result,errUpdate := Database.PutBookById(id,PutBookRequestToModelRequest(updateBook))

	if errUpdate != nil{
		return echo.NewHTTPError(http.StatusInternalServerError,errUpdate.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success get all books",
		"data": ModelToBookDetailResponse(result),
	})
}

func DeleteBookByIdController(ctx echo.Context)error{
	id,err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,err.Error())
	}

	_,errDelete := Database.DeleteBookById(id)

	if errDelete != nil{
		return echo.NewHTTPError(http.StatusNotFound,err.Error())
	}

	return ctx.JSON(http.StatusOK,map[string]interface{}{
		"message": "success delete book",
	})

}
