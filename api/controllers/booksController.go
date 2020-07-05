package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/NiteshStha/go-books-api/api/models"
	"github.com/NiteshStha/go-books-api/api/responses"
	"github.com/gorilla/mux"
)

// GetBooks returns all books
func (a *App) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetBooks(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, books)
	return
}

// GetBookByID returns a single book from the id provided
func (a *App) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	book, err := models.GetBookByID(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, book)
	return
}

// CreateBooks creates a new book
func (a *App) CreateBooks(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Book created succesfully"}

	book := &models.Book{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = json.Unmarshal(body, &book)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	book.Prepare() // strips all the white space
	if err = book.Validate(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	if bk, _ := book.GetBook(a.DB); bk != nil {
		resp["status"] = "failed"
		resp["message"] = "Book already registered"
		responses.JSON(w, http.StatusBadRequest, resp)
		return
	}

	bookCreated, err := book.Save(a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["book"] = bookCreated
	responses.JSON(w, http.StatusCreated, resp)
	return
}

// UpdateBooks updates a book
func (a *App) UpdateBooks(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Book updated successfully"}
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bookUpdate := models.Book{}
	if err = json.Unmarshal(body, &bookUpdate); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	bookUpdate.Prepare()
	if err = bookUpdate.Validate(); err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	_, err = bookUpdate.UpdateBook(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	resp["book"] = bookUpdate
	responses.JSON(w, http.StatusOK, resp)
	return
}

// DeleteBooks deletes a book
func (a *App) DeleteBooks(w http.ResponseWriter, r *http.Request) {
	var resp = map[string]interface{}{"status": "success", "message": "Book deleted successfully"}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	_, err := models.GetBookByID(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	err = models.DeleteBook(id, a.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, resp)
	return
}
