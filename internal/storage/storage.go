package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	mongo *mongo.Client
	psql  *pgx.Conn
}

func New(ctx context.Context) *Storage {
	return &Storage{}
}

func (s *Storage) Сonnect(ctx context.Context, PsqlConn, MongoConn string) error {

	// Connect psql
	pgx, err := pgx.Connect(context.Background(), PsqlConn)
	if err != nil {
		panic(err)
	}
	s.psql = pgx

	// err = s.makeMigrations(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// Connect MongoDB
	mongo, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoConn))
	if err != nil {
		panic(err)
	}
	s.mongo = mongo

	return err
}

func (s *Storage) Disconnect(ctx context.Context) error {
	err := s.mongo.Disconnect(context.Background())
	if err != nil {
		panic(err)
	}
	err = s.psql.Close(context.Background())
	if err != nil {
		panic(err)
	}

	return err
}

func (p *Storage) makeMigrations(ctx context.Context) error {
	migration, err := os.ReadFile("./internal/storage/migrations/0001_init.up.sql")
	// log.Print(string(migration))

	log.Printf("Применяем миграцию:\n%s", string(migration))

	// Выполнение SQL-запроса
	if _, err := p.psql.Exec(ctx, string(migration)); err != nil {
		return fmt.Errorf("ошибка выполнения миграции: %v", err)
	}

	return err
}
