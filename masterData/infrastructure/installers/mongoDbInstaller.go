package installers

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"natsMicros/buildingBlocks/application/configurations"
	"time"
)

func NewMongoDbClient(lc fx.Lifecycle, setting *configurations.MongoDbSetting) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(setting.ConnectionString))
	if err != nil {
		panic("Error connecting to MongoDB: " + err.Error())
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			cancel()
			return nil
		},
	})
	return client
}
