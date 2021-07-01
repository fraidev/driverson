package main

import (
	"context"
	"fmt"

	"google.golang.org/api/drive/v3"
)

func main() {
	srv, err := drive.NewService(context.Background())

	fmt.Print(err)

	list, err := srv.Files.List().PageSize(10).
		Fields("nextPageToken, files(id, name)").Do()

	fmt.Print(err)

	for _, i := range list.Files {
		fmt.Printf("%s (%s)\n", i.Name, i.Id)
	}
}
