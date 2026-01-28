package particles_test

import (
	"math"
	"project-particles/config"
	"project-particles/core"
	"project-particles/particles"
	"testing"
)

func TestUpdate_SpeedApplication(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.SpeedFactor = 3.0

	system := particles.NewSystem()
	p := system.Content.Front().Value.(*core.Particle)

	initialX, initialY := p.PositionX, p.PositionY
	speedX, speedY := p.SpeedX, p.SpeedY

	system.Update()

	expectedX := initialX + speedX
	expectedY := initialY + speedY

	if p.PositionX != expectedX {
		t.Errorf("PositionX attendu : %f obtenu : %f", expectedX, p.PositionX)
	}
	if p.PositionY != expectedY {
		t.Errorf("PositionY attendu : %f obtenu : %f", expectedY, p.PositionY)
	}
}

func TestUpdate_SpawnAccumulator(t *testing.T) {
	// Cas 1 : SpawnRate < 1 → aucune nouvelle particule
	config.General.InitNumParticles = 0
	config.General.SpawnRate = 0.5

	system := particles.NewSystem()
	initialLen := system.Content.Len()

	system.Update() // Accumulateur : 0 → 0.5

	if system.Content.Len() != initialLen {
		t.Errorf("SpawnRate 0.5 : attendu %d particules, obtenu %d", initialLen, system.Content.Len())
	}
	if system.SpawnAccumulator != 0.5 {
		t.Errorf("SpawnRate 0.5 : accumulateur attendu 0.5, obtenu %f", system.SpawnAccumulator)
	}

	// Cas 2 : Accumulateur atteint 1 → une particule générée
	system.Update() // Accumulateur : 0.5 → 1 → spawn → 0

	if system.Content.Len() != initialLen+1 {
		t.Errorf("SpawnRate 0.5 x2 : attendu %d particules, obtenu %d", initialLen+1, system.Content.Len())
	}
	if system.SpawnAccumulator != 0 {
		t.Errorf("SpawnRate 0.5 x2 : accumulateur attendu 0.0, obtenu %f", system.SpawnAccumulator)
	}

	// Cas 3 : SpawnRate > 1 → plusieurs particules générées
	config.General.SpawnRate = 2.7
	system2 := particles.NewSystem()

	system2.Update() // Accumulateur : 0 → 2.7 → spawn 2 particules → reste 0.7

	if system2.Content.Len() != 2 {
		t.Errorf("SpawnRate 2.7 : attendu 2 particules, obtenu %d", system2.Content.Len())
	}

	if math.Abs(system2.SpawnAccumulator-0.7) > 1e-6 {
		t.Errorf("SpawnRate 2.7 : accumulateur attendu 0.7, obtenu %f", system2.SpawnAccumulator)
	}
}
