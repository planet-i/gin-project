package cron

import (
	"log"
	"time"

	"github.com/planet-i/gin-project/models"
	"github.com/robfig/cron/v3"
)

func xronJob() {
	log.Println("Starting....")
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag")
		models.CleanAllTag()
	})

	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		// 阻塞 select 等待 channel
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
