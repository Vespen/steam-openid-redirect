package main

import (
	"github.com/Vespen/steam-openid-redirect/bootstrap"
	"github.com/sirupsen/logrus"
)

func main() {
	bootstrap.LoadConfigs(
		"server",
		"openid",
	)

	s := bootstrap.NewServer()

	if err := s.ListenAndServe(); err != nil {
		logrus.Fatal(err)
	}
}
