package organization

import (
	"errors"
)

type Service interface {
	Save(input OrganizationCreateInput) (Organization, error)
	FindById(organizationId int) (Organization, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Save(input OrganizationCreateInput) (Organization, error) {
	organization := Organization{}
	organization.Name = input.Name
	organization.Status = 1

	newOrganization, err := s.repository.Save(organization)
	if err != nil {
		return newOrganization, err
	}

	return newOrganization, nil
}

func (s *service) FindById(organizationId int) (Organization, error) {
	organization, err := s.repository.FindById(organizationId)
	if err != nil {
		return organization, err
	}

	if organization.ID == 0 {
		return organization, errors.New("Organization not found")
	}

	return organization, nil
}
