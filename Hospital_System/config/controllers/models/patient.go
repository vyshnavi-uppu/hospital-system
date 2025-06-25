package models

type Patient struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Gender  string `json:"gender"`
    Phone   string `json:"phone"`
    Address string `json:"address"`
}
