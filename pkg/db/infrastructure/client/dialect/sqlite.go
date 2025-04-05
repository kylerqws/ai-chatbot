package dialect

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	ctrdlt "github.com/kylerqws/chatbot/pkg/db/contract/client/dialect"
	ctrcfg "github.com/kylerqws/chatbot/pkg/db/contract/config"

	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type sqliteDialect struct {
	config ctrcfg.Config
	db     *bun.DB
}

func NewSQLite(cfg ctrcfg.Config) ctrdlt.Dialect {
	return &sqliteDialect{config: cfg}
}

func (d *sqliteDialect) Connect() error {
	dsn, err := d.prepareDSN(d.config.GetDSN())
	if err != nil {
		return fmt.Errorf("dialect: failed to prepare DSN: %w", err)
	}

	sqldb, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("dialect: failed to open SQLite connection: %w", err)
	}

	d.db = bun.NewDB(sqldb, sqlitedialect.New())

	if err := d.init(); err != nil {
		return fmt.Errorf("dialect: failed to initialize database: %w", err)
	}

	if err := d.db.Ping(); err != nil {
		return fmt.Errorf("dialect: ping to database failed: %w", err)
	}

	return nil
}

func (d *sqliteDialect) Close() error {
	if err := d.db.Close(); err != nil {
		return fmt.Errorf("dialect: failed to close SQLite database: %w", err)
	}

	return nil
}

func (d *sqliteDialect) DB() *bun.DB {
	return d.db
}

func (d *sqliteDialect) prepareDSN(dsn string) (string, error) {
	if strings.TrimSpace(dsn) == "" {
		dsn = "var/database.sqlite"
	}
	if strings.HasPrefix(dsn, ":memory:") {
		return dsn, nil
	}

	dir := filepath.Dir(dsn)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", fmt.Errorf("dialect: failed to create directory for DSN %q: %w", dir, err)
	}

	return dsn, nil
}

func (d *sqliteDialect) init() error {
	if d.config.IsDebug() {
		d.db.AddQueryHook(bundebug.NewQueryHook())
	}

	if _, err := d.db.Exec(`PRAGMA foreign_keys = ON;`); err != nil {
		return fmt.Errorf("dialect: failed to enable foreign keys: %w", err)
	}

	return nil
}
