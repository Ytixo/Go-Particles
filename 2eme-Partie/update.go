package main

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *game) Update() error {
	// Gestion des entrées clavier pour les différentes fonctionnalités
	if inpututil.IsKeyJustPressed(ebiten.KeyG) {
		config.General.Gravity = !config.General.Gravity
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		config.General.GravityStrength += 0.01
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		config.General.GravityStrength -= 0.01
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		config.General.Fashion = !config.General.Fashion
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyY) {
		config.General.ColorMode += 1
		if config.General.ColorMode > 3 {
			config.General.ColorMode = 0
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyF) {
		config.General.UseLife = !config.General.UseLife
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		config.General.LifeSpan += 10
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyV) {
		config.General.LifeSpan -= 10
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		config.General.OutOfScreen = !config.General.OutOfScreen
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		config.General.SpawnRate += 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		config.General.SpawnRate -= 1
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		config.General.RandomSpawn = !config.General.RandomSpawn
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		config.General.Modesouris = !config.General.Modesouris
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		config.General.CollisionMode = !config.General.CollisionMode
	}
	g.system.Update()
	return nil
}
