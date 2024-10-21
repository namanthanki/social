package main

import (
	"github.com/joho/godotenv"
	"github.com/namanthanki/social/internal/db"
	"github.com/namanthanki/social/internal/env"
	"github.com/namanthanki/social/internal/store"
	"go.uber.org/zap"
)

const version = "0.0.1"

//	@title			social
//	@description	social media API for golang practice.

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	v1

//	@securityDefinitions.apiKey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	err := godotenv.Load()
	if err != nil {
		// log.Fatalf("Error loading .env file: %e", err)
		logger.Panicf("Error loading .env file: %e", err)
	}

	cfg := config{
		address: env.GetString("API_ADDRESS", ":1337"),
		apiURL:  env.GetString("EXTERNAL_URL", "localhost:1337"),
		db: dbConfig{
			address:      env.GetString("DB_ADDRESS", "postgresql://postgres:password@localhost:5432/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 5),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 2),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "dev"),
	}

	db, err := db.New(
		cfg.db.address,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	if err != nil {
		// log.Fatalf("Error connecting to the database: %e", err)
		logger.Panic(err)
	}

	defer db.Close()
	// log.Printf("Database Pool established")
	logger.Info("Database pool established")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()

	// log.Fatal(app.run(mux))
	logger.Fatal(app.run(mux))
}
