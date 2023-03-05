package test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	firstname := r.PostForm.Get("firstname")
	lastname := r.PostForm.Get("lastname")
	fmt.Fprintf(w, "%s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=Andrian&lastname=raihannudin")
	request := httptest.NewRequest("POST", "htt://localhost/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
