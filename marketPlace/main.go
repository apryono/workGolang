package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type item struct {
	itemID   int
	itemName string
	qty      int
	trxID    int
	price    int
	total    int
	grandTot int
}

type trx struct {
	ID          int
	trxDate     string
	cashierName string
}

type month struct {
	month      int
	jlhTrx     int
	grandTotal int
	jlhQty     int
}

var trs []trx
var brg []item

var getNumb int
var getInput, name string
var sym2 = strings.Repeat("-", 68)
var sym = strings.Repeat("-", 22)
var sym1 = strings.Repeat("-", 15)

func main() {
	fmt.Println("Karakter Lebih dari 2 dan Kurang dari 21 ")
	fmt.Print("Nama Kasir : ")
	name = input()
	getName := strings.ContainsAny(name, "1234567890!@#$%%^&*()_+-=~")
	if getName {
		fmt.Println("Sorry, Just Input Abjad")
		return
	}
	menu(name)
}

func menu(data string) {
	for {
		fmt.Println(sym2)
		fmt.Println(sym, "[+] + Main Menu + [+]", sym)
		fmt.Println(sym2)
		fmt.Println(`
	Silahkan Pilih Menu 
	1. Input Transaction
	2. View Report
	3. Exit
	`)
		fmt.Println(sym2)
		fmt.Print("Masukkan Menu yang dipilih : ")
		getInput = input()
		getNumb, _ = strconv.Atoi(getInput)

		switch getNumb {
		case 1:
			inputTrx(data)
		case 2:
			viewReport()
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Sorry, Input Sesuai Pilihan")
		}
	}
}

func input() string {
	getData := bufio.NewScanner(os.Stdin)
	getData.Scan()
	data := getData.Text()
	return data
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_toko")
	if err != nil {
		fmt.Println("Sorry Any Problem, Check Your Database Name !")
		return nil, err
	}

	return db, nil
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error, Please Check it Again")
		fmt.Println(err.Error())
	}
}

func inputTrx(data string) {
	db, err := connect()
	handleErr(err)

	defer db.Close()

	tx, err := db.Begin()
	handleErr(err)

	res, err := tx.Exec("insert into transactions (tr_date, cashier_name) values (current_date(),?)", data)

	if err != nil {
		tx.Rollback()
		fmt.Println(err.Error())
	}

	id, err := res.LastInsertId()
	handleErr(err)
	for {
		fmt.Print("Item Name : ")
		getInput = input()
		getName := strings.ContainsAny(getInput, "1234567890!@#$%%^&*()_+-=~")
		if !getName && len(getInput) > 2 && len(getInput) < 21 {
			for {
				fmt.Print("Quantity : ")
				getQ := input()
				getQty, _ := strconv.Atoi(getQ)
				if getQty > 0 {
					for {
						fmt.Print("Price : ")
						getP := input()
						getPrice, _ := strconv.Atoi(getP)
						if getPrice > 0 {
							res, err = tx.Exec("insert into items (item_name,qty,price,transaction_id) values (?,?,?,?)", getInput, getQty, getPrice, id)

							if err != nil {
								tx.Rollback()
								fmt.Println(err.Error())
							}

							fmt.Println("Do you want to add again ? yes, no")
							question := strings.ToLower(input())
							if question == "yes" {
								fmt.Println("Silahkan input kembali")
							}
							if question == "no" {
								var sumTotal int
								sumTotal += getQty * getPrice
								fmt.Println("Total : ", getQty*getPrice)
								fmt.Println("Grand Total : ", sumTotal)
								err = tx.Commit()
								if err != nil {
									fmt.Println("Something Problem, Please Check Your History")
								}

								log.Println("Done.")
								break
							}

						}
						break
					}
				}
				break
			}
		}
	}
}

func viewReport() {
	for {
		fmt.Println(sym2)
		fmt.Println(sym, "[+] + Sub Menu + [+]", sym)
		fmt.Println(sym2)
		fmt.Println(`
	Silahkan Pilih Menu 
	1. View All Transactions
	2. Transaction Detail
	3. Monthly Report
	4. Exit
	`)
		fmt.Println(sym2)
		fmt.Print("Masukkan Menu yang dipilih : ")
		getInput = input()
		getNumb, _ = strconv.Atoi(getInput)

		switch getNumb {
		case 1:
			viewAllTrx()
		case 2:
			detailTrx()
		case 3:
			monthlyReport()
		case 4:
			os.Exit(0)
		default:
			fmt.Println("Sorry, Input Sesuai Pilihan")
		}
	}
}

func viewAllTrx() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query(`select tr_id from transactions as t join items on t.tr_id = items.transaction_id`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var each = trx{}
		var err = rows.Scan(&each.ID)

		if err != nil {
			fmt.Println("There's Nothing ID")
			return
		}

		trs = append(trs, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Data Transaction Number : ")
	for _, each := range trs {
		fmt.Printf("%d\n", each.ID)
	}
}

func detailTrx() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var trs = trx{}
	var brg = item{}

	for {
		fmt.Print("Input ID Transaction Number : ")
		var getTr, er = strconv.Atoi(input())
		if er != nil {
			fmt.Println("Mohon Input ID Transaction yang Benar!")
		}
		err = db.
			QueryRow(`select tr_id, cashier_name, tr_date, qty, sum(qty * price), item_name, (qty * price), price
		from transactions as t join items on t.tr_id = items.transaction_id where tr_id = ?`, getTr).
			Scan(&trs.ID, &trs.cashierName, &trs.trxDate, &brg.qty, &brg.grandTot, &brg.itemName, &brg.total, &brg.price)

		if err != nil {
			fmt.Println("Data yang diinput tidak sesuai dengan data transaksi")
			return
		}

		fmt.Println(sym2)
		fmt.Println(sym1, "[+] + Transaction Detail + [+]", sym1)
		fmt.Println(sym2)
		fmt.Printf("Transaction Number\t: %d\nCashier Name\t\t: %s\nTransaction Date\t: %s\nGrand Total\t\t: %d\n", trs.ID, trs.cashierName, trs.trxDate, brg.grandTot)
		fmt.Println(sym2)
		fmt.Printf("Nama Barang |\tQty |\tPrice |\tTotal\n")
		fmt.Printf("\t%s\t%d\t%d\t%d\n", brg.itemName, brg.qty, brg.price, brg.total)
		break
	}
}

func monthlyReport() {
	db, err := connect()
	if err != nil {
		fmt.Println("Error.Check your Database")
	}
	defer db.Close()

	rows, err := db.Query("select extract(month from tr_date) as month, count(tr_id) as jlhTrx, sum(qty*price) as GrandTotal, sum(qty) as jlhQty from transactions join items on items.transaction_id = transactions.tr_id group by month")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []month

	for rows.Next() {
		var each = month{}
		var err = rows.Scan(&each.month, &each.jlhTrx, &each.grandTotal, &each.jlhQty)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Month\tJlhTrx\tGrandTot\tJlhQty\n")
	for _, each := range result {
		fmt.Printf("%d\t%d\t%d\t\t%d\n", each.month, each.jlhTrx, each.grandTotal, each.jlhQty)
	}

}
