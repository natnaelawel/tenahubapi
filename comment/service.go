package comment

import (
	"github.com/TenaHub/api/entity"
)


// CommentService is
type CommentService interface {
	Comment(id uint)(*entity.Comment, []error)
	Comments(id uint)([]entity.Comment, []error)
	CheckUser(cmt *entity.Comment)([]error)
	UpdateComment(comment *entity.Comment)(*entity.Comment, []error)
	StoreComment(comment *entity.Comment)(*entity.Comment, []error)
	DeleteComment(id uint)(*entity.Comment, []error)
}