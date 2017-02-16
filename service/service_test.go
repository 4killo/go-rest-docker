package service

import "testing"
import "github.com/4killo/go-rest-docker/data"

// TestInsert is testing data inserting.
func TestInsertData(t *testing.T) {
	t.Log("Testing All MySql operation... (expected score: 0)")
	var r data.Request
	r.Brand = "AMAZON"
	r.Email = "yyyy@g.com"
	r.CardDetail.CardNum = "12345678"
	r.CardDetail.CardType = data.DEFAULT
	r.CardDetail.PinNum = "123"
	r.CardDetail.Amount.Currency = "USD"

	if res, _ := Insert(&r); res != 200 {
		t.Error("Insert is failed! response code = ", res)
	}

	// if res, _ := service.Insert(&r); res == 200 {
	// 	t.Error("Insert is Duplicate! response code = ", res)
	// }
}
