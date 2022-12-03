package service

import (
	"context"
	"go-pagination/entity"
	"go-pagination/repository"
)

type SiteService interface {
	GetTotalRows(ctx context.Context) (int64, error)
	GetAllSites(ctx context.Context) ([]entity.Site, error)
	GetAllSitesPaginatedOffset(ctx context.Context, pagination entity.Pagination) (entity.Pagination, error)
	GetAllSitesPaginatedSeek(ctx context.Context, pagination entity.Pagination) (entity.Pagination, error)
}

type siteService struct {
	siteRepository repository.SiteRepository
}

func NewSiteService(sr repository.SiteRepository) SiteService {
	return &siteService{
		siteRepository: sr,
	}
}

func (s *siteService) GetTotalRows(ctx context.Context) (int64, error) {
	return s.siteRepository.GetTotalRows(ctx)
}

func (s *siteService) GetAllSites(ctx context.Context) ([]entity.Site, error) {
	return s.siteRepository.GetAllSites(ctx)
}

func (s *siteService) GetAllSitesPaginatedOffset(ctx context.Context, pagination entity.Pagination) (entity.Pagination, error) {
	return s.siteRepository.GetAllSitesPaginatedOffset(ctx, pagination)
}

func (s *siteService) GetAllSitesPaginatedSeek(ctx context.Context, pagination entity.Pagination) (entity.Pagination, error) {
	return s.siteRepository.GetAllSitesPaginatedSeek(ctx, pagination)
}
