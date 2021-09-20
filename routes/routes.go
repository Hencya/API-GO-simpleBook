package Routes

import (
	Controller "Book_CRUD/controllers"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	echoRoutes := echo.New()

	echoRoutes.POST("/book",Controller.CreateBookController)
	echoRoutes.GET("/books",Controller.GetAllBooksController)
	echoRoutes.GET("/book/:id",Controller.GetBookByIdController)
	echoRoutes.PUT("/book/:id",Controller.EditBookByIdController)
	echoRoutes.DELETE("/book/:id",Controller.DeleteBookByIdController)

	return echoRoutes
}
