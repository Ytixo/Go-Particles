package extension

import (
	"math"
	"project-particles/config"
	"project-particles/core"
	"testing"
)

func TestLifeSpan_NotExpired(t *testing.T) {
	config.General.LifeSpan = 10
	config.General.OpacitySpeed = 0.1

	p := &core.Particle{
		LifeCounter: 5,
		Opacity:     1.0,
	}
	// Teste une particule qui n'a pas encore expiré
	dead := LifeSpan(p)

	if dead {
		t.Errorf("La particule ne devrait pas être morte")
	}
	if p.LifeCounter != 6 {
		t.Errorf("LifeCounter attendu = 6, obtenu %d", p.LifeCounter)
	}
	if math.Abs(p.Opacity-0.9) > 1e-9 {
		t.Errorf("Opacity attendue = 0.9, obtenue %f", p.Opacity)
	}
}

func TestLifeSpan_Expired(t *testing.T) {
	config.General.LifeSpan = 3
	config.General.OpacitySpeed = 0.2

	p := &core.Particle{
		LifeCounter: 2,
		Opacity:     0.8,
	}
	// Teste une particule qui a expiré
	dead := LifeSpan(p)

	if !dead {
		t.Errorf("La particule devrait être morte")
	}
	if p.LifeCounter != 3 {
		t.Errorf("LifeCounter attendu = 3, obtenu %d", p.LifeCounter)
	}
	if math.Abs(p.Opacity-0.6) > 1e-9 {
		t.Errorf("Opacity attendue = 0.6, obtenue %f", p.Opacity)
	}
}

func TestLifeSpan_ZeroOpacity(t *testing.T) {
	config.General.LifeSpan = 5
	config.General.OpacitySpeed = 0.2
	// Teste une particule avec une opacité initiale de 0
	p := &core.Particle{
		LifeCounter: 0,
		Opacity:     0,
	}

	LifeSpan(p)

	if p.Opacity != 0 {
		t.Errorf("Opacity devrait rester à 0")
	}
}

// TestLifeSpan_ImmediateDeath teste une particule qui devrait mourir immédiatement
func TestLifeSpan_ImmediateDeath(t *testing.T) {
	config.General.LifeSpan = 1
	config.General.OpacitySpeed = 0.1

	p := &core.Particle{
		LifeCounter: 0,
		Opacity:     1.0,
	}

	dead := LifeSpan(p)

	if !dead {
		t.Errorf("La particule devrait mourir immédiatement")
	}
}
