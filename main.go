package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
           Addr: "database:6379",
        })

	err := client.Set(client.Context(),"name", "Laura", 0).Err()
	if err != nil {
		panic(err)
	}

	nameVal, err := client.Get(client.Context(), "name").Result()
	if err != nil {
		panic(err)
	}

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Your Redis value is: %q", nameVal)
	})

	log.Fatal(http.ListenAndServe(":9001", nil))

}
