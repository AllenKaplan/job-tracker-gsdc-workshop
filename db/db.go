package db

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/AllenKaplan/jobTracker/models"
)

type Database struct {
	db *gorm.DB
}

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

func NewDB(dbCfg Config) (Repository, error) {
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbCfg.Host, dbCfg.User, dbCfg.Password, dbCfg.DBName, dbCfg.Port)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(models.App{})
	return Database{db: db}, nil
}

type Repository interface {
	GetApp(context.Context, int) (models.App, error)
	GetApps(context.Context) ([]models.App, error)
	CreateApp(context.Context, models.App) (models.App, error)
	UpdateApp(context.Context, models.App) (models.App, error)
	DeleteApp(context.Context, int) (bool, error)
}

func (db Database) GetApp(ctx context.Context, id int) (models.App, error) {
	var app models.App
	result := db.db.First(&app, "id = ?", id)
	if result.Error != nil {
		return models.App{}, result.Error
	}
	return app, nil
}

func (db Database) GetApps(ctx context.Context) ([]models.App, error) {
	var apps []models.App
	result := db.db.Find(&apps)
	if result.Error != nil {
		return nil, result.Error
	}
	return apps, nil
}

func (db Database) CreateApp(ctx context.Context, app models.App) (models.App, error) {
	result := db.db.Clauses(clause.Returning{}).Create(app)
	if result.Error != nil {
		return models.App{}, result.Error
	}
	return app, nil
}

func (db Database) UpdateApp(ctx context.Context, app models.App) (models.App, error) {
	result := db.db.Clauses(clause.Returning{}).Updates(app)
	if result.Error != nil {
		return models.App{}, nil
	}
	return app, nil

}

func (db Database) DeleteApp(ctx context.Context, id int) (bool, error) {
	result := db.db.Delete(&models.App{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
