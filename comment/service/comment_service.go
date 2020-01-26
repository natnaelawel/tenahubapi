package repository

import (
	"github.com/TenaHub/api/comment"
	"github.com/TenaHub/api/entity"
)

// CommentService implements comment.CommentRepository
type CommentService struct {
	cmtRepo comment.CommentRepository
}
// NewCommentService creates object of CommentService
func NewCommentService(repo comment.CommentRepository) *CommentService {
	return &CommentService{cmtRepo: repo}
}

// Comments returns all health center comments
func (cs *CommentService) Comments(id uint) ([]entity.Comment, []error) {
	comments, errs := cs.cmtRepo.Comments(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return comments, nil
}

// Comment returns single healthcenter comment
func (cs *CommentService) Comment(id uint) (*entity.Comment, []error) {
	comment, errs := cs.cmtRepo.Comment(id)

	if len(errs) > 0 {
		return nil, errs
	}
	return comment, nil
}

// UpdateComment updates comment
func (cs *CommentService) UpdateComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmt, errs := cs.cmtRepo.UpdateComment(comment)

	if len(errs) > 0 {
		return nil, errs
	}
	return cmt, nil
}

// StoreComment stores comment
func (cs *CommentService) StoreComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmt, errs := cs.cmtRepo.StoreComment(comment)

	if len(errs) > 0 {
		return nil, errs
	}

	return cmt, nil
}

// DeleteComment deletes single comment
func (cs *CommentService) DeleteComment(id uint) (*entity.Comment, []error) {
	comment, errs := cs.cmtRepo.DeleteComment(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return comment, nil
}

// CheckUser checks if user is valid to give feedback
func (cs *CommentService) CheckUser(cmt *entity.Comment) []error {
	errs := cs.cmtRepo.CheckUser(cmt)
	if len(errs) > 0 {
		return errs
	}
	return nil
}
