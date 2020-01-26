package test_handler

import (
	"testing"
	"github.com/julienschmidt/httprouter"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"encoding/json"
	commentRepo "github.com/TenaHub/api/comment/repository"
	commentServ "github.com/TenaHub/api/comment/service"
	"github.com/TenaHub/api/entity"
	"github.com/TenaHub/api/delivery/http/handler"
	"reflect"
)

func TestComments(t *testing.T) {

	commentRepo := commentRepo.NewCommentGormRepo(nil)
	commentServ := commentServ.NewCommentService(commentRepo)
	commentHandler :=handler.NewCommentHandler(commentServ)

	mux := httprouter.New()
	mux.GET("/v1/comments", commentHandler.GetComments)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/comments")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}
	var mockComment []entity.Comment
	var comment []entity.Comment
	_ = json.Unmarshal(body, &comment)
	mockComment = append(mockComment, entity.MockComment, entity.MockComment)

	if !reflect.DeepEqual(mockComment, comment) {
		t.Errorf("want body to contain \n%q, but\n%q", mockComment, comment)

	}
}

func TestComment(t *testing.T) {

	commentRepo := commentRepo.NewCommentGormRepo(nil)
	commentServ := commentServ.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentServ)

	mux := httprouter.New()
	mux.GET("/v1/comments/:id", commentHandler.GetComments)
	ts := httptest.NewTLSServer(mux)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/comments/1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}
	var mockComment entity.Comment
	var comment entity.Comment
	_ = json.Unmarshal(body, &comment)
	mockComment = entity.MockComment
	if !reflect.DeepEqual(mockComment, comment) {
		t.Errorf("want body to contain \n%q, but\n%q", mockComment, comment)

	}
}

