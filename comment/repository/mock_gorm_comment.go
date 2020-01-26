package repository

import (
	"github.com/TenaHub/api/entity"
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/comment"
	"errors"
)

// CommentGormRepo implements comment.CommentRepository
type MockCommentGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo creates object of CommentGormRepo
func NewMockCommentGormRepo(conn *gorm.DB) comment.CommentRepository {
	return &MockCommentGormRepo{conn: conn}
}

// Comments returns all health center comments from database
func (cr *MockCommentGormRepo) Comments(id uint) ([]entity.Comment, []error) {
	// comments := []entity.Comment{}
	usercmt := []entity.Comment{}
	usercmt = append(usercmt, entity.MockComment, entity.MockComment)
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return usercmt, nil
}

// Comment returns single healthcenter comment from database
func (cr *MockCommentGormRepo) Comment(id uint) (*entity.Comment, []error) {
	comment := entity.MockComment
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &comment, nil
}

// UpdateComment updates comment from the database
func (cr *MockCommentGormRepo) UpdateComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmt := comment
	return cmt, nil
}

// StoreComment stores comment to the database
func (cr *MockCommentGormRepo) StoreComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmt := comment
	return cmt, nil
}

// DeleteComment deletes single comment from the database
func (cr *MockCommentGormRepo) DeleteComment(id uint) (*entity.Comment, []error) {
	comment := entity.MockComment
	if id != 1 {
		return nil, []error{errors.New("Not found")}
	}
	return &comment, nil
}

// CheckUser checks if user is valid to give feedback
func (cr *MockCommentGormRepo) CheckUser(cmt *entity.Comment) []error {
	return nil
}
