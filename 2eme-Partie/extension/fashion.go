package extension

import (
	"math"
	"math/rand/v2"
	"project-particles/config"
	"project-particles/core"
)

// ApplyFashion applique un style visuel aux particules en fonction des paramètres
func ApplyFashion(p *core.Particle) {
	if config.General.Fashion {
		if config.General.ColorMode == 0 {
			p.Rotation = rand.Float64() * 2 * math.Pi
			p.ColorRed = 1.0
			p.ColorGreen = 0.3 + rand.Float64()*0.5
			p.ColorBlue = rand.Float64() * 0.1
			p.Opacity = 0.6 + rand.Float64()*0.4
		} else if config.General.ColorMode == 1 {
			p.Rotation = rand.Float64() * 2 * math.Pi
			p.ColorRed = rand.Float64() * 0.1
			p.ColorGreen = 0.3 + rand.Float64()*0.5
			p.ColorBlue = 0.7 + rand.Float64()*0.3
			p.Opacity = 0.6 + rand.Float64()*0.4
		} else if config.General.ColorMode == 2 {
			p.Rotation = rand.Float64() * 2 * math.Pi
			p.ColorRed = 0.1 + rand.Float64()*0.4
			p.ColorGreen = 0.5 + rand.Float64()*0.5
			p.ColorBlue = 0.1 + rand.Float64()*0.4
			p.Opacity = 0.6 + rand.Float64()*0.4
		} else if config.General.ColorMode == 3 {
			p.Rotation = rand.Float64() * 2 * math.Pi
			p.ColorRed = rand.Float64()
			p.ColorGreen = rand.Float64()
			p.ColorBlue = rand.Float64()
			p.Opacity = 0.6 + rand.Float64()*0.4
		}
	}
}
