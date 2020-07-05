package main

import (
	"testing"

	"github.com/Bios-Marcel/discordemojimap"
)

func TestBlobPrefixName(t *testing.T) {
	var tests = []struct{ emoji, prefix string }{
		{"👨‍❤️‍💋‍👨", "emoji_u1f468_200d_2764_200d_1f48b_200d_1f468"},
	}

	for _, test := range tests {
		if name := blobPrefixName(test.emoji); name != test.prefix {
			t.Fatalf("Unexpected name %q for %q", name, test.emoji)
		}
	}
}

func TestDiscordEmojisExist(t *testing.T) {
	for name, emoji := range discordemojimap.EmojiMap {
		p, err := BlobExists(emoji)
		if err != nil {
			t.Logf("Emoji %q does not exist at %q, error %v.", name, p, err)
		}
	}
}
