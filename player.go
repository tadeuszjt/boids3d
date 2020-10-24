package main

import (
	"github.com/tadeuszjt/geom/32"
    //"fmt"
)

const (
	playerSpeed           = 2
    playerLookSensitivity = 0.015
)

var (
	player = struct {
		position geom.Vec3
		bearing  geom.Angle
		pitch    geom.Angle
	}{
		position: geom.Vec3{0, 0, 3},
	}

	keys struct{ w, a, s, d bool }
)

func playerUpdate() {
    forward := geom.Vec3NormPitchYaw(player.pitch, player.bearing)
    right   := geom.Vec3NormPitchYaw(0, player.bearing.Plus(geom.Angle90Deg))

	if keys.w {
        player.position.PlusEquals(forward.ScaledBy(playerSpeed))
	}
	if keys.a {
        player.position.MinusEquals(right.ScaledBy(playerSpeed))
	}
	if keys.s {
        player.position.MinusEquals(forward.ScaledBy(playerSpeed))
	}
	if keys.d {
        player.position.PlusEquals(right.ScaledBy(playerSpeed))
	}
}

func playerLook(dx, dy float32) {
    player.bearing.PlusEquals(geom.MakeAngle(dx * playerLookSensitivity))
    player.pitch.PlusEquals(geom.MakeAngle(dy * playerLookSensitivity))
    player.pitch.Clamp(-geom.Angle90Deg, geom.Angle90Deg)
}
