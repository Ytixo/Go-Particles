package particles

import (
	"container/list"
	"math"
	"math/rand/v2"
	"project-particles/config"
	"project-particles/core"
	"project-particles/extension"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet.
func NewSystem() System {
	var posX, posY float64

	angle := (rand.Float64() * 360) * (math.Pi / 180)
	side := 10 * config.General.ScaleX
	marge := side * math.Sqrt2 // Marge : diagonale du carré, utilisée pour éviter qu’il dépasse de l’écran quelle que soit sa rotation.

	minX, minY := marge, marge
	maxX := float64(config.General.WindowSizeX) - marge
	maxY := float64(config.General.WindowSizeY) - marge

	l := list.New()

	for i := 0; i < config.General.InitNumParticles; i++ {

		if config.General.RandomSpawn {
			posX = minX + rand.Float64()*(maxX-minX)
			posY = minY + rand.Float64()*(maxY-minY)
		} else {
			posX = float64(config.General.SpawnX)
			posY = float64(config.General.SpawnY)
		}

		speedX := (rand.Float64()*config.General.SpeedFactor*2 - config.General.SpeedFactor)
		speedY := (rand.Float64()*config.General.SpeedFactor*2 - config.General.SpeedFactor)

		p := &core.Particle{
			ScaleX:     config.General.ScaleX,
			ScaleY:     config.General.ScaleY,
			PositionX:  posX,
			PositionY:  posY,
			Rotation:   angle,
			ColorRed:   1,
			ColorGreen: 1,
			ColorBlue:  1,
			Opacity:    1,
			SpeedX:     speedX,
			SpeedY:     speedY,
		}

		extension.ApplyFashion(p)

		l.PushFront(p)
	}

	return System{Content: l, SpawnAccumulator: 0}
}
