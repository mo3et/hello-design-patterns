package main

import "fmt"

type Bike struct {
	frame string
	tires string
}

func (b *Bike) setFrame(frame string) {
	b.frame = frame
}

func (b *Bike) setTires(tires string) {
	b.tires = tires
}

func (b *Bike) String() string {
	return b.frame + " " + b.tires
}

// Bike builder interface
type BikeBuilder interface {
	buildFrame()
	buildTires()
	getResult() *Bike
}

type MountainBikeBuilder struct {
	bike *Bike
}

func NewMountainBikeBuilder() *MountainBikeBuilder {
	return &MountainBikeBuilder{
		bike: &Bike{},
	}
}

func (mbb *MountainBikeBuilder) buildFrame() {
	mbb.bike.setFrame("Aluminum Frame")
}

func (mbb *MountainBikeBuilder) buildTires() {
	mbb.bike.setTires("Knobby Tires")
}

func (mbb *MountainBikeBuilder) getResult() *Bike {
	return mbb.bike
}

// Road Bike Builder
type RoadBikeBuilder struct {
	bike *Bike
}

func NewRoadBikeBuilder() *RoadBikeBuilder {
	return &RoadBikeBuilder{
		bike: &Bike{},
	}
}

func (rbb *RoadBikeBuilder) buildFrame() {
	rbb.bike.setFrame("Carbon Frame")
}

func (rbb *RoadBikeBuilder) buildTires() {
	rbb.bike.setTires("Slim Tires")
}

func (rbb *RoadBikeBuilder) getResult() *Bike {
	return rbb.bike
}

// Bike Director, responsible for build bike
type BikeDirector struct{}

func (bd *BikeDirector) construct(builder BikeBuilder) *Bike {
	builder.buildFrame()
	builder.buildTires()
	return builder.getResult()
}

func main() {
	var N int
	fmt.Scan(&N) // 订单数量

	director := &BikeDirector{}

	for i := 0; i < N; i++ {
		var bikeType string
		fmt.Scan(&bikeType)

		var builder BikeBuilder
		// 根据输入类别，创建不同类型的具体建造者
		if bikeType == "mountain" {
			builder = NewMountainBikeBuilder()
		} else {
			builder = NewRoadBikeBuilder()
		}
		// Director负责指导生产产品
		bike := director.construct(builder)
		fmt.Println(bike)
	}
}
