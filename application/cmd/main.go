package main

import (
	"gSheets/application/internal/application_container"
	"gSheets/application/internal/internal_config"
	"gSheets/application/internal/metadata"
	"github.com/brimb0r-org/scheduler/scheduler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func main() {
	runLogFile, _ := os.OpenFile(
		"gsheet.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0700,
	)
	multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
	metadata.NewMetaDataClient()
	c := internal_config.Configure()
	buildRunner := application_container.BuildJob(c)
	sched := scheduler.New()
	sched.ScheduleInterval(buildRunner, time.Second*time.Duration(c.Schedule))
	sched.Wait()
}
