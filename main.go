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

func mustConfig () Config{
	tz := getenv("TIMEZONE", "Asia/Bangkok")
	dueDay := mustInt(getenv("DUE_DAY", "15"))

	lineToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")
	if lineToken == ""{
		log.Fatal("missing LINE_CHANNEL_ACCESS_TOKEN (inject from Secret Manager in Cloud Run)")
	}

	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == ""{
		projectID = os.Getenv("GCP_PROJECT")
	}
	if projectID == ""{
		log.Fatal("missing project id")
	}

	return Config{
		Timezone: tz,
		DueDay: dueDay,
		GroupID: os.Getenv("GROUP_ID"),
		AmountTHB: getenv("AMOUNT_THB", "150"),
		MessageTemplate: getenv("MESSAGE_TEMPLATE", "อย่าลืมจ่ายค่า Apple One เดือนนี้ %s บาทนะครับ 🙏"),
		LineAccessToken: lineToken,
		LineChannelSecret: os.Getenv("LINE_CHANNEL_SECRET"),
		FirestoreProjectID: projectID,		
	}
}

func main() {
	cfg := mustConfig()
}