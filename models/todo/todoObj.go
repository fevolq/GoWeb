package todo

import (
	"fmt"
	"go-web/dao"
	"log"
	"strings"
)

type Todo struct {
	ID     int
	Title  string
	Memo   string
	Status bool
}

const table = "todo"

func (td *Todo) Init(tdID int) error {
	td.ID = tdID
	sql := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table)
	err := dao.MysqlDb.QueryRow(sql, tdID).Scan(&td.ID, &td.Title, &td.Memo, &td.Status)
	if err != nil {
		log.Println(err)
	}
	return err
}

func (td *Todo) Add(title, memo string) (interface{}, error) {
	td.Title = title
	td.Memo = memo
	td.Status = true
	log.Println(*td)

	sql := fmt.Sprintf("INSERT INTO %s (title, memo, status) VALUE(?, ?, ?)", table)
	res, err := dao.MysqlDb.Exec(sql, title, memo, false)
	return res, err
}

func (td *Todo) Edit(info map[string]interface{}) error {
	log.Println(*td)
	log.Println(info)

	items := []string{}
	args := []interface{}{}

	title, ok := info["title"]
	if ok {
		items = append(items, "title = ?")
		args = append(args, title)
	}
	memo, ok := info["memo"]
	if ok {
		items = append(items, "memo = ?")
		args = append(args, memo)
	}
	status, ok := info["status"]
	if ok {
		items = append(items, "status = ?")
		args = append(args, status)
	}

	sql := fmt.Sprintf("UPDATE %s SET %s WHERE id = ?", table,
		strings.Join(items, " AND "))
	args = append(args, td.ID)
	log.Println(sql, args)

	res, err := dao.MysqlDb.Exec(sql, args...)
	if err != nil {
		log.Println("更新失败", err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
