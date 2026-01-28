package extension

import (
	"math"
	"testing"

	"project-particles/config"
	"project-particles/core"
)

func TestApplyFashion_ColorMode0(t *testing.T) {
	config.General.Fashion = true
	config.General.ColorMode = 0

	p := &core.Particle{}
	ApplyFashion(p)
	// Vérifications des plages de valeurs
	if p.Rotation < 0 || p.Rotation > 2*math.Pi {
		t.Errorf("Rotation hors limites: %f", p.Rotation)
	}
	if p.ColorRed != 1.0 {
		t.Errorf("ColorRed attendu = 1.0, obtenu %f", p.ColorRed)
	}
	if p.ColorGreen < 0.3 || p.ColorGreen > 0.8 {
		t.Errorf("ColorGreen hors limites: %f", p.ColorGreen)
	}
	if p.ColorBlue < 0.0 || p.ColorBlue > 0.1 {
		t.Errorf("ColorBlue hors limites: %f", p.ColorBlue)
	}
	if p.Opacity < 0.6 || p.Opacity > 1.0 {
		t.Errorf("Opacity hors limites: %f", p.Opacity)
	}
}

func TestApplyFashion_ColorMode1(t *testing.T) {
	config.General.Fashion = true
	config.General.ColorMode = 1

	p := &core.Particle{}
	ApplyFashion(p)

	if p.ColorRed < 0.0 || p.ColorRed > 0.1 {
		t.Errorf("ColorRed hors limites: %f", p.ColorRed)
	}
	if p.ColorGreen < 0.3 || p.ColorGreen > 0.8 {
		t.Errorf("ColorGreen hors limites: %f", p.ColorGreen)
	}
	if p.ColorBlue < 0.7 || p.ColorBlue > 1.0 {
		t.Errorf("ColorBlue hors limites: %f", p.ColorBlue)
	}
}

func TestApplyFashion_ColorMode2(t *testing.T) {
	config.General.Fashion = true
	config.General.ColorMode = 2

	p := &core.Particle{}
	ApplyFashion(p)

	if p.ColorRed < 0.1 || p.ColorRed > 0.5 {
		t.Errorf("ColorRed hors limites: %f", p.ColorRed)
	}
	if p.ColorGreen < 0.5 || p.ColorGreen > 1.0 {
		t.Errorf("ColorGreen hors limites: %f", p.ColorGreen)
	}
	if p.ColorBlue < 0.1 || p.ColorBlue > 0.5 {
		t.Errorf("ColorBlue hors limites: %f", p.ColorBlue)
	}
}

func TestApplyFashion_ColorMode3(t *testing.T) {
	config.General.Fashion = true
	config.General.ColorMode = 3

	p := &core.Particle{}
	ApplyFashion(p)

	if p.ColorRed < 0.0 || p.ColorRed > 1.0 {
		t.Errorf("ColorRed hors limites: %f", p.ColorRed)
	}
	if p.ColorGreen < 0.0 || p.ColorGreen > 1.0 {
		t.Errorf("ColorGreen hors limites: %f", p.ColorGreen)
	}
	if p.ColorBlue < 0.0 || p.ColorBlue > 1.0 {
		t.Errorf("ColorBlue hors limites: %f", p.ColorBlue)
	}
}

// TestApplyFashion_Disabled teste que les valeurs ne changent pas lorsque Fashion est désactivé
func TestApplyFashion_Disabled(t *testing.T) {
	config.General.Fashion = false
	config.General.ColorMode = 3
	p := &core.Particle{
		ColorRed:   0.5,
		ColorGreen: 0.5,
		ColorBlue:  0.5,
		Opacity:    0.5,
		Rotation:   1.0,
	}

	ApplyFashion(p)
	// Vérifie que les valeurs n'ont pas changé
	if p.ColorRed != 0.5 || p.Rotation != 1.0 {
		t.Errorf("Les valeurs ne devraient pas être modifiées si Fashion=false")
	}
}
