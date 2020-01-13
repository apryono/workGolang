package usecase

import (
	"errors"
	"fmt"
	"tugas/GoPORT/dermaga"
	"tugas/GoPORT/model"
)

// DermagaUsecaseProcess struct
type DermagaUsecaseProcess struct {
	dermagaRepo dermaga.DermagaRepo
}

// CreateDermagaUsecase constructor
func CreateDermagaUsecase(dermagaRepo dermaga.DermagaRepo) dermaga.DermagaRepo {
	return &DermagaUsecaseProcess{dermagaRepo}
}

// GetAllDermaga function call for handler
func (d *DermagaUsecaseProcess) GetAllDermaga() (*[]model.Dermaga, error) {
	return d.dermagaRepo.GetAllDermaga()
}

// GetByID function call for handler
func (d *DermagaUsecaseProcess) GetByID(id int) (*model.Dermaga, error) {
	return d.dermagaRepo.GetByID(id)
}

// InsertData call for handler
func (d *DermagaUsecaseProcess) InsertData(dock *model.Dermaga) error {
	datadock, err := d.dermagaRepo.GetByID(dock.ID)
	if err != nil {
		return fmt.Errorf("[DermagaUsecaseProcess.InsertData.GetByID] Error when get InsertData by id' : %w", err)
	}

	if datadock != nil {
		return errors.New("Kode Dock sudah ada, silahkan masukkan kode lain")
	}

	return d.dermagaRepo.InsertData(dock)
}

// UpdateData function for handler
func (d *DermagaUsecaseProcess) UpdateData(id int, dock *model.Dermaga) error {
	return d.dermagaRepo.UpdateData(id, dock)
}

// DeleteByID function
func (d *DermagaUsecaseProcess) DeleteByID(id int) error {
	return d.dermagaRepo.DeleteByID(id)
}
