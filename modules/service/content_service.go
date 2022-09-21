package service

import (
	"github.com/google/uuid"
	"github.com/google/wire"
)

type IContentService interface {
	RandomContent() string
}

type ContentService struct {
}

var contentServiceSet = wire.NewSet(
	wire.Bind(new(IContentService), new(ContentService)),
	ContentServiceProvider,
)

func ContentServiceProvider() ContentService {
	return ContentService{}
}

func (cs ContentService) RandomContent() string {
	return uuid.NewString()
}
