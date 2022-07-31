package engine

import (
	"github.com/kenoyer130/wartgame/models"
)

type Game struct {
	CurrentGameState models.GameState
	BattleGround     models.BattleGround
	CurrentPhase     models.Phase
	Players          [models.MaxPlayers]models.Player
	CurrentPlayer    models.Player
	Assets           models.Assets	
}
