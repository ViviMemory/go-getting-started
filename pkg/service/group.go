package service

import (
	"github.com/heroku/go-getting-started/model"
	"github.com/heroku/go-getting-started/pkg/repository"
)

type GroupService struct {
	repo repository.Group
}

func NewGroupService(repo repository.Group) *GroupService {
	return &GroupService{repo: repo}
}

func (s *GroupService) DetailGroup(groupId int) ([]model.UserGroup, error) {
	return s.repo.DetailGroup(groupId)
}

func (s *GroupService) InviteUserInGroup(groupId int, phone string) (int, error) {
	return s.repo.InviteUserInGroup(groupId, phone)
}

func (s *GroupService) ListInviteUserInGroup(userId int) ([]model.GroupList, error) {
	return s.repo.ListInviteUserInGroup(userId)
}

func (s *GroupService) ActiveInviteUserInGroup(userId int, groupId int, isReject bool) error {
	return s.repo.ActiveInviteUserInGroup(userId, groupId, isReject)
}

func (s *GroupService) CreateGroup(title string, userId int) (int, error) {
	return s.repo.CreateGroup(title, userId)
}

func (s *GroupService) GetAllGroupUser(userId int) ([]model.Group, error) {
	return s.repo.GetAllGroupUser(userId)
}

func (s *GroupService) AddUserInGroup(groupAdd model.GroupAddUserInput) (int, error) {
	return s.repo.AddUserInGroup(groupAdd)
}
