package dermaga

import (
	"tugas/GoPORT/model"
)

// DermagaRepo interface will be use as dermaga.DermagaRepo
type DermagaRepo interface {
	GetAllDermaga() (*[]model.Dermaga, error)
	GetByID(id int) (*model.Dermaga, error)
	InsertData(dock *model.Dermaga) error
	UpdateData(id int, dock *model.Dermaga) error
	DeleteByID(id int) error
}
