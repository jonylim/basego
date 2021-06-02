package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/helper"
	"github.com/jonylim/basego/internal/pkg/common/logger"

	// See https://godoc.org/github.com/lib/pq.
	_ "github.com/lib/pq"
)

type config struct {
	DriverName         string
	Host               string
	Port               string
	SSLMode            string
	ConnectTimeout     string
	User               string
	Password           string
	DbName             string
	ConnMaxIdlePercent string
	ConnMaxLifetime    string
	CheckTable         string
}

var db *sql.DB

// Init generates and tests connection string using the given configurations.
func Init() {
	// Load environment variables.
	cfg := getConfig(envvar.Database)

	// Init database connection.
	var err error
	db, err = initDB(cfg)
	if err != nil {
		// Should not fail! This is the main database for this app.
		logger.Println("db", fmt.Sprintf("ERROR: initDB: %v", err))
		os.Exit(1)
	}
}

// CloseAll closes all open database connections.
func CloseAll() {
	logger.Println("db", "Closing database...")
	if db != nil {
		db.Close()
	}
}

// Get returns database pool.
func Get() *sql.DB {
	return db
}

func getConfig(dbEnv envvar.DatabaseEnvVars) config {
	return config{
		DriverName:         os.Getenv(dbEnv.DriverName),
		Host:               os.Getenv(dbEnv.Host),
		Port:               os.Getenv(dbEnv.Port),
		SSLMode:            os.Getenv(dbEnv.SSLMode),
		ConnectTimeout:     os.Getenv(dbEnv.ConnectTimeout),
		User:               os.Getenv(dbEnv.User),
		Password:           os.Getenv(dbEnv.Password),
		DbName:             os.Getenv(dbEnv.DbName),
		ConnMaxIdlePercent: os.Getenv(dbEnv.ConnMaxIdlePercent),
		ConnMaxLifetime:    os.Getenv(dbEnv.ConnMaxLifetime),
		CheckTable:         os.Getenv(dbEnv.CheckTable),
	}
}

func initDB(cfg config) (*sql.DB, error) {
	logger.Println("db", "Initializing database...")

	// Validate the config.
	if cfg.DriverName == "" {
		return nil, errors.New("initDB: Driver is undefined")
	} else if cfg.Port == "" {
		return nil, errors.New("initDB: Port is undefined")
	} else if _, err := helper.StringToInt(cfg.Port); err != nil {
		return nil, fmt.Errorf("initDB: %v", err)
	} else if cfg.User == "" {
		return nil, errors.New("initDB: User is undefined")
	} else if cfg.DbName == "" {
		return nil, errors.New("initDB: Database name is undefined")
	}
	if cfg.ConnectTimeout != "" {
		if _, err := helper.StringToInt(cfg.ConnectTimeout); err != nil {
			return nil, fmt.Errorf("initDB: %v", err)
		}
	}
	var connMaxIdlePercent int
	var connMaxLifetime int64
	if cfg.ConnMaxIdlePercent != "" {
		i, err := helper.StringToInt(cfg.ConnMaxIdlePercent)
		if err != nil {
			return nil, fmt.Errorf("initDB: %v", err)
		}
		connMaxIdlePercent = i
	}
	if cfg.ConnMaxLifetime != "" {
		i, err := helper.StringToInt64(cfg.ConnMaxLifetime)
		if err != nil {
			return nil, fmt.Errorf("initDB: %v", err)
		}
		connMaxLifetime = i
	}

	// Generate the connection string.
	connStr := "host=" + cfg.Host +
		" port=" + cfg.Port +
		" sslmode=" + cfg.SSLMode +
		" user=" + cfg.User +
		" password='" + cfg.Password + "' " +
		" dbname=" + cfg.DbName
	if cfg.ConnectTimeout != "" {
		connStr = connStr + " connect_timeout=" + cfg.ConnectTimeout
	}

	// Open the database.
	db, err := sql.Open(cfg.DriverName, connStr)
	if err != nil {
		return nil, fmt.Errorf("initDB: %v", logger.FromError(err))
	}

	// Set the maximum open & idle connections.
	maxConns, err := getMaxConnections(db)
	if err != nil {
		return nil, fmt.Errorf("initDB: %v", err)
	} else if maxConns >= 12 {
		maxConns -= 8
	}
	db.SetMaxOpenConns(maxConns)
	logger.Println("db", fmt.Sprintf("Maximum open connections set to %v", maxConns))
	if cfg.ConnMaxIdlePercent != "" {
		maxIdles := maxConns * connMaxIdlePercent / 100
		if maxIdles < 2 {
			maxIdles = 2
		}
		db.SetMaxIdleConns(maxIdles)
		logger.Println("db", fmt.Sprintf("Maximum idle connections set to %v", maxIdles))
	}
	if cfg.ConnMaxLifetime != "" {
		d := time.Duration(connMaxLifetime) * time.Second
		db.SetConnMaxLifetime(d)
		logger.Println("db", fmt.Sprintf("Connection maximum lifetime set to %v", d))
	}

	// Try the database.
	if err := tryDB(db, cfg.CheckTable); err != nil {
		return nil, fmt.Errorf("initDB: %v", err)
	}

	// Done.
	logger.Println("db", "Connected to database")
	return db, nil
}

func getMaxConnections(db *sql.DB) (maxConns int, err error) {
	err = db.QueryRow("SHOW max_connections").Scan(&maxConns)
	if err != nil {
		err = fmt.Errorf("getMaxConnections: %v", logger.FromError(err))
	}
	return
}

func tryDB(db *sql.DB, checkTable string) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("tryDB: %v", logger.FromError(err))
	}
	var checkRes int
	err = db.QueryRow("SELECT MIN(id) FROM " + checkTable).Scan(&checkRes)
	if err != nil {
		return fmt.Errorf("tryDB: %v", logger.FromError(err))
	}
	logger.Println("db", fmt.Sprintf("tryDB: Check result from %s: %v", checkTable, checkRes))
	return nil
}

func sqlTimestampToUnixMilliseconds(column string) string {
	return "COALESCE(FLOOR(EXTRACT(EPOCH FROM " + column + ") * 1000), 0)::BIGINT"
}

func sqlTimestampToUnixNanoseconds(column string) string {
	return "COALESCE(FLOOR(EXTRACT(EPOCH FROM " + column + ") * 1e9), 0)::BIGINT"
}
