package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"highest-twitterbot/utils"
	"net/url"
	"sort"
)

func main() {
	v := url.Values{}
	v.Set("count", utils.Count)
	v.Set("since_id", utils.BenchID)
	api := anaconda.NewTwitterApiWithCredentials(utils.AccessToken, utils.AccessTokenSecret, utils.ConsumerKey, utils.ConsumerSecret)
	search, _ := api.GetSearch(utils.Query, v)

	var users []utils.User
	for _, i := range search.Statuses {
		incremented := false
		for n, j := range users {
			if j.ScreenName == i.User.ScreenName {
				users[n].Tweets = j.Tweets + 10
				incremented = true
				break
			}
		}

		if !incremented {
			users = append(users, utils.User{Name: i.User.Name, ScreenName: i.User.ScreenName, Tweets: 1})
		}
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Tweets < users[j].Tweets
	})
	if len(users) != 0 {
		for _, tweetEntry := range users {
			if tweetEntry.Tweets == 1 {
				fmt.Printf("@%s made %d tweet \n", tweetEntry.ScreenName, tweetEntry.Tweets)
				continue
			}
			fmt.Printf("@%s made %d tweets \n", tweetEntry.ScreenName, tweetEntry.Tweets)

		}
		fmt.Println("\nTweet Retrieval and Sort done")
	}else{
		fmt.Println("No tweets based on your params")
	}

}
