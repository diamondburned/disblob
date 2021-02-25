package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/diamondburned/arikawa/v2/api/webhook"
	"github.com/diamondburned/arikawa/v2/discord"
	"github.com/diamondburned/arikawa/v2/utils/sendpart"
)

var webhookURL = regexp.MustCompile(`https://discord.com/api/webhooks/(\d+)/(\S+)`)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("Usage:", filepath.Base(os.Args[0]), "<webhook URL file> <files...>")
	}

	u, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("Failed to read webhook URL file:", err)
	}

	matches := webhookURL.FindStringSubmatch(string(u))
	if len(matches) != 3 {
		log.Fatalln("Invalid URL format.")
	}

	// Parse string ID.
	w, err := discord.ParseSnowflake(matches[1])
	if err != nil {
		log.Fatalln("Invalid snowflake in $WEBHOOK_ID:", err)
	}
	// Get the token.
	t := matches[2]

	client := webhook.New(discord.WebhookID(w), t)
	names := os.Args[2:]

	for _, name := range names {
		f, err := os.Open(name)
		if err != nil {
			log.Fatalln("Failed to open file:", err)
		}

		m, err := client.ExecuteAndWait(webhook.ExecuteData{
			Files: []sendpart.File{
				{Name: filepath.Base(name), Reader: f},
			},
		})
		if err != nil {
			log.Fatalln("Failed to send webhook:", err)
		}

		if len(m.Attachments) == 0 {
			log.Fatalln("No attachments found in message.")
		}

		os.Stdout.WriteString(m.Attachments[0].URL)
		os.Stdout.WriteString("\n")

		// Close file; no need to chech for errors here.
		f.Close()
	}
}
