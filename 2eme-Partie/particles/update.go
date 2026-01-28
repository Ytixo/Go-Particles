package particles

import (
	"math"
	"math/rand/v2"
	"project-particles/config"
	"project-particles/core"
	"project-particles/extension"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps.
func (s *System) Update() {
	s.SpawnAccumulator += config.General.SpawnRate

	var posX, posY float64

	side := 10 * config.General.ScaleX
	marge := side * math.Sqrt2 // marge pour éviter que la particule dépasse de l'écran selon sa rotation
	angle := (rand.Float64() * 360) * (math.Pi / 180)

	minX, minY := marge, marge
	maxX := float64(config.General.WindowSizeX) - marge
	maxY := float64(config.General.WindowSizeY) - marge

	// Déclenché lorsqu'une particule complète a été générée.
	for s.SpawnAccumulator >= 1 {
		if config.General.RandomSpawn {
			posX = minX + rand.Float64()*(maxX-minX)
			posY = minY + rand.Float64()*(maxY-minY)
		} else if config.General.Modesouris {
			x, y := ebiten.CursorPosition()
			posX = float64(x)
			posY = float64(y)
		} else {
			posX = float64(config.General.SpawnX)
			posY = float64(config.General.SpawnY)
		}
		speedX := rand.Float64()*config.General.SpeedFactor*2 - config.General.SpeedFactor
		speedY := rand.Float64()*config.General.SpeedFactor*2 - config.General.SpeedFactor
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

		s.Content.PushFront(p)

		s.SpawnAccumulator -= 1
	}

	// Mise à jour des particules existantes
	for e := s.Content.Front(); e != nil; {
		p := e.Value.(*core.Particle)
		nextElement := e.Next() // On sauvegarde le suivant avant potentielle suppression
		// 2. Mouvement classique (toujours actif)
		p.PositionX += p.SpeedX
		p.PositionY += p.SpeedY

		extension.Gravity(p)

		// 3. Gestion de la mort (Extensions Sortie d'écran + Durée de vie)
		isDead := false

		// On vérifie la sortie d'écran
		if config.General.OutOfScreen {
			if extension.OutOfBounds(p) {
				isDead = true
			}
		}

		// On vérifie la durée de vie (si pas déjà mort)
		if config.General.UseLife && !isDead {
			isDead = extension.LifeSpan(p)
		}

		// Si la particule est morte, on la supprime
		if isDead {
			s.Content.Remove(e)
		}

		e = nextElement
	}

	s.handleCollisions()
}
