# About
iTunesで再生中の曲をWSHで取得してGoでTwitterに投稿します。

# Usage
Twitter API を利用するためのトークンを設定します。

`twitter_oauth.env`
```twitter_oauth.env
TWITTER_CONSUMER_KEY=
TWITTER_CONSUMER_SECRET=
TWITTER_OAUTH_TOKEN=
TWITTER_OAUTH_TOKEN_SECRET=
```

あとは実行するだけ。

`go run main.go`
