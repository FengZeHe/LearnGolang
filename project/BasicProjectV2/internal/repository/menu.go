package repository

import (
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository/dao"
)

type MenuRepository interface {
	GetList() ([]domain.Menu, error)
}

type menuRepository struct {
	dao dao.GORMMenuDAO
}

func NewMenuRepository(dao dao.GORMMenuDAO) MenuRepository {
	return &menuRepository{
		dao: dao,
	}
}

func (repo *menuRepository) GetList() (menu []domain.Menu, err error) {
	temp, err := repo.dao.GetMenuList()
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
