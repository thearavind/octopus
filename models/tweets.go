package models

type Tweets struct {
	ID            int         `json:"id" bson:"id"`
	CreatedAt     string      `json:"created_at" bson:"createdAt"`
	FavoriteCount int         `json:"favorite_count" bson:"favoriteCount"`
	Favorited     bool        `json:"favorited" bson:"favorited"`
	RetweetCount  int         `json:"retweet_count" bson:"retweetCount"`
	Retweeted     bool        `json:"retweeted" bson:"retweeted"`
	Text          string      `json:"text" bson:"text"`
	Truncated     bool        `json:"truncated" bson:"truncated"`
}
