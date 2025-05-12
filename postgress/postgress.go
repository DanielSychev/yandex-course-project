package postgress

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Добавьте этот импорт
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

type Config struct {
	Host     string `yaml:"POSTGRES_HOST" env:"POSTGRES_HOST" envDefault:"postgres"` //localhost
	Port     uint16 `yaml:"POSTGRES_PORT" env:"POSTGRES_PORT" envDefault:"5432"`
	Username string `yaml:"POSTGRES_USER" env:"POSTGRES_USER" envDefault:"root"`
	Password string `yaml:"POSTGRES_PASSWORD" env:"POSTGRES_PASSWORD" envDefault:"1234"`
	Database string `yaml:"POSTGRES_DB" env:"POSTGRES_DB" envDefault:"postgres"`

	MinCon int32 `yaml:"POSTGRES_MIN_CONN" env:"POSTGRES_MIN_CONN" envDefault:"5"`
	MaxCon int32 `yaml:"POSTGRES_MAX_CONN" env:"POSTGRES_MAX_CONN" envDefault:"10"`
}

func New(ctx context.Context, config Config) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", //&pool_max_conns=%d&pool_min_conns=%d
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
		//config.MaxCon,
		//config.MinCon,
	)

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	//conn, err := pgx.Connect(ctx, "postgres://root:1234@localhost:5432/postgres?sslmode=disable")
	//if err != nil {
	//	log.Fatal("Connection failed: ", err)
	//}
	//fmt.Println("Connected to PostgreSQL!")

	m, err := migrate.New(
		"file://db/migrations", // Путь к файлам миграций (SQL-скриптам)
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			config.Username,
			config.Password,
			config.Host,
			config.Port,
			config.Database))
	if err != nil {
		return nil, fmt.Errorf("unable to create migrations: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, fmt.Errorf("unable to run migrations: %v", err)
	}
	fmt.Println("Created Migrations!")
	return conn, nil
}
