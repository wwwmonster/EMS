package userService

import (
	"context"
	"ems/mt/golang/sqlc/emsdb"
	"log"

	"github.com/jackc/pgx/v5"
)

func GetUserByEmail(email string) (emsdb.User, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "postgres://admin:123456@localhost:5432/Angular18")
	if err != nil {
		return emsdb.User{}, err
	}
	defer conn.Close(ctx)

	queries := emsdb.New(conn)

	// list all authors
	user, err := queries.GetUserByEmail(ctx, email)
	if err != nil {
		return emsdb.User{}, err
	}
	log.Println("========user:", user)

	// // create an author
	// insertedAuthor, err := queries.CreateAuthor(ctx, tutorial.CreateAuthorParams{
	// 	Name: "Brian Kernighan",
	// 	Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	// })
	// if err != nil {
	// 	return err
	// }
	// log.Println(insertedAuthor)

	// // get the author we just inserted
	// fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	// if err != nil {
	// 	return err
	// }

	// prints true
	// log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
	return user, nil
}
