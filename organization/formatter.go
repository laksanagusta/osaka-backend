package organization

import (
	"time"
)

type OrganizationFormatter struct {
	ID        int                       `json:"id"`
	Name      string                    `json:"name"`
	Status    int8                      `json:"status"`
	User      OrganizationUserFormatter `json:"user"`
	CreatedAt time.Time                 `json:"created_at"`
	UpdatedAt time.Time                 `json:"updated_at"`
}

type OrganizationUserFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func FormatOrganization(organization Organization) OrganizationFormatter {
	organizationFormatter := OrganizationFormatter{}
	organizationFormatter.ID = organization.ID
	organizationFormatter.Name = organization.Name
	organizationFormatter.Status = organization.Status
	organizationFormatter.CreatedAt = organization.CreatedAt
	organizationFormatter.UpdatedAt = organization.UpdatedAt

	userFormatter := OrganizationUserFormatter{}
	userFormatter.ID = organization.User.ID
	userFormatter.Name = organization.User.Name
	userFormatter.Username = organization.User.Username
	userFormatter.CreatedAt = organization.User.CreatedAt
	userFormatter.UpdatedAt = organization.User.UpdatedAt

	organizationFormatter.User = userFormatter

	return organizationFormatter

}
