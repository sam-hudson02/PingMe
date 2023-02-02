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

func (b *Bot) BuildPipelineNotify(channel_id string, buildRequest BuildRequest) error {
	fmt.Printf("Sending build notification to %s: %s", channel_id, buildRequest.BuildID)
	// build rich embed
	// set discord embed color
	color := 2031360
	Title := ""
	fmt.Printf("Build Result: %s", buildRequest.BuildResult)
	if buildRequest.BuildResult == "SUCCESS" {
		Title = "✅ " + buildRequest.BuildName + " Build Success! ✅"
		// green
		color = 2031360
	} else {
		Title = "❌ " + buildRequest.BuildName + " Build Failure! ❌"
		// red
		color = 16711680
	}

	msgEmbed := &discordgo.MessageEmbed{
		Title:       Title,
		Description: buildRequest.BuildURL,
		Color:       color,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Build Number",
				Value:  buildRequest.BuildID,
				Inline: false,
			},
			{
				Name:   "Commit Author",
				Value:  buildRequest.CommitAuthor,
				Inline: false,
			},
			{
				Name:   "Commit URL",
				Value:  buildRequest.CommitURL,
				Inline: false,
			},
			{
				Name:   "Build Duration",
				Value:  buildRequest.BuildDuration,
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: buildRequest.BuildDate,
		},
	}

	// get user channel
	b.Session.ChannelMessageSendEmbed(channel_id, msgEmbed)
	return nil
}
