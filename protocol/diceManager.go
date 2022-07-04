package protocol

type GameStateRequest struct {
	Status int `json:"status"` // 1:stop, 2: start
}

type GameStateResponse struct {
	Msg string `json:"msg"`
}

type GuessGameRequest struct {
	Text string `json:"text"`
}
type GuessGameResponse struct {
	Text   string `json:"text"`
	Status bool   `json:"status"`
	Remark string `json:"remark"`
	Number int    `json:"number"`
}
