package extension

import (
	"testing"

	"project-particles/config"
	"project-particles/core"
)

func TestOutOfBounds_Inside(t *testing.T) {
	config.General.Margin = 10
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	p := &core.Particle{
		PositionX: 400,
		PositionY: 300,
	}

	if OutOfBounds(p) {
		t.Errorf("La particule est à l'intérieur, elle ne devrait pas être tuée")
	}
}

func TestOutOfBounds_Left(t *testing.T) {
	config.General.Margin = 10
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	// Particule sortant par la gauche
	p := &core.Particle{
		PositionX: -11,
		PositionY: 300,
	}

	if !OutOfBounds(p) {
		t.Errorf("La particule est sortie à gauche, elle devrait être tuée")
	}
}

func TestOutOfBounds_Right(t *testing.T) {
	config.General.Margin = 10
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	// Particule sortant par la droite
	p := &core.Particle{
		PositionX: 811,
		PositionY: 300,
	}

	if !OutOfBounds(p) {
		t.Errorf("La particule est sortie à droite, elle devrait être tuée")
	}
}

func TestOutOfBounds_Top(t *testing.T) {
	config.General.Margin = 10
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	// Particule sortant par le haut
	p := &core.Particle{
		PositionX: 400,
		PositionY: -11,
	}

	if !OutOfBounds(p) {
		t.Errorf("La particule est sortie par le haut, elle devrait être tuée")
	}
}

func TestOutOfBounds_Bottom(t *testing.T) {
	config.General.Margin = 10
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	// Particule sortant par le bas
	p := &core.Particle{
		PositionX: 400,
		PositionY: 611,
	}

	if !OutOfBounds(p) {
		t.Errorf("La particule est sortie par le bas, elle devrait être tuée")
	}
}
