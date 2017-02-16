package data

//Amount that coming from request: optinal
type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

// CardDetail include cardNum and pin
type CardDetail struct {
	CardNum  string   `json:"cardNum"`
	CardType CardType `json:"cardType"`
	PinNum   string   `json:"pinNum"`
	//CardExpiryDate time.Time `json:"cardExpiryDate"`
	Amount Amount `json:"amount"`
}

// CardType Const
type CardType string

// Card constats
const (
	GIFTCARD CardType = "giftcard"
	DEFAULT  CardType = "giftcard"
)

type error struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// Request object
type Request struct {
	Email      string     `json:"email"`
	Brand      string     `json:"brand"`
	CardDetail CardDetail `json:"cardDetails"`
}
