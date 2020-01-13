package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"tugas/GoPORT/dermaga"
	"tugas/GoPORT/model"
)

// DermagaRepoMysql struct
type DermagaRepoMysql struct {
	db *sql.DB
}

// CreateDermagaRepoMysql function will call for func main() and send to usecase
func CreateDermagaRepoMysql(db *sql.DB) dermaga.DermagaRepo {
	return &DermagaRepoMysql{db}
}

// GetAllDermaga function
func (d *DermagaRepoMysql) GetAllDermaga() (*[]model.Dermaga, error) {

	qry := "select dock_id, kode_dock, status from dock where is_delete = 0"
	allDermaga := []model.Dermaga{}
	dermaga := model.Dermaga{}

	rows, err := d.db.Query(qry)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[DermagaRepoMysql.CreateDermagaRepoMysql.GetAllDermaga.QueryRow] Error when query data with : %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&dermaga.ID, &dermaga.Kode, &dermaga.Status); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, fmt.Errorf("[DermagaRepoMysql.CreateDermagaRepoMysql.GetAllDermaga.Scan] Error when query data with : %w", err)
		}

		allDermaga = append(allDermaga, dermaga)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("[DermagaRepoMysql.CreateDermagaRepoMysql.GetAllDermaga.rows.Err()] Error when rows.Err() data with : %w", err)
	}

	return &allDermaga, nil
}

// GetByID function
func (d *DermagaRepoMysql) GetByID(id int) (*model.Dermaga, error) {
	qry := "select dock_id, kode_dock, status from dock where dock_id = ? and is_delete = 0"

	dock := model.Dermaga{}

	err := d.db.QueryRow(qry, id).Scan(&dock.ID, &dock.Kode, &dock.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("[DermagaRepoMysql.CreateDermagaRepoMysqlImpl.GetByID.QueryRow] Error when running query '"+qry+"' : %w", err)
	}

	return &dock, nil
}

// InsertData calling for usecase
func (d *DermagaRepoMysql) InsertData(dock *model.Dermaga) error {
	query := "insert into dock (kode_dock, status,is_delete) values (?,?,?)"

	tx, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("[DermagaRepoMysql.]CreateDermagaRepoMysql.InsertData.Begin] Error when begin with : %w", err.Error())
	}

	_, err = tx.Exec(query, dock.Kode, dock.Status, dock.IsDelete)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[DermagaRepoMysql.]CreateDermagaRepoMysql.InsertData.Exec] Error when begin with : %w", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[KapalRepoMysql.InsertData.Commit] Error when commiting '"+query+"' : %w", err.Error())
	}
	return nil

}

// UpdateData function
func (d *DermagaRepoMysql) UpdateData(id int, dock *model.Dermaga) error {
	qry := "update dock set kode_dock = ?, status = ? where dock_id = ?"

	tx, err := d.db.Begin()
	if err != nil {
		return fmt.Errorf("[DermagaRepoMysql.]CreateDermagaRepoMysql.UpdateData.Begin] Error when begin with : %w", err.Error())
	}

	_, err = tx.Exec(qry, dock.Kode, dock.Status, id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[DermagaRepoMysql.]CreateDermagaRepoMysql.UpdateData.Exec] Error when begin with : %w", err.Error())
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("[KapalRepoMysql.InsertData.Commit] Error when commiting '"+qry+"' : %w", err.Error())
	}
	return nil
}

// DeleteByID function
func (d *DermagaRepoMysql) DeleteByID(id int) error {
	query := "update dock set is_delete = 1 where dock_id = ?"

	_, err := d.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("[KapalRepoMysql.DeleteByID.Exec] Error when Exec '"+query+"' : %w", err.Error())
	}
	return nil
}
