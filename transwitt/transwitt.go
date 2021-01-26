package transwitt

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (tc TwitterConfig) Vaildate() error {
	return nil
}

func getTimeline(clientTwitter *twitter.Client, timelineParams twitter.UserTimelineParams) ([]twitter.Tweet, int, error) {
	arTweets, httpResponse, err := clientTwitter.Timelines.UserTimeline(&timelineParams)
	if err != nil {
		return nil, -1, err
	}
	if httpResponse.StatusCode != 200 {
		body, err := ioutil.ReadAll(httpResponse.Body)
		if err != nil {
			return nil, httpResponse.StatusCode, err
		}
		return nil, httpResponse.StatusCode, errors.New(string(body))
	}
	return arTweets, httpResponse.StatusCode, nil
}

func Run(opconf OperateConfig) error {
	if err := opconf.Twitter.Vaildate(); err != nil {
		return err
	}

	var (
		telegram *tgbotapi.BotAPI
		nAdminID int64
		err      error
	)
	if opconf.Messenger.Telegram != (TelegramConfig{}) {
		telegram, err = tgbotapi.NewBotAPI(opconf.Messenger.Telegram.Token)
		if err != nil {
			return err
		}
		nAdminID = opconf.Messenger.Telegram.Admin
		log.Printf("Telegram bot [%s] is authorized", telegram.Self.UserName)
		/*
			ucUpdates := tgbotapi.NewUpdate(0)
			ucUpdates.Timeout = 60
			chanUpdate, err := telegram.GetUpdatesChan(ucUpdates)
			if err != nil {
				return err
			}
			// telegram command listener

			go func() {
				for {
					select {
					case u := <-chanUpdate:
						log.Println("updates : ", chanUpdate)
						log.Println("From :", u.Message.From.String())
						log.Println("Date :", u.Message.Date)
						log.Println("Chat.ID :", u.Message.Chat.ID)
						log.Println("Chat.UserName :", u.Message.Chat.UserName)
						log.Println("Text :", u.Message.Text)
					}
				}
			}()
		*/
	}

	// oauth2 configures a client that uses app credentials to keep a fresh token
	confTwitter := &clientcredentials.Config{
		ClientID:     opconf.Twitter.ConsumerKey,
		ClientSecret: opconf.Twitter.ConsumerSecret,
		TokenURL:     "https://api.twitter.com/oauth2/token",
	}
	// http.Client will automatically authorize Requests
	httpClient := confTwitter.Client(oauth2.NoContext)

	// Twitter client
	clientTwitter := twitter.NewClient(httpClient)
	go func() {
		var (
			bExcludeReplies  = false
			bIncludeRetweets = true
			timelineParams   = twitter.UserTimelineParams{}
			tInterval        = time.Millisecond * time.Duration((1/opconf.Twitter.TPS)*1000)
			tLastRequested   = time.Time{}
			tCreatedAt       = time.Time{}
			Users            = TwitterUsers{}
		)

		// init
		for _, u := range opconf.Twitter.Users {
			timelineParams = twitter.UserTimelineParams{
				ScreenName:      u.ScreenID,
				Count:           1,
				ExcludeReplies:  &bExcludeReplies,
				IncludeRetweets: &bIncludeRetweets,
				TweetMode:       "extended",
			}
			for {
				if time.Now().Sub(tLastRequested) > tInterval {
					break
				}
				time.Sleep(time.Millisecond * 1) // Thread safe
			}
			arTweets, nStatusCode, err := getTimeline(clientTwitter, timelineParams)
			tLastRequested = time.Now()
			if err != nil {
				log.Printf("Fail to init user %s with error %s (%d). Remove this user\r\n", timelineParams.ScreenName, err, nStatusCode)
				continue
			}
			if arTweets[0].Retweeted {
				u.TweetTime, err = arTweets[0].RetweetedStatus.CreatedAtTime()
			} else {
				u.TweetTime, err = arTweets[0].CreatedAtTime()
			}
			if err != nil {
				log.Printf("Fail to init user %s with error %s. Remove this user\r\n", timelineParams.ScreenName, err)
				continue
			}
			if u.Nickname == "" {
				u.Nickname = u.ScreenID
			}
			log.Println("Init Twitter user", u.ScreenID, "as", u.Nickname)
			Users = append(Users, u)
		}
		for {
			for i, u := range Users {
				for {
					if time.Now().Sub(tLastRequested) > tInterval {
						break
					}
					time.Sleep(time.Millisecond * 1) // Thread safe
				}

				timelineParams = twitter.UserTimelineParams{
					ScreenName:      u.ScreenID,
					Count:           5,
					ExcludeReplies:  &bExcludeReplies,
					IncludeRetweets: &bIncludeRetweets,
					TweetMode:       "extended",
				}
				arTweets, nStatusCode, err := getTimeline(clientTwitter, timelineParams)
				tLastRequested = time.Now()
				if err != nil {
					log.Printf("Fail to get timeline %s with error %s (%d)\r\n", timelineParams.ScreenName, err, nStatusCode)
					continue
				}

				for _, t := range arTweets {
					// Convert time
					if t.Retweeted {
						tCreatedAt, err = t.RetweetedStatus.CreatedAtTime()
					} else {
						tCreatedAt, err = t.CreatedAtTime()
					}
					if err != nil {
						log.Printf("Fail to convert CreatedAtTime user %s with error %s\r\n", timelineParams.ScreenName, err)
						continue
					}

					// Compare time
					if u.TweetTime.Before(tCreatedAt) {
						log.Println(tCreatedAt.Local().String())
						log.Println(t.FullText)
						Users[i].TweetTime = tCreatedAt

						sMessage := fmt.Sprintf("[%s] from [%s]\r\n%s", tCreatedAt.Local().String(), u.Nickname, t.FullText)
						msg := tgbotapi.NewMessage(nAdminID, sMessage)
						_, err = telegram.Send(msg)
						if err != nil {
							log.Println("Fail to send to Telegram with error", err)
						}
					}
				}
			}
		}
	}()

	return nil
}
