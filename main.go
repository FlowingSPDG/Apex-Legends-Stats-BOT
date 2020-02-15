package main

import (
	"encoding/json"
	"flag"
	"fmt"
	models "github.com/FlowingSPDG/Apex-Legends-Stats-BOT/src"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// StatsCommand is Discord Command for gettings stats
	StatsCommand = ".stats"
)

var (
	// DiscordToken for discordgo
	DiscordToken *string
	// TrackerToken for tracker.gg
	TrackerToken *string
	stopBot      = make(chan struct{})
)

func init() {
	DiscordToken = flag.String("discord", "", "Discord APP token. e.g. NTQwXX...")
	TrackerToken = flag.String("tracker", "", "tracker.gg APP token. e.g. 592efec1-XXX...")
	flag.Parse()
	*DiscordToken = "Bot " + *DiscordToken
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}
	if strings.HasPrefix(m.Content, StatsCommand) {
		commands := strings.Split(m.Content, " ")
		log.Printf("Command Received from user : %s. Commands : %v\n", m.Author.Username, commands)
		if len(commands) < 2 {
			embed := &discordgo.MessageEmbed{
				Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
				Title:     "❌エラー",
				Color:     0xff0000, // RED
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "コマンド送信者",
						Value:  fmt.Sprintf("<@%s>", m.Author.ID),
						Inline: false,
					},
					&discordgo.MessageEmbedField{
						Name:   "エラー",
						Value:  "引数が足りません。",
						Inline: false,
					},
				},
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if err != nil {
				log.Println("Discord send message ERR : ", err)
				return
			}
			return
		}
		platform := "origin"
		id := commands[1]
		client := &http.Client{}
		req, err := http.NewRequest("GET", fmt.Sprintf("https://public-api.tracker.gg/v2/apex/standard/profile/%s/%s", platform, id), nil)
		req.Header.Add("TRN-Api-Key", *TrackerToken)
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("tracker.gg Request ERR : %v\n", err)
			embed := &discordgo.MessageEmbed{
				Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
				Title:     "❌エラー",
				Color:     0xff0000, // RED
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "コマンド送信者",
						Value:  fmt.Sprintf("<@%s>", m.Author.ID),
						Inline: false,
					},
					&discordgo.MessageEmbedField{
						Name:   "エラー",
						Value:  fmt.Sprintf("HTTP GET ERROR : %v", err.Error()),
						Inline: false,
					},
				},
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if err != nil {
				log.Println("Discord send message ERR : ", err)
				return
			}
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("tracker.gg Request ERR : %v, BODY : %s\n", err, string(body))
			embed := &discordgo.MessageEmbed{
				Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
				Title:     "❌エラー",
				Color:     0xff0000, // RED
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "コマンド送信者",
						Value:  fmt.Sprintf("<@%s>", m.Author.ID),
						Inline: false,
					},
					&discordgo.MessageEmbedField{
						Name:   "エラー",
						Value:  fmt.Sprintf("JSON Unmarshal : %v", err.Error()),
						Inline: false,
					},
					&discordgo.MessageEmbedField{
						Name:   "JSON Body",
						Value:  string(body),
						Inline: false,
					},
				},
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if err != nil {
				log.Println("Discord send message ERR : ", err)
				return
			}
			return
		}
		stats := models.ProfileStatsResponse{}
		err = json.Unmarshal(body, &stats)
		if err != nil {
			log.Printf("JSON Unmarshaling ERR : %v, BODY : %s\n", err, string(body))
			embed := &discordgo.MessageEmbed{
				Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
				Title:     "❌エラー",
				Color:     0xff0000, // RED
				Fields: []*discordgo.MessageEmbedField{
					&discordgo.MessageEmbedField{
						Name:   "コマンド送信者",
						Value:  fmt.Sprintf("<@%s>", m.Author.ID),
						Inline: false,
					},
					&discordgo.MessageEmbedField{
						Name:   "エラー",
						Value:  fmt.Sprintf("HTTP GET ERROR : %v", err.Error()),
						Inline: false,
					},
				},
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
			if err != nil {
				log.Println("Discord send message ERR : ", err)
				return
			}
			return
		}
		log.Printf("stats : %v\n", stats)

		var rank string
		switch {
		case stats.Data.Segments[0].Stats.RankScore.Value > 10000:
			rank = "Master"
		case stats.Data.Segments[0].Stats.RankScore.Value > 7200:
			rank = "Diamond"
		case stats.Data.Segments[0].Stats.RankScore.Value > 4800:
			rank = "Platinum"
		case stats.Data.Segments[0].Stats.RankScore.Value > 4300:
			rank = "Gold I"
		case stats.Data.Segments[0].Stats.RankScore.Value > 3800:
			rank = "Gold II"
		case stats.Data.Segments[0].Stats.RankScore.Value > 3300:
			rank = "Gold III"
		case stats.Data.Segments[0].Stats.RankScore.Value > 2800:
			rank = "Gold IV"
		case stats.Data.Segments[0].Stats.RankScore.Value > 2400:
			rank = "Silver I"
		case stats.Data.Segments[0].Stats.RankScore.Value > 2000:
			rank = "Silver II"
		case stats.Data.Segments[0].Stats.RankScore.Value > 1600:
			rank = "Silver III"
		case stats.Data.Segments[0].Stats.RankScore.Value > 1200:
			rank = "Silver IV"
		case stats.Data.Segments[0].Stats.RankScore.Value > 1200:
			rank = "Bronze I"
		case stats.Data.Segments[0].Stats.RankScore.Value > 900:
			rank = "Bronze II"
		case stats.Data.Segments[0].Stats.RankScore.Value > 600:
			rank = "Bronze III"
		case stats.Data.Segments[0].Stats.RankScore.Value > 300:
			rank = "Bronze IV"
		}
		embed := &discordgo.MessageEmbed{
			Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
			Title:     "✅STATS取得成功",
			Color:     0x00ff00, // Green
			Image: &discordgo.MessageEmbedImage{
				URL: stats.Data.PlatformInfo.AvatarURL,
			},
			Fields: []*discordgo.MessageEmbedField{
				&discordgo.MessageEmbedField{
					Name:   "コマンド送信者",
					Value:  fmt.Sprintf("<@%s>", m.Author.ID),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "STATS取得対象ID",
					Value:  stats.Data.PlatformInfo.PlatformUserID,
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "LEVEL",
					Value:  strconv.Itoa(int(stats.Data.Segments[0].Stats.Level.Value)),
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "RANK",
					Value:  rank,
					Inline: false,
				},
				&discordgo.MessageEmbedField{
					Name:   "Kills",
					Value:  strconv.Itoa(int(stats.Data.Segments[0].Stats.Kills.Value)),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "Damages",
					Value:  strconv.Itoa(int(stats.Data.Segments[0].Stats.Damage.Value)),
					Inline: true,
				},
				&discordgo.MessageEmbedField{
					Name:   "現在のレジェンド",
					Value:  stats.Data.Metadata.ActiveLegendName,
					Inline: false,
				},
			},
		}
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)
		if err != nil {
			log.Println("Discord send message ERR : ", err)
			return
		}
	}
}

func main() {
	//Discordのセッションを作成
	discord, err := discordgo.New()
	discord.Token = *DiscordToken
	if err != nil {
		log.Println("Error logging in")
		log.Println(err)
	}

	discord.AddHandler(onMessageCreate) //全てのWS APIイベントが発生した時のイベントハンドラを追加
	// websocketを開いてlistening開始
	openerr := discord.Open()
	if openerr != nil {
		log.Printf("ERR : %v\n", openerr)
		panic(openerr)
	}
	botName := fmt.Sprintf("<@!%s>", discord.State.User.ID)
	log.Printf("Logged in Discord BOT %s\n", botName)
	<-stopBot //プログラムが終了しないようロック
	return
}
