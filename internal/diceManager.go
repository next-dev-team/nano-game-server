package internal

import (
	"fmt"
	"log"
	"nano-game-server/protocol"
	"time"

	"github.com/lonng/nano/component"
	"github.com/lonng/nano/scheduler"
	"github.com/lonng/nano/session"
)

type DiceManager struct {
	component.Base
	timer *scheduler.Timer
}

func NewDiceManager() *DiceManager {
	return &DiceManager{}
}

// do game state handler (start/stop the game)
func (m *DiceManager) DoGameState(s *session.Session, msg *protocol.GameStateRequest) error {
	log.Println("msg:", msg)
	id := s.ID()
	s.Bind(id)
	if msg.Status == 1 {
		stopRollingDice()
		return s.Response(protocol.GameStateResponse{
			Msg: "Game stopped",
		})
	}
	rollDice()
	return s.Response(protocol.GameStateResponse{
		Msg: "Game started",
	})
}

func (m *DiceManager) GuessGame(s *session.Session, msg *protocol.GuessGameRequest) error {
	status, text, remark, pickDice := guessDice(msg.Text)
	return s.Response(protocol.GuessGameResponse{
		Status: status,
		Text:   text,
		Remark: remark,
		Number: pickDice,
	})
}

// AfterInit component lifetime callback
func (mgr *DiceManager) AfterInit() {
	session.Lifetime.OnClosed(func(s *session.Session) {
		fmt.Println("closing connection here...")
	})
	mgr.timer = scheduler.NewTimer(time.Minute, func() {
		fmt.Println("dice manager trigger every minute...")

	})
}
