package extension

import (
	"project-particles/config"
	"project-particles/core"
)

// applique l'accélération de la gravité à la vitesse Y
func Gravity(p *core.Particle) {
	if config.General.Gravity {
		p.SpeedY += config.General.GravityStrength // accélération de la gravité
	}
}
