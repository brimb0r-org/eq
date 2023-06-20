package application_container

import (
	"context"
	"github.com/brimb0r-org/eq/application/internal/eq_repo"
	"github.com/brimb0r-org/eq/application/internal/internal_config"
	"github.com/brimb0r-org/eq/application/internal/internal_excel"
	"github.com/brimb0r-org/eq/application/internal/translator/eq_translator"
	"github.com/brimb0r-org/eq/application/pkg/excel"
	"github.com/brimb0r-org/eq/application/pkg/produce"
	"github.com/brimb0r-org/scheduler/scheduler"
	"github.com/rs/zerolog/log"
	"time"
)

type runner struct {
	excelsrv   *excel.ExcelConfig
	eqRepo     eq_repo.IEqRepo
	producesrv produce.IProduce
}

func BuildJob(c *internal_config.Configuration) scheduler.Job {
	excelsrv := &excel.ExcelConfig{
		Apikey:      c.Excel.Apikey,
		AccessToken: c.Excel.AccessToken,
	}
	producer := produce.NewProducer()
	return &runner{
		excelsrv:   excelsrv,
		eqRepo:     &eq_repo.Repo{Database: c.Mongo.MongoDatabase()},
		producesrv: producer,
	}
}

func (app runner) Run() error {
	log.Print("begin")
	begin := time.Now()

	statusReport, err := internal_excel.GetDataToProcess(app.excelsrv)
	if err != nil {
		log.Print(err)
	}

	eq, err := app.eqRepo.QueryEq(context.Background())
	eqChan := make(chan eq_translator.ITranslator, len(eq))

	for _, t := range eq {
		t.Activity = statusReport
		eqChan <- &eq_translator.EqTranslator{
			Eq:   t,
			Repo: app.eqRepo,
		}
	}

	close(eqChan)
	err = app.producesrv.Produce(eqChan)

	log.Print(time.Since(begin))
	return err
}
