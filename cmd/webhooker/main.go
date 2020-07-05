package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/diamondburned/arikawa/api"
	"github.com/diamondburned/arikawa/discord"
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

	var names = os.Args[2:]
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

	m, err := c.ExecuteWebhook(w, t, true, api.ExecuteWebhookData{Files: files})
	if err != nil {
		log.Fatalln("Failed to send webhook:", err)
	}

	if len(m.Attachments) == 0 {
		log.Fatalln("No attachments found in message.")
	}

	os.Stdout.WriteString(m.Attachments[0].URL)
}
