package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type User struct {
	ID       int
	Username string
	Password string
}

func main() {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address and port
		Password: "",               // Redis password (if any)
		DB:       0,                // Redis database index
	})

	// Close the Redis connection when the program finishes
	defer client.Close()

	// Create a User struct
	user := User{
		ID:       1,
		Username: "cecep",
		Password: "123",
	}

	// Serialize the User struct to JSON
	userJSON, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	// Store the JSON data in Redis with a key
	err = client.Set(context.Background(), "user", userJSON, time.Hour).Err()
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the JSON data from Redis using the key
	val, err := client.Get(context.Background(), "user").Result()
	if err != nil {
		log.Fatal(err)
	}

	// Deserialize the JSON data back into a User struct
	var retrievedUser User
	err = json.Unmarshal([]byte(val), &retrievedUser)
	if err != nil {
		log.Fatal(err)
	}

	// Print the retrieved User struct
	fmt.Println("ID:", retrievedUser.ID)
	fmt.Println("Username:", retrievedUser.Username)
	fmt.Println("Password:", retrievedUser.Password)

	// Delete the data from Redis
	err = client.Del(context.Background(), "user").Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data has been deleted")
}
