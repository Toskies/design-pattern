package builder

/*
separate the construction of a complex object from representation
设计思想：

	构建器(builder)模式：
		将一个复杂类的表示与其构造相分离，使得相同的构建过程能够得出不同的表示.它通过分步组装对象的方式，解决对象构造过程中参数过多、逻辑复杂或需要灵活组合的问题

	构建器模式包含以下核心角色：

		​​产品（Product）​​:最终要构建的复杂对象（如一辆汽车、一份订单 Vehicle）。

		​​抽象构建者（Builder）​​:定义构建产品的通用接口（如 buildPartA()、buildPartB()），隐藏具体实现, 对应具体的函数声明。

		​​具体构建者（Concrete Builder）​​:实现抽象构建者的接口，完成产品各部分的具体构造逻辑，对应具体的函数实现。

		​​指挥者（Director）​​:控制构建过程（可选），定义构建步骤的顺序（如先装引擎，再装车身），  对应Construct()。

	*Builder interface (包含1.biuld_XX method 返回的是biulder接口，2.get_XX 返回对象)
	*父struct
	*Director struct, 属性为 Builder, 实现Construct()和SetBuilder()方法
	*不同的子struct组合，实现接口builder
*/
type Wheel int

// 定义父struct
type Vehicle struct {
	Wheels    Wheel
	Seats     int
	Structure string
}

// builder interface
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	GetVehicle() Vehicle
}

// Director
type Director struct {
	builder Builder
}

func (director *Director) Construct() {
	director.builder.SetWheels().SetSeats().SetStructure() //链式调用
}

func (director *Director) SetBuilder(builder Builder) {
	director.builder = builder
}

// car struct
type Car struct {
	vehicle Vehicle
}

// 实现继承Builder
func (car *Car) SetWheels() Builder {
	car.vehicle.Wheels = 4
	return car
}
func (car *Car) SetSeats() Builder {
	car.vehicle.Seats = 4
	return car
}
func (car *Car) SetStructure() Builder {
	car.vehicle.Structure = "Car"
	return car
}
func (car *Car) GetVehicle() Vehicle {
	return car.vehicle
}
