package db

import (
	"context"
	"errors"

	"github.com/AllenKaplan/jobTracker/models"
)

type FakeDatabase struct {
	db map[int]models.App
}

func newFakeDB() Repository {
	return FakeDatabase{db: make(map[int]models.App)}
}

func (db FakeDatabase) GetApp(ctx context.Context, id int) (models.App, error) {
	app, ok := db.db[id]
	if !ok {
		return models.App{}, errors.New("App does not exist")
	}
	return app, nil
}

func (db FakeDatabase) GetApps(ctx context.Context) ([]models.App, error) {
	var apps []models.App
	for _, app := range db.db {
		apps = append(apps, app)
	}
	return apps, nil
}

func (db FakeDatabase) CreateApp(ctx context.Context, app models.App) (models.App, error) {
	db.db[len(db.db)+1] = app
	return app, nil
}

func (db FakeDatabase) UpdateApp(ctx context.Context, app models.App) (models.App, error) {
	db.db[app.Id] = app
	return app, nil
}

func (db FakeDatabase) DeleteApp(ctx context.Context, id int) (bool, error) {
	delete(db.db, id)
	return true, nil
}
