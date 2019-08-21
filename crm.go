package amocrm

import (
	"github.com/vozerov/amocrm/models"
)

type amo struct {
	Account models.Acc
	Company models.Cmpn
	Lead    models.Ld
	Contact models.Ct
	Task    models.Tsk
	Note    models.Nt
}

func NewAmo(login, key, domain string) *amo {
	models.OpenConnection(login, key, domain)
	return &amo{}
}
