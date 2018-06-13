package pattern

import "fmt"

//抽象工厂
type AbstractFactory interface {
	Produce() pen //生产笔
}

//笔
type pen interface {
	//写字
	Write()
}

type PencilFactory struct {
}

type BrushPen struct {
}
//工厂
type PenFactory struct {
}


type pencil struct {
}

type brushPen struct {
}

func (PencilFactory) Produce() pen {
	return new(pencil)
}

func (BrushPen) Produce() pen {
	return new(brushPen)
}

func (p *pencil) Write() {
	fmt.Println("铅笔")
}

func (p *brushPen) Write() {
	fmt.Println("毛笔")
}

func (this PenFactory) Produce(typ string) pen {
	switch typ {
	case "pencil":
		return this.ProducePencil()
	case "brush":
		return this.ProduceBrushPen()
	default:
		return nil
	}
}

func (PenFactory) ProducePencil() pen {
	return new(pencil)
}

func (PenFactory) ProduceBrushPen() pen {
	return new(brushPen)
}
