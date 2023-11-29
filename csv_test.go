package csv

import (
	"fmt"
	"testing"
)

func Test_NewWriter(t *testing.T) {
	c, err := NewWriter("test.csv", 10)
	if err != nil {
		t.Error(err)
	}
	if err := c.SetTitle("name", "age"); err != nil {
		t.Error(err)
	}
	for i := 0; i < 1000000; i++ {
		if err := c.PushRow("name", fmt.Sprintf("age:%d", i)); err != nil {
			t.Error(err)
		}
	}
	c.Flush()
	c.Close()
}
