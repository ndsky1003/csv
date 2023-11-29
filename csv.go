/*
非线程安全
本身代码就很简单
只是为了方便使用，导出好几百万的数据，避免全部存在内存当中，设置一个detect_row，每detect_row行就刷新一次
*/
package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
)

type CSV struct {
	wc         io.WriteCloser
	w          *csv.Writer
	FileName   string
	row        uint
	detect_row uint
}

func NewWriter(fileName string, detectRow uint) (*CSV, error) {
	finfo, err := os.Stat(fileName)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
	} else if finfo.IsDir() {
		return nil, errors.New("file is dir")
	} else {
		return nil, errors.New("file exist")
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	if detectRow == 0 {
		detectRow = 1000
	}
	return &CSV{
		wc:         f,
		w:          csv.NewWriter(f),
		FileName:   fileName,
		detect_row: detectRow,
	}, nil
}

func (this *CSV) Close() {
	if this.wc != nil {
		this.wc.Close()
	}
}

func (this *CSV) SetTitle(titles ...string) error {
	return this.w.Write(titles)
}

func (this *CSV) PushRow(datas ...string) error {
	err := this.w.Write(datas)
	this.row++
	if this.row%this.detect_row == 0 {
		this.w.Flush()
	}
	return err
}

func (this *CSV) Flush() {
	this.w.Flush()
}
