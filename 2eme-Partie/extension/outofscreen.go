package extension

import (
	"project-particles/config"
	"project-particles/core"
)

// OutOfBounds vérifie si la particule est sortie de la zone autorisée
// Retourne true si la particule doit être tuée
func OutOfBounds(p *core.Particle) bool {
	m := config.General.Margin
	wX := float64(config.General.WindowSizeX)
	wY := float64(config.General.WindowSizeY)

	if p.PositionX < -m || p.PositionX > wX+m || p.PositionY < -m || p.PositionY > wY+m {
		return true
	}
	return false
}
