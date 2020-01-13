package dermaga

import (
	"tugas/GoPORT/model"
)

// DermagaUsecase interface will be use as dermaga.DermagaUsecase
type DermagaUsecase interface {
	GetAllDermaga() (*[]model.Dermaga, error)
	GetByID(id int) (*model.Dermaga, error)
	InsertData(dock *model.Dermaga) error
	UpdateData(id int, dock *model.Dermaga) error
	DeleteByID(id int) error
}
