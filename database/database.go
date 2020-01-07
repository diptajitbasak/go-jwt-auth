package database

import ("github.com/go-redis/redis"
"fmt"
)

func ExampleNewClient() {
	fmt.Println("database connected")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}