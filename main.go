package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Class struct {
	ID   int
	Name string
	// PhoneNumber string
}

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	bgCtx := context.Background()
	conn, err := pgx.Connect(bgCtx, "postgres://postgres:Paramore_12345@localhost:5432/postgres")
	if err != nil {
		log.Fatal("cannot connect to database")
		return
	}
	defer conn.Close(bgCtx)

	err = conn.Ping(bgCtx)
	if err != nil {
		log.Fatal("belom konek coy")
		return
	}

	// err = initTables(conn)
	// if err != nil {
	// 	log.Fatal("failed to init tables")
	// 	return
	// }

	// class, err := readRow(conn, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// fmt.Println(class)

	// classes, err := readRows(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println(classes)

	// class := Class{Name: "James"}
	// err = insertRow(conn, class)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println()

	// class := Class{
	// 	Name: "Juan",
	// 	ID:   1,
	// }
	// err = updateRow(conn, class)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println()

	// class := Class{Name: "James"}
	// err = insertRow(conn, class)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println()

	// class := Class{ID: 1}
	// err = deleteRow(conn, class)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// fmt.Println()
}

func initTables(conn *pgx.Conn) error {
	sqlFilePath := `C:\Users\yoris\Desktop\go-wokspace\src\exercisee\init.sql`

	c, err := os.ReadFile(sqlFilePath)
	if err != nil {
		return err
	}
	sql := string(c)

	_, err = conn.Exec(context.Background(), sql)
	if err != nil {
		return err
	}

	return nil
}

func insertRow(conn *pgx.Conn, class Class) error {
	ctx := context.Background()

	queryString :=
		`
	INSERT INTO class (name)
	VALUES ($1) 
	`

	_, err := conn.Exec(ctx, queryString, class.Name)
	if err != nil {
		return err
	}
	return nil
}

func updateRow(conn *pgx.Conn, class Class) error {
	ctx := context.Background()

	queryString :=
		`
	UPDATE class 
	SET name = $1
	WHERE class_id = $2
	`

	_, err := conn.Exec(ctx, queryString, class.Name, class.ID)
	if err != nil {
		return err
	}
	return nil
}

func deleteRow(conn *pgx.Conn, class Class) error {
	ctx := context.Background()

	queryString :=
		`
	DELETE from class
	WHERE class_id = $1
	`
	_, err := conn.Exec(ctx, queryString, class.ID)
	if err != nil {
		return err
	}
	return nil
}

func readRow(conn *pgx.Conn, classID int) (Class, error) {
	ctx := context.Background()

	queryString :=
		`
	SELECT * FROM class
	WHERE class_id = $1
	`

	var class Class
	row := conn.QueryRow(ctx, queryString, classID)
	err := row.Scan(&class.ID, &class.Name)
	if err != nil {
		return class, err
	}

	return class, nil
}

func readRows(conn *pgx.Conn) ([]Class, error) {
	ctx := context.Background()

	queryString :=
		`
	SELECT * FROM class
	`

	rows, err := conn.Query(ctx, queryString)
	if err != nil {
		return nil, err
	}

	classes := []Class{}
	for rows.Next() {
		var class Class
		err := rows.Scan(&class.ID, &class.Name)
		if err != nil {
			return nil, err
		}

		classes = append(classes, class)
	}

	return classes, nil
}
