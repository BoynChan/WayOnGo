package main

/*
模拟将物品放入背包中
*/

type Bag struct {
	items []int
}

//在这个函数中 (b *Bag)就是接收器
//它作为方法作用的目标,可以使得作为结构体的对象方法,进行调用
func (b *Bag) Insert(itemId int) {
	b.items = append(b.items, itemId)
}

//定义了一个Property类,演示get/set方法的使用
type Property struct {
	value int
}

func (p *Property) SetValue(v int) {
	p.value = v
}

func (p *Property) Value() int {
	return p.value
}

func main() {
	b := new(Bag)
	b.Insert(5)
	b.Insert(1)
}
