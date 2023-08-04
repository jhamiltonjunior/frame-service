package main

import (
	"fmt"

	"github.com/jhamiltonjunior/saas/src/domain/frame/validation"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	// user := os.Getenv("USER_MAIL")
	// fmt.Println(user)

	// Exec()
	title := validation.Title{}

	if _, err := title.Create("jose"); err != nil {
		fmt.Println(err)
	}


}
