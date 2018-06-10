package pattern


type Books struct {
	title string
	author string
	subject string
	book_id int
}

type Student struct {
	Id string
	Name string
	age int
}



//抽象工厂
//add for test
type AbstractFactory interface {
	Produce() pen //生产笔
}

type PencilFactory struct {
}

func (PencilFactory) Produce() pen {
	return new(pencil)
}

type BrushPen struct {
}

func (BrushPen) Produce() pen {
	return new(brushPen)
}
