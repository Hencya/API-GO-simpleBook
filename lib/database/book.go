package Database

import (
	"Book_CRUD/config"
	Models "Book_CRUD/models"
)

func PostBook(book *Models.Books)(*Models.Books,error){
	if err := config.DB.Create(&book).Error; err != nil{
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

func GetBookById(id int)(*Models.Books,error){
	var book *Models.Books

	if err := config.DB.Where("id = ?",id).First(&book).Error; err != nil{
		return nil,err
	}

	return book,nil
}

func PutBookById(id int,book *Models.Books)(*Models.Books,error){
	if err := config.DB.Where("id = ?",id).Updates(&book).Error; err != nil{
		return nil,err
	}
	return book,nil
}

func DeleteBookById(id int)(*Models.Books,error){
	var book *Models.Books

	if err := config.DB.Delete(&book,id).Error; err != nil{
		return nil,err
	}
	return book,nil
}