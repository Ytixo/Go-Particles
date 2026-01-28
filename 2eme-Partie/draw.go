package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/core"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*core.Particle)
		if ok {
			if p.Opacity <= 0 {
				continue
			}

			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}
	// Tous les status et infos pour les touches de contrôle
	if config.General.Debug {
		gravityStatus := "désactivée"
		if config.General.Gravity {
			gravityStatus = "activée"
		}
		colorStatus := "désactivée"
		if config.General.Fashion {
			colorStatus = "activée"
		}
		lifeStatus := "désactivée"
		if config.General.UseLife {
			lifeStatus = "activée"
		}
		oosStatus := "désactivée"
		if config.General.OutOfScreen {
			oosStatus = "activée"
		}
		randomStatus := "désactivée"
		if config.General.RandomSpawn {
			randomStatus = "activée"
		}
		modesourisStatus := "désactivée"
		if config.General.Modesouris {
			modesourisStatus = "activée"
		}
		colormodeStatus := "désactivée"
		if config.General.ColorMode == 0 {
			colormodeStatus = "Fire"
		} else if config.General.ColorMode == 1 {
			colormodeStatus = "Ocean"
		} else if config.General.ColorMode == 2 {
			colormodeStatus = "Forest"
		} else if config.General.ColorMode == 3 {
			colormodeStatus = "Chromatic"
		}
		collisionmodeStatus := "désactivée"
		if config.General.CollisionMode {
			collisionmodeStatus = "activée"
		}
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f\nParticles: %d\nGravité: %s Force : %.2f ('G' ON/OFF 'T' + 'B' -) \nColor : %s -> %s ('H' ON/OFF 'Y' Changement de couleur) \nDurée de vie : %d et %s ('F' ON/OFF 'R' + 'V' -) \nSortie d'écran : %s ('J' ON/OFF) \nSpawnRate : %.0f ('E' + 'C' -) \nSpawn Random : %s ('L' ON/OFF) \nMode Souris : %s ('S' ON/OFF) \nCollision : %s ('A' ON/OFF)", ebiten.CurrentTPS(), g.system.Content.Len(), gravityStatus, config.General.GravityStrength, colorStatus, colormodeStatus, config.General.LifeSpan, lifeStatus, oosStatus, config.General.SpawnRate, randomStatus, modesourisStatus, collisionmodeStatus))
	}

}
