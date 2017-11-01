package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"
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
	return elevator{currentFloor: 0, direction: DOWN}
}

func (e elevator) move() {

}

func startElevator() {
	elevator := newElevaror()
	for true {
		elevator.move()
	}
}

func main() {

	go startElevator()

	fmt.Println("I'm an elevator.")

	floors := flag.Int("floors", 1, "Number of floors")
	height := flag.Float64("height", 1, "Height of a floor")
	speed := flag.Float64("speed", 1, "Lift speed")
	openTime := flag.Float64("speed", 1, "Lift speed")

	floorSpeed := time.Duration(*height / *speed * 1000000000) // nanoseconds for 1 floor
	time.Sleep(floorSpeed)

	var input string

	for input != "exit" {

		fmt.Scan(&input)
		fmt.Println(input)

		if input == "T" {

		}

		if input != "T" {
			floor, err := strconv.ParseInt(input, 10, 64)
			if err != nil {
				fmt.Println("Floor is an integer.")
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
