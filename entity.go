package main

type Entity interface {
	Width() int
	Height() int
	X() float64
	Y() float64
}

type GameObject struct {
	width  int
	height int
	x      float64
	y      float64
}

func (gameObj *GameObject) Width() int {
	return gameObj.width
}

func (gameObj *GameObject) Height() int {
	return gameObj.height
}

func (gameObj *GameObject) X() float64 {
	return gameObj.x
}

func (gameObj *GameObject) Y() float64 {
	return gameObj.y
}

// CheckCollision 检查是否有碰撞
func CheckCollision(entityA, entityB Entity) bool {
	top, left := entityA.Y(), entityA.X()
	bottom, right := entityA.Y()+float64(entityA.Height()), entityA.X()+float64(entityA.Width())
	// 左上角
	x, y := entityB.X(), entityB.Y()
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 右上角
	x, y = entityB.X()+float64(entityB.Width()), entityB.Y()
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 左下角
	x, y = entityB.X(), entityB.Y()+float64(entityB.Height())
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	// 右下角
	x, y = entityB.X()+float64(entityB.Width()), entityB.Y()+float64(entityB.Height())
	if y > top && y < bottom && x > left && x < right {
		return true
	}

	return false
}
