package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
	_ "errors"
	_ "os"
	"os"
)

const ( // iota is reset to 0
	DOWN = iota
	UP
)

type elevator struct {
	openTime     time.Duration
	passTime     time.Duration
	maxFloor     int
	currentFloor int
	direction    int
	buttons      [10]bool
}

func newElevaror() elevator {
	return elevator{currentFloor: 0, direction: UP}
}

func (e *elevator) move(moves chan struct{}) {

	// direction  change logic
	if e.currentFloor == e.maxFloor - 1 {
		e.direction = DOWN
	}

	if e.currentFloor == 0 {
		e.direction = UP
	}

	if e.direction == UP && e.nextFloor() == -1 {
		e.direction = DOWN
	}
	if e.direction == DOWN && e.nextFloor() == -1 {
		e.direction = UP
	}


	fmt.Fprintf(os.Stderr, "Floor %d. Moving %d\n", e.currentFloor, e.direction)
	//time.Sleep(time.Duration(e.passTime))
	time.Sleep(time.Duration(time.Second))


	if e.direction == UP {
		e.currentFloor++
	}
	if e.direction == DOWN {
		e.currentFloor--
	}

	if e.buttons[e.currentFloor] { //the floor button we've arrived at was on
		fmt.Fprintf(os.Stderr, "Opening doors on floor %d\n", e.currentFloor)
		time.Sleep(time.Duration(e.passTime))
		e.buttons[e.currentFloor] = false
	}

	moves <- struct{}{}
}

func (e elevator) nextFloor() int {
	var nextFloor int

	if e.direction == UP {
		nextFloor = e.currentFloor + 1
	}
	if e.direction == DOWN {
		nextFloor = e.currentFloor - 1
	}

	if (nextFloor > e.maxFloor - 1) || (nextFloor < 0) {
		return -1
	}

	if e.direction == UP {
		for i := e.currentFloor + 1; i < e.maxFloor; i++ {
			if e.buttons[i] {
				return i
			}
		}
	}

	if e.direction == DOWN {
		for i := e.currentFloor - 1; i >= 0; i-- {
			if e.buttons[i] {
				return i
			}
		}
	}

	return nextFloor
}

func startElevator(commands chan int) {

	elevator := newElevaror()
	moves := make(chan struct{})
	go elevator.move(moves)

	for true {
		select {
		case move := <-moves:
			fmt.Println("Got a move!\n\n")
			_ = move
			go elevator.move(moves)

		case command := <-commands:
			fmt.Printf("GOT COMMAND: %d", command)
			elevator.buttons[command-1] = true
		default:

		}
	}

}

func main() {

	fmt.Println("I'm an elevator.")

	floors := flag.Int64("floors", 10, "Number of floors")
	height := flag.Float64("height", 1, "Height of a floor")
	speed := flag.Float64("speed", 1, "Lift speed")
	//openTime := flag.Float64("open", 1, "Door open time")

	flag.Parse()

	commands := make(chan int, *floors)
	go startElevator(commands)

	floorSpeed := time.Duration(*height / *speed * 1000000000) // nanoseconds for 1 floor
	time.Sleep(floorSpeed)

	var input string
	for input != "exit" {

		fmt.Scan(&input)
		//fmt.Println(input)

		if input == "T" {
			commands <- 0
		}

		if input != "T" {
			floor, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Floor is an integer.")
			} else {

				if (floor > *floors - 1) || (floor < 0) {
					fmt.Printf("Floor is an integer between 0 and %d\n", *floors-1)
				} else {
					commands <- int(floor)
				}
			}
		}
	}
}

/*
Предлагаю написать программу «симулятор лифта».

Программа запускается из коммандной строки, в качестве параметров задается:
1) кол-во этажей в подъезде - N (от 5 до 20);
2) высота одного этажа;
3) скорость лифта при движении в метрах в секунду (ускорением пренебрегаем, считаем, что когда лифт едет - он сразу едет с определенной скоростью);
4) время между открытием и закрытием дверей.

После запуска программа должна постоянно ожидать ввода от пользователя и выводить действия лифта в реальном времени.
События, которые нужно выводить:
1) лифт проезжает некоторый этаж;
2) лифт открыл двери;
3) лифт закрыл двери.

Возможный ввод пользователя:
1) вызов лифта на этаж из подъезда;
2) нажать на кнопку этажа внутри лифта.

Считаем, что пользователь не может помешать лифту закрыть двери.
Все данные, которых не хватает в задаче можно выбрать на свое усмотрение.

Хочется видеть более-менее современный лифт. Например, чтобы можно было нажать вызов на нескольких этажах и он поехал в ближайший, а не первый нажатый.
*/
