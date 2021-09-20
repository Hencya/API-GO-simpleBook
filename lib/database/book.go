package Database

import (
	"Book_CRUD/config"
	Models "Book_CRUD/models"
	"github.com/labstack/echo/v4"
)

func PostBook(ctx echo.Context)(interface{},error){
	var book Models.Books
	ctx.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil{
		return nil,err
	}

	return  book,nil
}

func GetBooks()(interface{},error){
	var books []Models.Books

	if err := config.DB.Find(&books).Error;err != nil{
		return  nil,err
	}

	return books,nil
}

func GetBookById(id int)(interface{},error){
	var book Models.Books

	if err := config.DB.Where("id = ?",id).First(&book).Error; err != nil{
		return nil,err
	}

	return book,nil
}

func PutBookById(id int,book Models.Books)(interface{},error){
	if err := config.DB.Where("id = ?",id).Updates(&book).Error; err != nil{
		return nil,err
	}
	return book,nil
}

func DeleteBookById(id int)(interface{},error){
	var book Models.Books

	if err := config.DB.Where("id = ?",id).Delete(&book).Error; err != nil{
		return nil,err
	}
	return book,nil
}