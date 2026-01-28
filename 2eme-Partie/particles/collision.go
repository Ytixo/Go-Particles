package particles

import (
	"math"
	"project-particles/config"
	"project-particles/core"
)

func (s *System) handleCollisions() {
	// si les collisions sont désactivé on fait rien
	if !config.General.CollisionMode {
		return
	}

	distMin := config.General.CollisionDistance

	// on parcourt toute la liste avec deux boucles pour comparer chaque particule
	for e1 := s.Content.Front(); e1 != nil; e1 = e1.Next() {
		p1 := e1.Value.(*core.Particle)

		for e2 := e1.Next(); e2 != nil; e2 = e2.Next() {
			p2 := e2.Value.(*core.Particle)

			// calcule de la distance entre p1 et p2 (formule classique)
			dx := p1.PositionX - p2.PositionX
			dy := p1.PositionY - p2.PositionY
			distance := math.Sqrt(dx*dx + dy*dy)

			// si elles sont trop proches on considere une collision
			if distance < distMin {
				//  gestion de la collision

				// on échange les vitesses pour simuler un rebond simple
				p1.SpeedX, p2.SpeedX = p2.SpeedX, p1.SpeedX
				p1.SpeedY, p2.SpeedY = p2.SpeedY, p1.SpeedY

				// on les écarte un peu pour éviter qu'elles se superpose
				// sinon elles vont se collisionner en boucle
				overlap := (distMin - distance) / 2
				p1.PositionX += (dx / distance) * overlap
				p1.PositionY += (dy / distance) * overlap
				p2.PositionX -= (dx / distance) * overlap
				p2.PositionY -= (dy / distance) * overlap
			}
		}
	}
}
