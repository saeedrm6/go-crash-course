package routes

import (
	"encoding/json"
	"golang-rest-api/entity"
	"golang-rest-api/repository"
	"math/rand"
	"net/http"
)



var (
	repo repository.PostRepository = repository.NewPostRepository()
)

// GetPosts /*
func GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error getting the posts"}`))
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	post.Id = rand.Int63()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}
