package service

import (
	"context"
	"errors"
	"github.com/basicprojectv2/internal/domain"
	"github.com/basicprojectv2/internal/repository"
	"github.com/casbin/casbin/v2"
	"log"
)

type SysService interface {
	GetMenuByUserID(ctx context.Context, userid string) (menus []domain.Menu, err error)
	GetMenuByRole(ctx context.Context, role string) (menus []domain.Menu, err error)

	GetApiByUserID(ctx context.Context, userid string) (apis []domain.API, err error)
	GetMenu(ctx context.Context) ([]domain.Menu, error)
	GetRole(ctx context.Context) ([]domain.Role, error)
	GetAPI(ctx context.Context) ([]domain.API, error)

	AddCasbinPolicy(ctx context.Context, req domain.AddCasbinRulePolicyReq) error
	UpdateCasbinPolicy(ctx context.Context, req domain.UpdateCasbinPolicyReq) error
	DeleteCasbinPolicy(ctx context.Context, req domain.RemoveCasbinPolicyReq) error
	UpdateCasbinPolicies(ctx context.Context, req domain.TransactionPolicyReq) error
}

type sysService struct {
	repo     repository.SysRepository
	enforcer *casbin.Enforcer
}

func NewSysService(repo repository.SysRepository, enforcer *casbin.Enforcer) SysService {
	return &sysService{repo: repo, enforcer: enforcer}
}

func (s *sysService) UpdateCasbinPolicies(ctx context.Context, req domain.TransactionPolicyReq) (err error) {
	log.Println(req.NewPolicies, req.OldPolicies)
	if len(req.NewPolicies) > 0 {
		for _, v := range req.NewPolicies {
			exists, _ := s.enforcer.HasPolicy(v)
			if !exists {
				ok, err := s.enforcer.AddPolicies(req.NewPolicies)
				if err != nil || !ok {
					log.Println(err, ok)
				}
			}
		}
	}

	if len(req.OldPolicies) > 0 {
		for _, v := range req.OldPolicies {
			exists, _ := s.enforcer.RemovePolicy(v)
			if !exists {
				ok, err := s.enforcer.RemovePolicies(req.OldPolicies)
				if err != nil || !ok {
					log.Println(ok, err)
				}
			}
		}

	}
	err = s.enforcer.SavePolicy()

	err = s.enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	return nil
}

func (s *sysService) GetMenuByRole(ctx context.Context, role string) (menus []domain.Menu, err error) {
	menus, err = s.repo.GetMenuByRole(ctx, role)
	if err != nil {
		return nil, err
	}
	return menus, nil
}

func (s *sysService) GetApiByUserID(ctx context.Context, userid string) (apis []domain.API, err error) {
	apis, err = s.repo.GetApiByUserID(ctx, userid)
	if err != nil {
		return nil, err
	}
	//处理用户返回菜单
	return apis, err
}

func (s *sysService) AddCasbinPolicy(ctx context.Context, policy domain.AddCasbinRulePolicyReq) (err error) {
	ok, err := s.enforcer.AddPolicy(policy.NewPolicy)
	if err != nil {
		return err
	}
	if !ok {
		return err
	}
	return nil
}

func (s *sysService) UpdateCasbinPolicy(ctx context.Context, req domain.UpdateCasbinPolicyReq) (err error) {
	// 先删除 再添加
	ok, err := s.enforcer.RemovePolicy(req.OldPolicy)
	if err != nil {
		log.Println("remove policy fail", err)
		return err
	}
	ok, err = s.enforcer.AddPolicy(req.NewPolicy)
	if err != nil {
		log.Println("add policy fail", err)
		return err
	}

	if ok {
		log.Println("update policy success")
		return nil
	}

	return nil
}

func (s *sysService) DeleteCasbinPolicy(ctx context.Context, req domain.RemoveCasbinPolicyReq) (err error) {
	ok, err := s.enforcer.RemovePolicy(req.RemovePolicy)
	if err != nil {
		log.Println("remove policy fail", err)
		return err
	}
	if !ok {
		log.Println("delete policy fail")
		err = errors.New("deletePolicyFail")
		return err
	}
	return nil
}

func (s *sysService) GetRole(ctx context.Context) ([]domain.Role, error) {
	rl, err := s.repo.GetRole(ctx)
	return rl, err
}

func (s *sysService) GetAPI(ctx context.Context) ([]domain.API, error) {
	rl, err := s.repo.GetAPI(ctx)
	return rl, err
}

func (s *sysService) GetMenu(ctx context.Context) (menus []domain.Menu, err error) {
	menus, err = s.repo.GetMenu(ctx)
	return menus, err
}

func (s *sysService) GetMenuByUserID(ctx context.Context, userid string) (menus []domain.Menu, err error) {
	menus, err = s.repo.GetMenuByUserID(ctx, userid)
	if err != nil {
		return nil, err
	}
	//处理用户返回菜单
	return menus, err
}
