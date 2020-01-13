package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"tugas/GoPORT/kapal"
	"tugas/GoPORT/model"
)

// KapalRepoMysql struct
type KapalRepoMysql struct {
	db *sql.DB
}

// CreateKapalRepoMysql function will call for func main() and send to usecase
func CreateKapalRepoMysql(db *sql.DB) kapal.KapalRepo {
	return &KapalRepoMysql{db}
}

// GetAllKapal function
func (k *KapalRepoMysql) GetAllKapal() (*[]model.Kapal, error) {
	stats := 0
	qry := "select kapal_id, kode_kapal, muatan, status from kapal where is_delete = ?"
	allkapal := []model.Kapal{}
	kapal := model.Kapal{}

	rows, err := k.db.Query(qry, stats)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[KapalRepoMysql.CreateKapalRepoMysql.GetAllKapal.QueryRow] Error when query data with : %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&kapal.ID, &kapal.Kode, &kapal.Muatan, &kapal.Status); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[KapalRepoMysql.CreateKapalRepoMysql.GetAllKapal.Scan] Error when query data with : %w", err)
		}
		allkapal = append(allkapal, kapal)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[KapalRepoMysql.CreateKapalRepoMysql.GetAllKapal.rows.Err()] Error when rows.Err() data with : %w", err)
	}

	return &allkapal, nil
}

// GetByID function call for usecase
func (k *KapalRepoMysql) GetByID(id int) (*model.Kapal, error) {
	qry := "select kapal_id, kode_kapal, muatan, status from kapal where kapal_id = ? and is_delete = 0"

	kapal := model.Kapal{}

	err := k.db.QueryRow(qry, id).Scan(&kapal.ID, &kapal.Kode, &kapal.Muatan, &kapal.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[KapalRepoMysql.CreateTableRepoMysqlImpl.GetByID.QueryRow] Error when running query '"+qry+"' : %w", err)
	}
	return &kapal, nil
}

// DeleteByID function to soft delete
func (k *KapalRepoMysql) DeleteByID(id int) error {
	qry := "update kapal set is_delete = 1 where kapal_id = ?"

	_, err := k.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("[KapalRepoMysql.CreateTableRepoMysqlImpl.DeleteByID.Exec] Error when running query '"+qry+"' : %w", err)
	}
	return nil
}

// InsertData function to soft delete
func (k *KapalRepoMysql) InsertData(kapal *model.Kapal) error {
	qry := "insert into kapal (kode_kapal,muatan,status,is_delete) values (?,?,?,?)"

	tx, err := k.db.Begin()
	if err != nil {
		return fmt.Errorf("[KapalRepoMysql.InsertData] Error when begin transaction : %w", err)
	}

	_, err = tx.Exec(qry, kapal.Kode, kapal.Muatan, kapal.Status, kapal.IsDelete)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[KapalRepoMysql.InsertData.Exec] Error when running query '"+qry+"' : %w", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[KapalRepoMysql.InsertData.Commit] Error when commiting '"+qry+"' : %w", err)
	}
	return nil
}

// UpdateData function call for usecase
func (k *KapalRepoMysql) UpdateData(id int, kapal *model.Kapal) error {
	qry := "update kapal set kode_kapal = ?, muatan= ?, status = ? where kapal_id = ?"

	tx, err := k.db.Begin()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[KapalRepoMysql.UpdateData] Error when begin transaction : %w", err)
	}

	_, err = tx.Exec(qry, kapal.Kode, kapal.Muatan, kapal.Status, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[KapalRepoMysql.UpdateData] Error when Exec '"+qry+"' : %w", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[KapalRepoMysql.UpdateData.Commit] Error when commiting '"+qry+"' : %w", err)
	}
	return nil
}
