package service

import "github.com/google/uuid"

type IContentService interface {
    RandomContent() string
}

type ContentService struct {
}

func (cs ContentService) RandomContent() string {
    return uuid.NewString()
}
