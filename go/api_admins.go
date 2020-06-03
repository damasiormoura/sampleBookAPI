/*
 * Simple Inventory API
 *
 * Sample Books API
 *
 * API version: 1.0.0
 * Contact: damasiormoura@gmail.com.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/http"
)

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	book := BookFromJSON(body)
	isbn, created := CreateBook(book)
	if created {
		w.Header().Add("Location", "/damasiormoura/sample_books/1.0.0/books/"+isbn)
		w.WriteHeader(http.StatusCreated)
	} else {
		w.WriteHeader(http.StatusConflict)
	}
}
