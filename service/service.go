package service

import (
	"fmt"
	"log"

	"github.com/4killo/go-rest-docker/data"
)

// Insert card
func Insert(r *data.Request) (int, error) {
	fmt.Println("request::insert", r)
	db := connect()
	defer db.Close()

	stm, err := db.Prepare("INSERT INTO gift_cards (id, brand, gift_card_number, pin, currency) VALUES (?,?,?,?,?)")
	if err != nil {
		log.Fatal("Error", err.Error())
		return 400, err
	}
	id, err := stm.Exec(r.Email, r.Brand, r.CardDetail.CardNum, r.CardDetail.PinNum, r.CardDetail.Amount.Currency)
	if err != nil {
		log.Print(err.Error())
		return 400, err
	}
	log.Println(id)

	return 200, nil
}

// GetCards to get card details
func GetCards(r *data.Request) []data.CardDetail {
	db := connect()
	defer db.Close()
	fmt.Println("Data = " + r.Email + r.Brand)
	results, err := db.Query("SELECT gift_card_number, pin, currency from gift_cards  where id = ? and brand = ?", r.Email, r.Brand)

	if err != nil {
		fmt.Println(err.Error())
	}

	var cardDetails []data.CardDetail

	for results.Next() {
		var cardDetail data.CardDetail
		err := results.Scan(&cardDetail.CardNum, &cardDetail.PinNum, &cardDetail.Amount.Currency)
		if err != nil {
			log.Print(err.Error())
		}
		fmt.Println(cardDetail.CardNum + " " + cardDetail.PinNum)
		cardDetail.CardType = data.DEFAULT
		cardDetails = append(cardDetails, cardDetail)
	}
	fmt.Println(cardDetails)
	return cardDetails
}
