package migrator

import (
	"context"
	"github.com/uptrace/bun/migrate"

	ctr "github.com/kylerqws/chatbot/pkg/db/contract"
	ctrcl "github.com/kylerqws/chatbot/pkg/db/contract/client"
)

func Migrate(ctx context.Context, db ctr.DB) (err error) {
	cl := db.Client()
	if err = cl.Connect(); err != nil {
		return err
	}

	defer func(cl ctrcl.Client) {
		if clerr := cl.Close(); clerr != nil {
			if err == nil {
				err = clerr
			}
		}
	}(cl)

	err = db.Migrator().Migrate(ctx, migrations())
	if err != nil {
		return err
	}

	return err
}

func Rollback(ctx context.Context, db ctr.DB) (err error) {
	cl := db.Client()
	if err = cl.Connect(); err != nil {
		return err
	}

	defer func(cl ctrcl.Client) {
		if clerr := cl.Close(); clerr != nil {
			if err == nil {
				err = clerr
			}
		}
	}(cl)

	err = db.Migrator().Rollback(ctx, migrations())
	if err != nil {
		return err
	}

	return err
}

func migrations() *migrate.Migrations {
	migs := migrate.NewMigrations()
	for i := range migrationRegistry {
		migrationRegistry[i](migs)
	}

	return migs
}
