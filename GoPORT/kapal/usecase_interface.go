package kapal

import "tugas/GoPORT/model"

// KapalUsecase interface will be use as kapal.KapalUsecase
type KapalUsecase interface {
	GetAllKapal() (*[]model.Kapal, error)
	GetByID(id int) (*model.Kapal, error)
	InsertData(kapal *model.Kapal) error
	UpdateData(id int, kapal *model.Kapal) error
	DeleteByID(id int) error
}
