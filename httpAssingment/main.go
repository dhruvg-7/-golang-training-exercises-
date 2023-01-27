package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func RootHandler() func(w http.ResponseWriter, req *http.Request) {
	visited := make([]string, 0)

	return func(w http.ResponseWriter, req *http.Request) {
		root := req.URL.Path[1:]
		for i := range visited {
			if visited[i] == root {
				ans := "Welcome back " + root + "!"
				fmt.Fprintln(w, ans)
				return
			}
		}
		visited = append(visited, root)
		ans := "Greetings " + root + "!"
		fmt.Fprintln(w, ans)

	}
}

func main() {

	f := RootHandler()

	err := http.ListenAndServe(":8000", http.HandlerFunc(f))
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
