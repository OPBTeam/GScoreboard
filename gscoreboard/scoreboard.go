package gscoreboard

import (
	"time"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
)

type Scoreboard struct {
	Title string
	Line  []string
	Delay int
}

func (s Scoreboard) Score(p *player.Player) {
	score := scoreboard.New(s.Title)
	score.RemovePadding()
	for i := 0; i < len(s.Line); i = i + 1 {
		score.Set(i, s.Line[i])
	}
	p.SendScoreboard(score)
}

func (s Scoreboard) Send(p *player.Player) {
	ticker := time.NewTicker(time.Second * time.Duration(s.Delay))
	for range ticker.C {
		if p == nil {
			ticker.Stop()
			return
		}
		s.Score(p)
	}
}
