package main

import (
	"database/sql"
	"fmt"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/golang-migrate/migrate/v4"
	mpostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	withEmbeddedDB(
		func(db *gorm.DB, database string) {
			setup(db, database)
			generate(db)
		},
	)
}

func setup(db *gorm.DB, database string) {
	must(
		assert(
			migrate.NewWithDatabaseInstance(
				"file://resources/migrations",
				database,
				assert(mpostgres.WithInstance(
					assert(db.DB()),
					&mpostgres.Config{DatabaseName: database},
				)),
			),
		).Up(),
	)
	must(db.Migrator().DropTable("schema_migrations"))
}

func generate(db *gorm.DB) {
	generator := gen.NewGenerator(gen.Config{
		OutPath:      "generated/dao",
		ModelPkgPath: "generated/dao/model",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	generator.UseDB(db)
	generator.ApplyBasic(generator.GenerateModel("city"))
	generator.ApplyBasic(generator.GenerateModel("forecast"))
	generator.Execute()
}

func connect(driver string, url string) *gorm.DB {
	return assert(gorm.Open(
		postgres.New(
			postgres.Config{
				DriverName: driver,
				Conn:       assert(sql.Open(driver, fmt.Sprintf("%s?sslmode=disable", url))),
			},
		),
		&gorm.Config{},
	))
}

func withEmbeddedDB(callback func(*gorm.DB, string)) {
	config := embeddedpostgres.DefaultConfig()
	db := embeddedpostgres.NewDatabase(config)
	defer func(db *embeddedpostgres.EmbeddedPostgres) {
		must(db.Stop())
	}(assert(db, db.Start()))

	callback(connect("postgres", config.GetConnectionURL()), "postgres")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func assert[R any](result R, err error) R {
	if err != nil {
		panic(err)
	}

	return result
}
