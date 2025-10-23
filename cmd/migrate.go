package cmd

import (
	"fmt"
	"goframe/internal/core/bootstrap"
	"goframe/internal/core/config"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

const migrationPath = "database/migrations"

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		bootstrap.InitDB(cfg)

		db, err := bootstrap.DB.DB()
		if err != nil {
			fmt.Println("❌ Failed to get *sql.DB:", err)
			os.Exit(1)
		}

		driver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			fmt.Println("❌ Failed to create driver:", err)
			os.Exit(1)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://"+migrationPath,
			"mysql",
			driver,
		)
		if err != nil {
			fmt.Println("❌ Migration init failed:", err)
			os.Exit(1)
		}

		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			fmt.Println("❌ Migration failed:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Migration completed")
	},
}

var rollbackCmd = &cobra.Command{
	Use:   "migrate:rollback",
	Short: "Rollback last migration",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig()
		bootstrap.InitDB(cfg)

		db, err := bootstrap.DB.DB()
		if err != nil {
			fmt.Println("❌ Failed to get *sql.DB:", err)
			os.Exit(1)
		}

		driver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			fmt.Println("❌ Failed to create driver:", err)
			os.Exit(1)
		}

		m, err := migrate.NewWithDatabaseInstance(
			"file://"+migrationPath,
			"mysql",
			driver,
		)
		if err != nil {
			fmt.Println("❌ Migration init failed:", err)
			os.Exit(1)
		}

		if err := m.Steps(-1); err != nil {
			fmt.Println("❌ Rollback failed:", err)
			os.Exit(1)
		}

		fmt.Println("✅ Rollback completed")
	},
}

var createMigrationCmd = &cobra.Command{
	Use:   "create:migration",
	Short: "Create a new migration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		timestamp := time.Now().Format("20060102150405")

		// Ensure migrations folder exists
		if _, err := os.Stat(migrationPath); os.IsNotExist(err) {
			if err := os.MkdirAll(migrationPath, os.ModePerm); err != nil {
				fmt.Println("❌ Failed to create migrations folder:", err)
				return
			}
		}

		upFile := filepath.Join(migrationPath, fmt.Sprintf("%s_%s.up.sql", timestamp, name))
		downFile := filepath.Join(migrationPath, fmt.Sprintf("%s_%s.down.sql", timestamp, name))

		upTemplate := "-- +migrate Up\n\n"
		downTemplate := "-- +migrate Down\n\n"

		if err := os.WriteFile(upFile, []byte(upTemplate), 0644); err != nil {
			fmt.Println("❌ Failed to create UP migration:", err)
			return
		}
		if err := os.WriteFile(downFile, []byte(downTemplate), 0644); err != nil {
			fmt.Println("❌ Failed to create DOWN migration:", err)
			return
		}

		fmt.Printf("✅ Migration files created:\n  %s\n  %s\n", upFile, downFile)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(rollbackCmd)
	rootCmd.AddCommand(createMigrationCmd)
}
