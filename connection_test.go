package dm_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/gaoyuan98/dm"
)

const localTestDSN = "dm://SYSDBA:SYSDBA123#@120.53.45.167:5236?socketTimeout=5"

func TestDMConnectionLocalhost(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := sql.Open("dm", localTestDSN)
	if err != nil {
		t.Fatalf("open dm connection: %v", err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if err := db.PingContext(ctx); err != nil {
		t.Fatalf("ping dm database %q: %v", localTestDSN, err)
	}

	var got int
	if err := db.QueryRowContext(ctx, "select 1").Scan(&got); err != nil {
		t.Fatalf("query dm database: %v", err)
	}
	if got != 1 {
		t.Fatalf("query result = %d, want 1", got)
	}
}
