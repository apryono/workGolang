package model

// Detail for bill
type Detail struct {
	Menu  string `json:"Nama_menu"`
	Qty   int    `json:"Qty"`
	Price int    `json:"Harga"`
	Total int    `json:"Total"`
}

// Transaction for Data Bill
type Transaction struct {
	TransID  int      `json:"TransactionID"`
	MejaID   int      `json:"Nomor_meja"`
	StatusTR string   `json:"Status Transaksi"`
	GranTot  int      `json:"Grand_total,omitempty"`
	Detail   []Detail `json:"detail,omitempty"`
}
