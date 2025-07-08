package main

type Card struct {
	ID       string `json:"_id"`
	SerialNo string `json:"serialNo"`
	LockerNo string `json:"lockerNo"`
	Rev      string `json:"_rev,omitempty"`
	Doctype  string `json:"doctype"`
}

type CardResponse struct {
	Status   string
	Body     []Card `json:"docs"`
	Bookmark string `json:"bookmark"`
	Warning  string `json:"warning"`
}

type Charge struct {
	ID         string `json:"_id"`
	Name       string `json:"name"`
	Card       Card   `json:"card"`
	DeviceType string `json:"deviceType"`
	Gender     string `json:"gender"`
	Adult      string `json:"adult"`
	Price      string `json:"price"`
	Collected  string `json:"collected"`
	Date       string `json:"date"`
	Time       string `json:"time"`
	Day        int    `json:"day"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Rev        string `json:"_rev,omitempty"`
	Doctype    string `json:"doctype"`
}

type ChargeResponse struct {
	Status   string
	Body     []Charge `json:"docs"`
	Bookmark string   `json:"bookmark"`
	Warning  string   `json:"warning"`
}
