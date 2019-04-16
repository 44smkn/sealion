package datastore_test

import (
	"context"
	"database/sql"
	"fmt"
	"sealion/config"
	"sealion/domain/model"
	"sealion/infrastructure/persistence/datastore"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var createTable = `
CREATE TABLE IF NOT EXISTS tasks (
	 id        INTEGER PRIMARY KEY AUTO_INCREMENT
	,category  VARCHAR(255)
    ,name      VARCHAR(255)
	,do_today  boolean
	,deadline  DATE
	,ticket_id VARCHAR(255)
	,archive   boolean
);
`
var insertRecord = `
INSERT INTO tasks (category, name, do_today, deadline, ticket_id, archive)
    VALUES (?, ?, ?, ?, ?, ?)
`

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func TestGetAll(t *testing.T) {
	conn, err := config.GetDbConn()
	check(err)

	defer conn.Close()
	initTable(conn)
	//defer dropTable(conn)
	initRecords(conn)

	ctx := context.Background()
	r := datastore.NewTaskRepository(conn)
	tasks, err := r.GetAll(ctx)
	check(err)

	fmt.Println(tasks)
	expected := newExpectedData()
	for i := 0; i < len(expected); i++ {
		if diff := cmp.Diff(tasks[i], expected[i]); diff != "" {
			t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
		}
	}
}

func initTable(conn *sql.DB) {
	stmt, err := conn.Prepare(createTable)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)
}

func initRecords(conn *sql.DB) {
	stmt, err := conn.Prepare(insertRecord)
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec("TICKETS", "Prometheusの設定変更", true, "2019-03-04", "MICROUD-1242", false)
	check(err)

	_, err = stmt.Exec("CHORE", "健康診断の申し込み", false, "2019-03-07", "NULL", false)
	check(err)
}

func dropTable(conn *sql.DB) {
	stmt, err := conn.Prepare("DROP TABLE tasks")
	check(err)
	defer stmt.Close()

	_, err = stmt.Exec()
	check(err)
}

func newExpectedData() []model.Task {
	return []model.Task{
		{
			Id:       1,
			Category: "TICKETS",
			Name:     "Prometheusの設定変更",
			DoToday:  true,
			Deadline: time.Date(2019, 3, 4, 0, 0, 0, 0, time.UTC),
		},
		{
			Id:       2,
			Category: "CHORE",
			Name:     "健康診断の申し込み",
			DoToday:  false,
			Deadline: time.Date(2019, 3, 7, 0, 0, 0, 0, time.UTC),
		},
	}
}
