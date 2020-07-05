package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/discord"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Missing files. Usage:", filepath.Base(os.Args[0]), "<files...>")
	}

	var token = os.Getenv("WEBHOOK_TOKEN")
	var strid = os.Getenv("WEBHOOK_ID")

	// Parse string ID.
	w, err := discord.ParseSnowflake(strid)
	if err != nil {
		log.Fatalln("Invalid snowflake in $WEBHOOK_ID:", err)
	}

	if token == "" {
		log.Fatalln("$WEBHOOK_TOKEN is empty.")
	}

	var names = os.Args[1:]
	var files = make([]api.SendMessageFile, len(names))

	for i, name := range names {
		f, err := os.Open(name)
		if err != nil {
			log.Fatalln("Failed to open file:", err)
		}

		// Close the file at the END of the MAIN FUNCTION. This defer is
		// intentional!
		defer f.Close()

		files[i] = api.SendMessageFile{
			Name:   filepath.Base(name),
			Reader: f,
		}
	}

	c := api.NewClient("")

	m, err := c.ExecuteWebhook(w, token, true, api.ExecuteWebhookData{Files: files})
	if err != nil {
		log.Fatalln("Failed to send webhook:", err)
	}

	if len(m.Attachments) == 0 {
		log.Fatalln("No attachments found in message.")
	}

	os.Stdout.WriteString(m.Attachments[0].URL)
}
