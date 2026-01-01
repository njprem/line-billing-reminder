package main

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Timezone string
	DueDay int
	GroupID string
	AmountTHB string
	MessageTemplate string
	LineAccessToken string
	LineChannelSecret string
	FirestoreProjectID string
}

func getenv(k, def string) string{
	v := os.Getenv(k)
	if v == "" {
		return def
	}
	return v
}

func mustInt (s string) int{
	i, err := strconv.Atoi(s)
	if err != nil{
		log.Fatalf("invalid int: %s", s)
	}
	return i
}

func main() {
	
}