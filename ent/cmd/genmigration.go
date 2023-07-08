//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"log"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a local migration directory able to understand golang-migrate migration files for replay.
	d, err := sqltool.NewGolangMigrateDir("../gen")
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed creating atlas migration directory: %v", err))
	}

	// Load the graph.
	graph, err := entc.LoadGraph("../schema", &gen.Config{})
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed load the graph: %v", err))
	}

	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalf(fmt.Sprintf("failed get the tables: %v", err))
	}

	opts := []schema.MigrateOption{
		schema.WithDir(d), // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL),           // Ent dialect to use
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithForeignKeys(false),
	}

	if err := schema.Diff(context.Background(), "AtlasDsn", "changes", tbls, opts...); err != nil {
		log.Fatalf(err.Error())
	}

	log.Fatalf("generate migration done")
}
