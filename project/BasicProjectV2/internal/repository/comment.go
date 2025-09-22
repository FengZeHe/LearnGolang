package repository

import "github.com/basicprojectv2/internal/repository/dao"

type commentRepository struct {
	dao dao.CommentDao
}

type CommentRepository interface {
}

func NewCommentRepository(dao dao.CommentDao) CommentRepository {
	return &commentRepository{
		dao: dao,
	}
}
