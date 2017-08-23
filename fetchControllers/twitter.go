package fetchControllers

import (
	"github.com/kapitol-app/octopus/fetchers"
	"github.com/kapitol-app/octopus/db"
	"github.com/kapitol-app/octopus/models"
	"github.com/kapitol-app/octopus/logger"
)

type TweetsFetchController struct{}

type tweetResponse struct {
	Statuses []models.Tweets `json:"statuses"`
}

var accessToken string

func (tfc *TweetsFetchController) fetchAccessToken() {
	accessToken = fetchers.FetchTwitterBearerToken()
}

func (tfc *TweetsFetchController) fetchMembersTweet(memberDbCollection db.CollectionType) {
	tfc.fetchAccessToken()
	queryResult := []models.Member{}
	tweetResult := tweetResponse{}
	query, _ := db.Find(nil, memberDbCollection)
	query.All(&queryResult)
	for _, congressMember := range queryResult {
		if congressMember.TwitterAccount == "" {
			logger.Log("Member", congressMember.FirstName, "does not have a twitter account")
		} else {
			fetchers.FetchTweets(congressMember.TwitterAccount, accessToken, &tweetResult)
			if len(tweetResult.Statuses) == 0 {
				logger.Log("Cant find new tweets by", congressMember.TwitterAccount)
			} else if congressMember.Tweets == nil {
				congressMember.Tweets = &tweetResult.Statuses
				err := db.Update(congressMember, congressMember.ID, memberDbCollection)
				if err != nil {
					logger.Error("Failed to update the tweets of", congressMember.TwitterAccount, err)
				} else {
					logger.Log("Fetched", len(tweetResult.Statuses), "new tweets from", congressMember.TwitterAccount)
				}
			} else {
				tweetCount := 0
				for _, tweet := range tweetResult.Statuses {
					if !checkTweetExists(*congressMember.Tweets, tweet.ID) {
						*congressMember.Tweets = append(*congressMember.Tweets, tweet)
						tweetCount = tweetCount + 1
					}
				}
				if tweetCount != 0 {
					err := db.Update(congressMember, congressMember.ID, memberDbCollection)
					if err != nil {
						logger.Error("Failed to update the new tweet of", congressMember.TwitterAccount, err)
					} else {
						logger.Log("Fetched", tweetCount, "new tweets by", congressMember.TwitterAccount)
					}
				} else {
					logger.Log("Cant find new tweets by", congressMember.TwitterAccount)
				}
			}
		}
	}
}

func (tfc *TweetsFetchController) FetchSenatorsAndHouseTweet() {
	tfc.fetchMembersTweet(db.SenateCollection)
	tfc.fetchMembersTweet(db.HouseCollection)
}

func checkTweetExists(tweets []models.Tweets, tweetId int) bool {
	for _, oldTweet := range tweets {
		if oldTweet.ID == tweetId {
			return true
			break
		}
	}
	return false
}
