package geometry

import "errors"

func CubeVolume(side float64) (float64, error) {
	if side != 0 {
		return side * side * side, nil
	} else {
		return 0, errors.New("side cannot be 0")
	}
}
