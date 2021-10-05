package main

import (
	//"debug/macho"
	"fmt"
	"image"
	"log"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex

type Message struct {
	Pos       image.Point
	LifeSigns int
	Rover     string
}

const (
	//火星日的时长
	dayLength = 24 * time.Second
	//在一天之中，探测器能够向地球发送消息的时长
	receiveTimePerDay = 2 * time.Second
)

//earthReceiver会接收从火星发来的消息
func earthReceiver(msgc chan []Message) {
	for {
		time.Sleep(dayLength - receiveTimePerDay)
		receiveMarsMessages(msgc)
	}
}

//接收火星消息
func receiveMarsMessages(msgc chan []Message) {
	finished := time.After(receiveTimePerDay)
	for {
		select {
		case <-finished:
			return
		case ms := <-msgc:
			for _, m := range ms {
				log.Printf("erath received report of life sign level %d from %s at %v", m.LifeSigns, m.Rover, m.Pos)

			}
		}
	}
}
func startDriver(name string, grid *MarsGrid, marsToEarth chan []Message) *RoverDriver {
	var o *Occupier
	for o == nil {
		startPoint := image.Point{X: rand.Intn(grid.Size().X), Y: rand.Intn(grid.Size().Y)}
		o = grid.Occupy(startPoint)

	}
	return NewRoverDriver(name, o, marsToEarth)
}

type command int

//驱动系统，驱动一台在火星表面行走的探测器
type RoverDriver struct {
	commandc chan command
	occupier *Occupier
	name     string
	radio    *Radio
}

//无线电发射器
type Radio struct {
	fromRover chan Message
}

//发送消息到地球,使用的是通道类型
func (r *Radio) SendToEarth(m Message) {
	r.fromRover <- m
}

//创建一个新的Radio实例
func NewRadio(toEarth chan []Message) *Radio {
	r := &Radio{
		fromRover: make(chan Message),
	}
	go r.run(toEarth) //开启并发
	return r
}

//将探测器发送的消息放进缓冲区，直到他们可以被发送至地球为止
func (r *Radio) run(toEarth chan []Message) {
	var buffered []Message
	for {
		toEarth1 := toEarth
		if len(buffered) == 0 {
			toEarth1 = nil
		}
		select {
		case m := <-r.fromRover:
			buffered = append(buffered, m)
		case toEarth1 <- buffered:
			buffered = nil
		}
	}
}

//NewRoverDriver启动一台新的RoverDriver
func NewRoverDriver(name string, occupier *Occupier, marsToEarth chan []Message) *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
		occupier: occupier,
		name:     name,
		radio:    NewRadio(marsToEarth),
	}
	go r.drive()
	return r
}

const (
	right command = 0
	left  command = 1
)

//驱动探测器
func (r *RoverDriver) drive() {
	log.Printf("%s initial position %v", r.name, r.occupier.Pos())

	direction := image.Point{X: 1, Y: 0}
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)
	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("%s new direction %v", r.name, direction) //会输出时间
		case <-nextMove:
			nextMove = time.After(updateInterval)
			newPos := r.occupier.Pos().Add(direction)
			if r.occupier.MoveTo(newPos) {
				log.Printf("%s move to %v", r.name, newPos)
				r.checkForLife()
				break
			}
			log.Printf("%s blocked trying to move from %v to %v", r.name, r.occupier.Pos(), newPos)
			//随机地选择其他一个方向
			//从下一轮迭代开始，探测器将朝着新的方向进发
			dir := rand.Intn(3) + 1
			for i := 0; i < dir; i++ {
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			}
			log.Printf("%s new random direction %v", r.name, direction)
		}
	}
}

func (r *RoverDriver) checkForLife() {
	sensorData := r.occupier.Sense()
	if sensorData.LifeSigns < 900 {
		return
	}
	r.radio.SendToEarth(Message{
		Pos:       r.occupier.Pos(),
		LifeSigns: sensorData.LifeSigns,
		Rover:     r.name,
	})
}
func (r *RoverDriver) Left() {
	r.commandc <- left
}
func (r *RoverDriver) Right() {
	r.commandc <- right
}

type MarsGrid struct {
	bounds image.Rectangle
	mu     sync.Mutex
	cells  [][]cell
}
type SensorData struct {
	LifeSigns int
}
type cell struct {
	groundData SensorData
	occupier   *Occupier
}

//创建一个火星grid地图
func NewMarsGrid(size image.Point) *MarsGrid {
	grid := &MarsGrid{
		bounds: image.Rectangle{
			Max: size,
		},
		cells: make([][]cell, size.Y),
	}
	for y := range grid.cells {
		grid.cells[y] = make([]cell, size.X)
		for x := range grid.cells[y] {
			cell := &grid.cells[y][x]
			cell.groundData.LifeSigns = rand.Intn(1000)
		}
	}
	return grid
}

//返回一个表示网格大小的Point
func (g *MarsGrid) Size() image.Point {
	return g.bounds.Max
}

//占领网格中给定坐标点上的单元格，
func (g *MarsGrid) Occupy(p image.Point) *Occupier {
	g.mu.Lock()
	defer g.mu.Unlock()
	cell := g.cell(p)
	if cell == nil || cell.occupier != nil {
		return nil
	}
	cell.occupier = &Occupier{
		grid: g,
		pos:  p,
	}
	return cell.occupier
}

type Occupier struct {
	grid *MarsGrid
	pos  image.Point
}

func (g *MarsGrid) cell(p image.Point) *cell {
	if !p.In(g.bounds) {
		return nil
	}
	return &g.cells[p.Y][p.X]
}
//将给定的Occupier移动至网格中其他单元格，移动失败就会停留在原地
func (o *Occupier) MoveTo(p image.Point) bool {
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	newCell:=o.grid.cell(p)
	if newCell==nil||newCell.occupier!=nil{
		return false
	}
	o.grid.cell(o.pos).occupier=nil
	newCell.occupier=o
	o.pos=p
	return true
}
func (o *Occupier) Sense()SensorData{
	o.grid.mu.Lock()
	defer o.grid.mu.Unlock()
	return o.grid.cell(o.pos).groundData
}
func (o *Occupier)Pos()image.Point{
	return o.pos
}
func main(){
	marsToEarth:=make(chan []Message)
	go earthReceiver(marsToEarth)
	gridSize:=image.Point{X:20,Y:10}
	grid:=NewMarsGrid(gridSize)
	rover:=make([]*RoverDriver,5)
	for i:=range rover{
		rover[i]=startDriver(fmt.Sprint("rover",i),grid,marsToEarth)
	}
	time.Sleep(60*time.Second)
}