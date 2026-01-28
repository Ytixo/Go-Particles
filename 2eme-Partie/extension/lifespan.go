package extension

import (
	"project-particles/config"
	"project-particles/core"
)

// LifeSpan gère la vie de la particule et son opacité
// Retourne true si la particule doit mourir de vieillesse
func LifeSpan(p *core.Particle) bool {
	p.LifeCounter++

	if p.Opacity > 0 {
		p.Opacity -= config.General.OpacitySpeed
		if p.Opacity < 0 {
			p.Opacity = 0
		}
	}

	return p.LifeCounter >= config.General.LifeSpan
}
