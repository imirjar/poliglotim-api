package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	// "github.com/jackc/pgx/v5"
)

type Storage struct {
	psql *pgxpool.Pool
}

func New(ctx context.Context) *Storage {
	return &Storage{}
}

func (s *Storage) Сonnect(ctx context.Context, PsqlConn string) error {

	dbConfig, err := pgxpool.ParseConfig(PsqlConn)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	// log.Print(dbConfig)
	// Connect psql
	pgx, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		panic(err)
	}
	s.psql = pgx

	// err = s.makeMigrations(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	return err
}

func (s *Storage) Disconnect(ctx context.Context) error {
	defer s.psql.Close()
	return nil
}

// func (s *Storage) makeMigrations(ctx context.Context) error {
// 	migration, err := os.ReadFile("./internal/storage/migrations/0001_init.up.sql")
// 	// log.Print(string(migration))

// 	log.Printf("Применяем миграцию:\n%s", string(migration))

// 	// Выполнение SQL-запроса
// 	if _, err := s.psql.Exec(ctx, string(migration)); err != nil {
// 		return fmt.Errorf("ошибка выполнения миграции: %v", err)
// 	}

// 	return err
// }
