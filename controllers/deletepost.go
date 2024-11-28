package controllers

import (
	"fmt"
	"net/http"

	"forumbackend/services"

	"github.com/gorilla/mux"
)

func DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	err := services.DeletePost(postID)
	if err != nil {
		if err == services.ErrNoPostFound {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Post with ID %s deleted successfully", postID)
}
