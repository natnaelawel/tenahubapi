package repository

import (
	"fmt"

	"github.com/TenaHub/api/entity"
	"github.com/jinzhu/gorm"
	"github.com/TenaHub/api/comment"
)

// CommentGormRepo implements comment.CommentRepository
type CommentGormRepo struct {
	conn *gorm.DB
}

// NewCommentGormRepo creates object of CommentGormRepo
func NewCommentGormRepo(conn *gorm.DB) comment.CommentRepository {
	return &CommentGormRepo{conn: conn}
}

// Comments returns all health center comments from database
func (cr *CommentGormRepo) Comments(id uint) ([]entity.Comment, []error) {
	// comments := []entity.Comment{}
	usercmt := []entity.Comment{}
	errs := cr.conn.Table("comments").Select("comments.*, users.first_name").Joins("left join users on users.id = comments.user_id").Where("comments.health_center_id = ?", id).Scan(&usercmt).GetErrors()

	fmt.Println(usercmt)

	if len(errs) > 0 {
		return nil, errs
	}

	return usercmt, nil
}

// Comment returns single healthcenter comment from database
func (cr *CommentGormRepo) Comment(id uint) (*entity.Comment, []error) {
	comment := entity.Comment{}
	errs := cr.conn.First(&comment, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &comment, nil
}

// UpdateComment updates comment from the database
func (cr *CommentGormRepo) UpdateComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmt := comment

	errs := cr.conn.Save(cmt).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return cmt, nil
}

// StoreComment stores comment to the database
func (cr *CommentGormRepo) StoreComment(comment *entity.Comment) (*entity.Comment, []error) {
	cmt := comment

	errs := cr.conn.Create(cmt).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return cmt, nil
}

// DeleteComment deletes single comment from the database
func (cr *CommentGormRepo) DeleteComment(id uint) (*entity.Comment, []error) {
	comment, errs := cr.Comment(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cr.conn.Delete(&comment, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return comment, nil
}

// CheckUser checks if user is valid to give feedback
func (cr *CommentGormRepo) CheckUser(cmt *entity.Comment) []error {
	comment := cmt
	errs := cr.conn.Where("user_id = ? and health_center_id = ?", comment.UserID, comment.HealthCenterID).First(comment).GetErrors()
	if len(errs) > 0 {
		return errs
	}
	return nil
}
