package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_ENDPOINT"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pong)

	return client
}

func setValue(keyID string, client *redis.Client) {
	err := client.Set(keyID, "hello world", 0).Err()

	if err != nil {
		log.Fatal(err)
	}
}

func getValue(keyID string, client *redis.Client) string {
	value, err := client.Get(keyID).Result()

	if err != nil {
		log.Fatal("error: ", err)
	}

	fmt.Printf("%v: %v\n", keyID, value)
	return value
}

func main() {
	sessionClient := newClient()
	setValue("mykey", sessionClient)
	getValue("mykey", sessionClient)
}
