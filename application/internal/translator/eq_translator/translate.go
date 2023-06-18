package eq_translator

import (
	"gSheets/application/internal/eq_repo"
)

type ITranslator interface {
	SendSuccessCallback() error
	Translate()
}

type EqTranslator struct {
	Eq   *eq_repo.Eq
	Repo eq_repo.IEqRepo
}

func (t *EqTranslator) SendSuccessCallback() error {
	return t.Repo.UpdateEqPublished(t.Eq)
}

func (t *EqTranslator) Translate() {}
