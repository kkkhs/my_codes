package model

type student struct {
	Name  string
	score float64
}

// 当结构体为私有，通过工厂模式解决
func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		score: s,
	}
}

// 属性为私有时：
func (this *student) GetScore() float64 {
	return this.score
}
