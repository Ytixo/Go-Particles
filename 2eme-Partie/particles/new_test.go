package particles_test

import (
	"math"
	"project-particles/config"
	"project-particles/core"
	"project-particles/particles"
	"testing"
)

func Test_NoRandomSpawn(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 100
	config.General.SpawnY = 200

	system := particles.NewSystem()

	if system.Content.Len() != 1 {
		t.Fatalf("on attendait 1 particule on a eu %d", system.Content.Len())
	}

	// Test si elles sont bien à la bonne position
	p := system.Content.Front().Value.(*core.Particle)

	if p.PositionX != 100 || p.PositionY != 200 {
		t.Errorf("La position devrait être : 100 et 200 alors que là elle est en : (%f, %f) ", p.PositionX, p.PositionY)
	}
}
func Test_RandomSpawn(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = true
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.ScaleX = 1.0 // S'assurer que le calcul de 'r' est cohérent

	system := particles.NewSystem()
	p := system.Content.Front().Value.(*core.Particle)

	// Calcul de la zone de spawn minimum/maximum (reproduit le calcul dans NewSystem)
	side := 10.0 * config.General.ScaleX
	r := side * math.Sqrt2
	minX := r
	minY := r
	maxX := float64(config.General.WindowSizeX) - r
	maxY := float64(config.General.WindowSizeY) - r

	if p.PositionX < minX || p.PositionX > maxX {
		t.Errorf("PositionX (%f) est en dehors de l'écran [%f, %f]", p.PositionX, minX, maxX)
	}
	if p.PositionY < minY || p.PositionY > maxY {
		t.Errorf("PositionY (%f) est en dehors de l'écran [%f, %f]", p.PositionY, minY, maxY)
	}
}

func Test_InitialSpeeds(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.SpeedFactor = 5.0

	system := particles.NewSystem()
	p := system.Content.Front().Value.(*core.Particle)

	// La plage de vitesse est [-5, 5] en x et y
	if p.SpeedX < -5 || p.SpeedX > 5 {
		t.Errorf("SpeedX (%f) est en dehors de ce qui était attendu [-5, 5]", p.SpeedX)
	}
	if p.SpeedY < -5 || p.SpeedY > 5 {
		t.Errorf("SpeedY (%f) est en dehors de ce qui était attendu [-5, 5]", p.SpeedY)
	}
}
