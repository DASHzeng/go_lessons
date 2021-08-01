package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

func findRow(key string, tableName string) error {
	//基础函数，表中找不到目标行，直接返回原始的err，不包装信息
	err := sql.ErrNoRows
	return err
}

func TestFunction() error {
	key := "1010"
	tableName := "t_sys_user"
	//业务逻辑中发生ErrNoRows错误，将操作信息添加到err上返回,其他错误或未出错则另做处理
	if err := findRow(key, tableName); err == sql.ErrNoRows {
		return errors.Wrap(err, "target "+key+" not found in"+tableName)
	} else {
		return err
	}
}
