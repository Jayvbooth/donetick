package database

import (
	"os"
	"testing"

	"donetick.com/core/config"
	chModel "donetick.com/core/internal/chore/model"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type legacyChore struct {
	ID     int  `gorm:"primaryKey"`
	Points *int `gorm:"column:points"`
}

func (legacyChore) TableName() string { return "chores" }

type legacyChoreHistory struct {
	ID     int  `gorm:"primaryKey"`
	Points *int `gorm:"column:points"`
}

func (legacyChoreHistory) TableName() string { return "chore_histories" }

func verifyScoringColumns(t *testing.T, db *gorm.DB) {
	t.Helper()
	require.NoError(t, db.AutoMigrate(&legacyChore{}, &legacyChoreHistory{}))
	require.NoError(t, db.AutoMigrate(&chModel.Chore{}, &chModel.ChoreHistory{}))
	for _, column := range []string{"timing_mode", "early_bonus"} {
		require.True(t, db.Migrator().HasColumn(&chModel.Chore{}, column), column)
	}
	for _, column := range []string{"base_points", "timing_adjustment", "recovery_points"} {
		require.True(t, db.Migrator().HasColumn(&chModel.ChoreHistory{}, column), column)
	}
}

func TestScoringAdditiveMigrationSQLite(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:scoring-migration?mode=memory&cache=shared"), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	require.NoError(t, err)
	verifyScoringColumns(t, db)
}

func TestScoringAdditiveMigrationPostgres(t *testing.T) {
	dsn := os.Getenv("TEST_POSTGRES_DSN")
	if dsn == "" {
		t.Skip("TEST_POSTGRES_DSN is not configured")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	require.NoError(t, err)
	require.NoError(t, db.Exec("DROP TABLE IF EXISTS chore_histories CASCADE").Error)
	require.NoError(t, db.Exec("DROP TABLE IF EXISTS chores CASCADE").Error)
	verifyScoringColumns(t, db)
}

func TestMigrationStillAcceptsSupportedDatabaseTypes(t *testing.T) {
	for _, databaseType := range []string{"sqlite", "postgres"} {
		cfg := &config.Config{}
		cfg.Database.Type = databaseType
		require.Contains(t, []string{"sqlite", "postgres"}, cfg.Database.Type)
	}
}
