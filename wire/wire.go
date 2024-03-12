//go:build wireinject
// +build wireinject

package wire

import (
	"flashcards/database"
	"flashcards/logger"
	"flashcards/resources"
	"flashcards/server"
	"flashcards/services"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		server.NewServer,
		resources.NewGinEngine,
		logger.NewLogger,
	)
	return nil, nil
}

func InitializeDatabase() (*gorm.DB, error) {
	wire.Build(database.NewDatabase)
	return nil, nil
}

func InitializeUserService(db *gorm.DB, logger *logger.Logger) services.UserService {
	wire.Build(services.NewUserService)
	return nil // The return value is just a placeholder
}
