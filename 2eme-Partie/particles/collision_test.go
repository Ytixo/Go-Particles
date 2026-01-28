package particles

import (
	"container/list"
	"math"
	"project-particles/config"
	"project-particles/core"
	"testing"
)

func newTestSystem(particles ...*core.Particle) *System {
	l := list.New()
	for _, p := range particles {
		l.PushBack(p)
	}
	return &System{Content: l}
}

//TEST 1 : collisions désactivées

func TestHandleCollisions_Disabled(t *testing.T) {
	config.General.CollisionMode = false

	p1 := &core.Particle{PositionX: 0, PositionY: 0, SpeedX: 1, SpeedY: 1}
	p2 := &core.Particle{PositionX: 0, PositionY: 0, SpeedX: -1, SpeedY: -1}

	sys := newTestSystem(p1, p2)
	sys.handleCollisions()

	if p1.SpeedX != 1 || p2.SpeedX != -1 {
		t.Errorf("les vitesses ne doivent pas changer si les collisions sont désactivées")
	}
}

// TEST 2 : collision détectée
func TestHandleCollisions_Collision(t *testing.T) {
	config.General.CollisionMode = true
	config.General.CollisionDistance = 10
	// Particules suffisamment proches pour une collision
	p1 := &core.Particle{
		PositionX: 0, PositionY: 0,
		SpeedX: 1, SpeedY: 0,
	}
	p2 := &core.Particle{
		PositionX: 5, PositionY: 0,
		SpeedX: -1, SpeedY: 0,
	}

	sys := newTestSystem(p1, p2)
	sys.handleCollisions()

	// vitesses échangées
	if p1.SpeedX != -1 || p2.SpeedX != 1 {
		t.Errorf("les vitesses n'ont pas été échangées correctement")
	}

	// distance finale >= distance minimale
	dx := p1.PositionX - p2.PositionX
	dy := p1.PositionY - p2.PositionY
	dist := math.Sqrt(dx*dx + dy*dy)

	if dist < config.General.CollisionDistance {
		t.Errorf("les particules ne sont pas assez séparées après la collision")
	}
}

//TEST 3 : pas de collision

func TestHandleCollisions_NoCollision(t *testing.T) {
	config.General.CollisionMode = true
	config.General.CollisionDistance = 5
	// Particules trop éloignées pour une collision
	p1 := &core.Particle{PositionX: 0, PositionY: 0, SpeedX: 1}
	p2 := &core.Particle{PositionX: 100, PositionY: 100, SpeedX: -1}

	sys := newTestSystem(p1, p2)
	sys.handleCollisions()

	if p1.SpeedX != 1 || p2.SpeedX != -1 {
		t.Errorf("les vitesses ne doivent pas changer s'il n'y a pas collision")
	}
}
