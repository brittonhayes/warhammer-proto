package server

import (
	"context"
	"os"

	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
)

func setupSession() *session.Session {
	channelID := os.Getenv("CHANNEL_ID")
	if channelID == "" {
		log.Fatal().Msg("No $CHANNEL_ID given.")
	}

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal().Msg("No $BOT_TOKEN given.")
	}

	s, err := session.New("Bot " + token)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open session")
	}

	s.AddIntents(gateway.IntentGuilds)
	s.AddIntents(gateway.IntentGuildMessages)
	if err := s.Open(context.Background()); err != nil {
		log.Panic().Err(err).Msg("failed to open session")
	}

	return s
}
