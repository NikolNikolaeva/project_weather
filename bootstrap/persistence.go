package bootstrap

import (
	"database/sql"
	"errors"
	"fmt"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/fx"
	gorm_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"project_weather/config"
	"project_weather/generated/dao"
	"strconv"
)

var FXModule_Persistence = fx.Module(
	"persistence",
	fx.Provide(
		createEmbeddedPostgres,
		createEntityManager,
		createDatabaseMigrator,
		createDatabaseDriver,
		createDatabaseConnection,
		createEntityManagerConnection,
	),

	fx.Invoke(
		performDatabaseSchemaMigration,
	),
)

func createEntityManagerConnection(db *sql.DB) (*gorm.DB, error) {
	return gorm.Open(
		gorm_postgres.New(gorm_postgres.Config{Conn: db}),
		&gorm.Config{
			TranslateError:         true,
			SkipDefaultTransaction: false,
		},
	)
}

func createEmbeddedPostgres(configuration *config.ApplicationConfiguration) (*embeddedpostgres.EmbeddedPostgres, error) {
	port, err := strconv.Atoi(configuration.DBPort)
	if err != nil {
		return nil, err
	}

	db := embeddedpostgres.NewDatabase(
		embeddedpostgres.DefaultConfig().
			Database(configuration.DBName).
			Port(uint32(port)).
			Username(configuration.DBUsername).
			Username(configuration.DBPassword),
	)

	return db, db.Start()
}

func createEntityManager(db *gorm.DB) *dao.Query {
	fmt.Printf("--------------------------------\n\n=====================\n")
	q := dao.Use(db)
	city, err := q.City.First()
	if err != nil {
		panic(err)
	}
	fmt.Printf("--------------------------------\n\n=====================%#v\n", city)
	return q
}

func createDatabaseMigrator(config *config.ApplicationConfiguration, driver database.Driver) (*migrate.Migrate, error) {
	return migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%v", "./resources/migrations"), "postgres", driver)
}
func performDatabaseSchemaMigration(migrator *migrate.Migrate) error {

	_, dirty, _ := migrator.Version()
	if dirty {
		_ = migrator.Drop()

		return fmt.Errorf("failed performing migrations, dirty DB state detected")
	}

	err := migrator.Up()

	if errors.Is(err, migrate.ErrNoChange) {
		return nil
	}

	return err
}

func createDatabaseDriver(db *sql.DB) (database.Driver, error) {
	return postgres.WithInstance(db, &postgres.Config{})
}

func createDatabaseConnection(config *config.ApplicationConfiguration) (*sql.DB, error) {
	db, err := sql.Open("postgres", buildDatabaseURL(config))
	if err != nil {
		return nil, err
	}

	return db, nil
}
func buildDatabaseURL(config *config.ApplicationConfiguration) string {

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}
