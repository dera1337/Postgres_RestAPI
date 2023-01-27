package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Class struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// PhoneNumber string
}

func InsertRow(conn *pgxpool.Pool, class Class) error {
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

func UpdateRow(conn *pgxpool.Pool, class Class) error {
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

func DeleteRow(conn *pgxpool.Pool, class Class) error {
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

func ReadRow(conn *pgxpool.Pool, classID int) (Class, error) {
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

func ReadRows(conn *pgxpool.Pool) ([]Class, error) {
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

// func sample() {
// 	class, err := readRow(Conn, 1)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}

// 	fmt.Println(class)

// 	classes, err := readRows(Conn)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println(classes)

// 	class := Class{Name: "James"}
// 	err = insertRow(Conn, class)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println()

// 	class := Class{
// 		Name: "Juan",
// 		ID:   1,
// 	}
// 	err = updateRow(Conn, class)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println()

// 	class := Class{Name: "James"}
// 	err = insertRow(Conn, class)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println()

// 	class := Class{ID: 1}
// 	err = deleteRow(Conn, class)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	fmt.Println()
// }
