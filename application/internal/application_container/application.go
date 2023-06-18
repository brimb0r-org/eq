package application_container

import (
	"context"
	"gSheets/application/internal/eq_repo"
	"gSheets/application/internal/internal_config"
	"gSheets/application/internal/internal_excel"
	"gSheets/application/internal/translator/eq_translator"
	"gSheets/application/pkg/excel"
	"gSheets/application/pkg/produce"
	"gSheets/application/pkg/worker_pool"
	"github.com/brimb0r-org/scheduler/scheduler"
	"github.com/rs/zerolog/log"
	"time"
)

type runner struct {
	srv     *excel.ExcelConfig
	eqRepo  eq_repo.IEqRepo
	produce produce.IProduce
}

func BuildJob(c *internal_config.Configuration) scheduler.Job {
	srv := &excel.ExcelConfig{
		Apikey:      c.Excel.Apikey,
		AccessToken: c.Excel.AccessToken,
	}
	return &runner{
		srv:     srv,
		eqRepo:  &eq_repo.Repo{Database: c.Mongo.MongoDatabase()},
		produce: &produce.Produce{},
	}
}

func (app runner) Run() error {
	log.Print("begin")
	begin := time.Now()

	statusReport, err := internal_excel.GetDataToProcess(app.srv)
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

	work := worker_pool.Worker(func(i interface{}) {
		err = app.produce.Produce(i.(chan eq_translator.ITranslator))
	})
	for i := 0; i < len(eqChan); i++ {
		work <- eqChan
	}
	close(work)

	log.Print(time.Since(begin))
	return err
}