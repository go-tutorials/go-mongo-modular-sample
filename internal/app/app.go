package app

import (
	"context"
	"github.com/core-go/health"
	"github.com/core-go/log"
	mgo "github.com/core-go/mongo"
	"github.com/core-go/search"
	"github.com/core-go/search/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"

	. "go-service/internal/usecase/user"
)

type ApplicationContext struct {
	Health *health.Handler
	User   UserHandler
}

func NewApp(ctx context.Context, conf Config) (*ApplicationContext, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Mongo.Uri))
	db := client.Database(conf.Mongo.Database)
	if err != nil {
		return nil, err
	}

	logError := log.ErrorMsg

	userType := reflect.TypeOf(User{})
	userQuery := query.UseQuery(userType)
	userSearchBuilder := mgo.NewSearchBuilder(db, "users", userQuery, search.GetSort)
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userSearchBuilder.Search, userService, logError)

	mongoChecker := mgo.NewHealthChecker(db)
	healthHandler := health.NewHandler(mongoChecker)

	return &ApplicationContext{
		Health: healthHandler,
		User:   userHandler,
	}, nil
}
