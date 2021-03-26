package amocrm

import (
	"github.com/bkmz/amocrm/models"
)

type Amo struct {
	Account models.Acc
	Company models.Cmpn
	Lead    models.Ld
	Contact models.Ct
	Task    models.Tsk
	Note    models.Nt
}

func NewAmo(login, key, domain string) *Amo {
	models.OpenConnection(login, key, domain)
	return &Amo{}
}
