package extension

import (
	"testing"

	"project-particles/config"
	"project-particles/core"
)

// TestGravity_Enabled teste l'application de la gravité lorsque activée
func TestGravity_Enabled(t *testing.T) {
	config.General.Gravity = true
	config.General.GravityStrength = 0.98
	// Teste la gravité sur une particule avec une vitesse Y positive
	p := &core.Particle{
		SpeedY: 2.0,
	}

	Gravity(p)

	expected := 2.0 + 0.98
	if p.SpeedY != expected {
		t.Errorf("SpeedY incorrecte, attendu %f, obtenu %f", expected, p.SpeedY)
	}
}

// TestGravity_ZeroStrength teste que la gravité n'affecte pas la particule lorsque la force est zéro
func TestGravity_ZeroStrength(t *testing.T) {
	config.General.Gravity = true
	config.General.GravityStrength = 0.0

	p := &core.Particle{
		SpeedY: -1.5,
	}

	Gravity(p)

	if p.SpeedY != -1.5 {
		t.Errorf("SpeedY ne devrait pas changer avec GravityStrength = 0")
	}
}

// TestGravity_Disabled teste que la gravité n'affecte pas la particule lorsque désactivée
func TestGravity_Disabled(t *testing.T) {
	config.General.Gravity = false
	config.General.GravityStrength = 10.0

	p := &core.Particle{
		SpeedY: 3.0,
	}

	Gravity(p)

	if p.SpeedY != 3.0 {
		t.Errorf("SpeedY ne devrait pas être modifiée si Gravity = false")
	}
}

func TestGravity_NegativeSpeedY(t *testing.T) {
	config.General.Gravity = true
	config.General.GravityStrength = 1.0
	// Teste la gravité sur une particule avec une vitesse Y négative
	p := &core.Particle{
		SpeedY: -5.0,
	}

	Gravity(p)

	if p.SpeedY != -4.0 {
		t.Errorf("SpeedY incorrecte, attendu -4.0, obtenu %f", p.SpeedY)
	}
}
