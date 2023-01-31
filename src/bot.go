package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// Bot is the main struct for the bot
type Bot struct {
	Session *discordgo.Session
}

// NewBot creates a new bot
func NewBot(token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		Session: session,
	}, nil
}

// Run starts the bot
func (b *Bot) Run() error {
	fmt.Println("Starting bot...")
	return b.Session.Open()
}

// Send dm to user
func (b *Bot) Send(userID, message string) error {
	fmt.Printf("Sending message to %s: %s", userID, message)
	channel, err := b.Session.UserChannelCreate(userID)
	if err != nil {
		return err
	}

	_, err = b.Session.ChannelMessageSend(channel.ID, message)
	if err != nil {
		return err
	}

	return nil
}
