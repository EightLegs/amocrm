package models

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const limit = 500
const delay = 25 * time.Millisecond
const refreshInterval = 1

type (
	amoSettings struct {
		Cfg    *config
		Client http.Client
	}

	config struct {
		Domain string `json:"domain"`
		Login  string `json:"login"`
		Key    string `json:"key"`
	}
)

var client *amoSettings

func OpenConnection(login, key, domain string) error {
	client = &amoSettings{
		Cfg: &config{
			Domain: domain,
			Key:    key,
			Login:  login,
		},
	}

	err := client.open()
	if err != nil {
		return err
	}

	go client.refresher()

	return nil
}

func (c *amoSettings) refresher() {
	ticker := time.NewTicker(refreshInterval * time.Minute)

	for {
		select {
		case t := <-ticker.C:
			log.Infof("Updating token at %s", t)
			err := client.open()
			if err != nil {
				log.Errorf("Got error while updating: %s", err)
			}
		}
	}
}

func (c *amoSettings) open() error {
	jar := newJar()
	cl := http.Client{Jar: jar, Timeout: 15 * time.Minute, CheckRedirect: nil, Transport: nil}
	// c.Client = http.Client{Jar: jar, Timeout: 15 * time.Minute, CheckRedirect: nil, Transport: nil}

	values := url.Values{
		"USER_LOGIN": {c.Cfg.Login},
		"USER_HASH":  {c.Cfg.Key},
	}

	time.Sleep(delay)
	resp, err := cl.PostForm(getUrl(authUrl), values)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}
	defer resp.Body.Close()

	c.Client = cl
	return nil
}

func getStrFromArr(arr []int) string {
	tmp := ""
	for _, value := range arr {
		if tmp != "" {
			tmp += ", "
		}
		tmp += strconv.Itoa(value)
	}

	return tmp
}
