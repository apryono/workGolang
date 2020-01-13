package usecase

import (
	"errors"
	"fmt"
	"tugas/GoPORT/kapal"
	"tugas/GoPORT/model"
)

// KapalUsecaseProcess struct
type KapalUsecaseProcess struct {
	kapalRepo kapal.KapalRepo
}

// CreateKapalUsecase constructor
func CreateKapalUsecase(kapalRepo kapal.KapalRepo) kapal.KapalUsecase {
	return &KapalUsecaseProcess{kapalRepo}
}

// GetAllKapal function call for handler
func (k *KapalUsecaseProcess) GetAllKapal() (*[]model.Kapal, error) {
	return k.kapalRepo.GetAllKapal()
}

// GetByID function call for handler
func (k *KapalUsecaseProcess) GetByID(id int) (*model.Kapal, error) {
	return k.kapalRepo.GetByID(id)
}

// DeleteByID function call for handler
func (k *KapalUsecaseProcess) DeleteByID(id int) error {
	return k.kapalRepo.DeleteByID(id)
}

// InsertData function call for handler
func (k *KapalUsecaseProcess) InsertData(kapal *model.Kapal) error {
	datakapal, err := k.kapalRepo.GetByID(kapal.ID)
	if err != nil {
		return fmt.Errorf("[KapalUsecaseProcess.InsertData.GetByID] Error when get InsertData by id' : %w", err)
	}

	if datakapal != nil {
		return errors.New("Kode yang di input sudah ada, silahkan masukkan kode lain")
	}

	return k.kapalRepo.InsertData(kapal)
}

// UpdateData function call for handler
func (k *KapalUsecaseProcess) UpdateData(id int, kapal *model.Kapal) error {
	return k.kapalRepo.UpdateData(id, kapal)
}
