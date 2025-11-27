package repository

import (
	"context"
	"log"

	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
)

type MenuRepository interface {
	GetMenuByID(role string) ([]domain.Menu, error)
	GetMenuByUserID(ctx context.Context, userID string) ([]domain.Menu, error)
}

type menuRepository struct {
	dao dao.GORMMenuDAO
}

func NewMenuRepository(dao dao.GORMMenuDAO) MenuRepository {
	return &menuRepository{
		dao: dao,
	}
}

func (repo *menuRepository) GetMenuByUserID(ctx context.Context, id string) ([]domain.Menu, error) {
	user, err := repo.dao.FindUserByID(ctx, id)
	menus, err := repo.dao.FindMenusByRole(ctx, user.Role)
	if err != nil {
		log.Println("repo Get Menus By Role Error", err)
	}
	return menus, err
}

func (repo *menuRepository) GetMenuByID(id string) (menu []domain.Menu, err error) {

	temp, err := repo.dao.GetMenuListByRole(id)
	if err != nil {
		return nil, err
	}
	for _, v := range temp {
		menu = append(menu, repo.toDomain(v))
	}

	return menu, nil
}

func (repo *menuRepository) toDomain(m dao.Menu) domain.Menu {
	return domain.Menu{
		ID:       m.ID,
		Name:     m.Name,
		Path:     m.Path,
		ParentID: m.ParentId,
		OrderNo:  m.OrderNo,
	}
}
