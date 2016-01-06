package main

import (
	"os/exec"
	"fmt"
	"log"
	"bufio"
	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"os"
	"net/url"
	"io/ioutil"
"strings"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/japanese"
)

func main() {
	err := godotenv.Load("./twitter_oauth.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))

	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_OAUTH_TOKEN"), os.Getenv("TWITTER_OAUTH_TOKEN_SECRET"))
	defer api.Close()

	ch := make(chan string)
	go func() {
		err := run(ch, "cscript.exe", "//nologo", "./now_playing.js")
		if err != nil {
			log.Fatal(err)
		}
	}()

	for v := range ch {
		api.PostTweet(v, url.Values{})
		fmt.Println(v)
	}
}

func run(ch chan<- string, command string, flags ...string) error {
	cmd := exec.Command(command, flags...)

	output, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("RunCommand: cmd.StdoutPipe(): %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("RunCommand: cmd.Start(): %v", err)
	}

	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		ch <- sjis_to_utf8(scanner.Text())
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("RunCommand: cmd.Wait(): %v", err)
	}

	return nil
}

func sjis_to_utf8(str string) string {
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		panic(err)
	}
	return string(ret)
}
