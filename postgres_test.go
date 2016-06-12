package tdgw

import (
	"database/sql"
	"io/ioutil"
	"testing"

	_ "github.com/lib/pq"
)

// before running test, create user and database
// CREATE USER tdgw_test;
// CREATE DATABASE  tdgw_test OWNER tdgw_test;

func testPgSetup(t *testing.T) (*sql.DB, func()) {
	conn, err := sql.Open("postgres", "user=tdgw_test dbname=tdgw_test sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	setupSQL, err := ioutil.ReadFile("./sql_test/postgres_setup.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = conn.Exec(string(setupSQL))
	if err != nil {
		t.Fatal(err)
	}
	cleanup := func() {
		conn.Close()
	}
	return conn, cleanup
}

func TestPgLoadColumnDef(t *testing.T) {
	conn, cleanup := testPgSetup(t)
	defer cleanup()

	schema := "public"
	table := "t1"
	cols, err := PgLoadColumnDef(conn, schema, table)
	if err != nil {
		t.Fatal(err)
	}
	for _, c := range cols {
		t.Logf("%+v", c)
	}
}

func TestPgLoadTableDef(t *testing.T) {
	conn, cleanup := testPgSetup(t)
	defer cleanup()

	schema := "public"
	tbls, err := PgLoadTableDef(conn, schema)
	if err != nil {
		t.Fatal(err)
	}
	for _, tbl := range tbls {
		t.Logf("%+v", tbl)
	}
}
