package internal

import (
	"gulnur.net/services/contact/internal/repository"
	"gulnur.net/services/contact/internal/usecase"
)

type Service struct {
	ContactUC usecase.ContactUseCase
	GroupUC   usecase.GroupUseCase
}

func NewService() *Service {
	contactRepo := repository.NewContactRepository()
	groupRepo := repository.NewGroupRepository()

	contactUC := usecase.NewContactUseCase(contactRepo,  groupRepo)
	groupUC := usecase.NewGroupUseCase(groupRepo)

	return &Service{
		ContactUC: contactUC,
		GroupUC:   groupUC,
	}
}
