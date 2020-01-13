package kapal

import (
	"tugas/GoPORT/model"
)

// KapalRepo interface will be use as kapal.KapalRepo
type KapalRepo interface {
	GetAllKapal() (*[]model.Kapal, error)
	GetByID(id int) (*model.Kapal, error)
	InsertData(kapal *model.Kapal) error
	UpdateData(id int, kapal *model.Kapal) error
	DeleteByID(id int) error
}
