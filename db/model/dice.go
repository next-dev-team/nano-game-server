package model

// Todo -
type Dice struct {
	ID      int    `json:"id" gorm:"primary_key,auto_increment"`
	Release int    `json:"release"`
	Remark  string `json:"remark"`
	Guess   string `json:"guess"`

	// -----------
	BaseModel
}
