package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/sarulabs/di"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    app := createApp()
    defer app.Delete()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Create a request and delete it once it has been handled.
        // Deleting the request will close the connection.
        request, _ := app.SubContainer()
        defer request.Delete()

        handler(w, r, request)
    })

    http.ListenAndServe(":8080", nil)
}

func createApp() di.Container {
    builder, _ := di.NewBuilder()

    builder.Add([]di.Def{
        {
            // Define the connection pool in the App scope.
            // There will be one for the whole application.
            Name:  "mysql-pool",
            Scope: di.App,
            Build: func(ctn di.Container) (interface{}, error) {
                db, err := sql.Open("mysql", "user:password@/")
                db.SetMaxOpenConns(1)
                return db, err
            },
            Close: func(obj interface{}) error {
                return obj.(*sql.DB).Close()
            },
        },
        {
            // Define the connection in the Request scope.
            // Each request will use its own connection.
            Name:  "mysql",
            Scope: di.Request,
            Build: func(ctn di.Container) (interface{}, error) {
                pool := ctn.Get("mysql-pool").(*sql.DB)
                return pool.Conn(context.Background())
            },
            Close: func(obj interface{}) error {
                return obj.(*sql.Conn).Close()
            },
        },
    }...)

    // Returns the app Container.
    return builder.Build()
}

func handler(w http.ResponseWriter, r *http.Request, ctn di.Container) {
    // Retrieve the connection.
    conn := ctn.Get("mysql").(*sql.Conn)

    var variable, value string

    row := conn.QueryRowContext(context.Background(), "SHOW STATUS WHERE `variable_name` = 'Threads_connected'")
    row.Scan(&variable, &value)

    // Display how many connections are opened.
    // As the connection is closed when the request is deleted,
    // the value should not be be higher than the number set with db.SetMaxOpenConns(1).
    w.Write([]byte(variable + ": " + value))
}