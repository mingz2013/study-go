package main

import (
	"sync"
	robot2 "github.com/mingz2013/study.go/test-910-game-server/robot"
	table2 "github.com/mingz2013/study.go/test-910-game-server/table"
	"github.com/mingz2013/study.go/test-910-game-server/msg"
	"log"
)

type Processor interface {
	Run()
}

func runProcessor(wg sync.WaitGroup, p Processor) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		p.Run()

	}()
}

type RobotContext struct {
	Robot  robot2.Robot
	MsgIn  chan msg.Msg
	MsgOut chan msg.Msg
}

func makeRobots() []RobotContext {
	var robots []RobotContext
	for i := 0; i < 4; i++ {
		robotMsgIn := make(chan msg.Msg)
		robotMsgOut := make(chan msg.Msg)

		robot := robot2.NewRobot(i+1000, "", robotMsgIn, robotMsgOut)

		robots = append(robots, RobotContext{robot, robotMsgIn, robotMsgOut})

	}
	return robots
}

func main() {
	//sdk := sdk2.MakerSdk("1")

	//tableManager := table.NewTableManager("1")

	var wg sync.WaitGroup // 工作goroutine个数

	//go sdk.Run()
	//go tableManager.Run()
	//RunProcessor(wg, sdk)
	//RunProcessor(wg, tableManager)

	tableMsgIn := make(chan msg.Msg)
	tableMsgOut := make(chan msg.Msg)

	table := table2.NewTable("1", tableMsgIn, tableMsgOut)

	//var robots []RobotContext

	robots := makeRobots()

	log.Println("make obj down")

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			m, ok := <-tableMsgOut
			log.Println("on msg table", m)
			if !ok {
				continue
			}
			id := m["id"].(int)

			for i := 0; i < len(robots); i++ {
				if robots[i].Robot.Id == id {
					robots[i].MsgIn <- m
				}
			}

		}

	}()

	for i := 0; i < len(robots); i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for {
				m, ok := <-robots[index].MsgOut
				log.Println("on msg robot", index, m)
				if !ok {
					continue
				}
				tableMsgIn <- m
			}

		}(i)
	}

	log.Println("bound ch down")

	runProcessor(wg, table)

	for i := 0; i < len(robots); i++ {

		runProcessor(wg, robots[i].Robot)
	}

	log.Println("run down")

	wg.Wait()
}
