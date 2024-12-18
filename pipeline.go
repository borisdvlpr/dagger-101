package main

import (
	"context"
	"fmt"

	"dagger.io/dagger"
)

func main() {
	// create a shared context
	ctx := context.Background()

	// run the stages of the pipeline
	if err := Build(ctx); err != nil {
		fmt.Println("error:", err)
		panic(err)
	}
}

func Build(ctx context.Context) error {
	// initialize Dagger client
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	rust := client.Container().From("rust:alpine").
		WithDirectory("hello-rust", client.Host().Directory("hello-rust")).
		WithExec([]string{"cargo", "--version"})

	rust = rust.
		WithWorkdir("/hello-rust").
		WithExec([]string{"cargo", "r"})

	_, err = rust.Directory("output").Export(ctx, "output")
	if err != nil {
		return err
	}

	return nil
}
