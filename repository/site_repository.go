package repository

import (
	"context"
	"go-pagination/entity"
	"go-pagination/utils"

	"gorm.io/gorm"
)

type SiteRepository interface {
	GetTotalRows(ctx context.Context) (int64, error)
	GetAllSites(ctx context.Context) ([]entity.Site, error)
	GetAllSitesPaginatedOffset(ctx context.Context, pagination entity.Pagination) (entity.Pagination, error)
	GetAllSitesPaginatedSeek(ctx context.Context, pagination entity.Pagination) (entity.Pagination, error)
}

type siteConnection struct {
	connection *gorm.DB
}

func NewSiteRepository(db *gorm.DB) SiteRepository {
	return &siteConnection{
		connection: db,
	}
}

func (db *siteConnection) GetTotalRows(ctx context.Context) (int64, error) {
	var totalRows int64
	tx := db.connection.Model(&entity.Site{}).Count(&totalRows)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return totalRows, nil
}

func (db *siteConnection) GetAllSites(ctx context.Context) ([]entity.Site, error) {
	var siteList []entity.Site
	tx := db.connection.Find(&siteList)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return siteList, nil
}

func (db *siteConnection) GetAllSitesPaginatedOffset(ctx context.Context, pagination entity.Pagination) (entity.Pagination, error) {
	var siteList []*entity.Site

	totalRows, _ := db.GetTotalRows(ctx)

	db.connection.Debug().Scopes(utils.PaginateOffset(totalRows, &pagination, db.connection)).Find(&siteList)
	pagination.Rows = siteList

	return pagination, nil
}

func (db *siteConnection) GetAllSitesPaginatedSeek(ctx context.Context, pagination entity.Pagination) (entity.Pagination, error) {
	var siteList []*entity.Site

	totalRows, _ := db.GetTotalRows(ctx)

	db.connection.Debug().Scopes(utils.PaginateSeek(totalRows, &pagination, db.connection)).Find(&siteList)
	pagination.Rows = siteList

	return pagination, nil
}
