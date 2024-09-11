package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Board struct {
	playing_field [4][4]int
}

func (b Board) getField() [4][4]int {
	return b.playing_field
}

func (b *Board) setField(f [4][4]int) {
	b.playing_field = f
}
func (b *Board) rotateMatrix() {

	matrix := b.playing_field
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	b.playing_field = matrix
}

func (b Board) outputField() {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(math.Pow(2, float64(b.playing_field[i][j])), " ")
		}
		fmt.Print("\n")
	}
}
func (b *Board) moveLeft() {
	var m [4]int
	var c int
	for i := 0; i < 4; i++ {
		c = 0
		m = [4]int{}
		for j := 0; j < 4; {
			if j == 3 {
				m[c] = b.playing_field[i][j]
			} else if b.playing_field[i][j] == b.playing_field[i][j+1] && b.playing_field[i][j] != 0 {
				m[c] = b.playing_field[i][j] + 1
				j++
				c++
			} else if b.playing_field[i][j] != 0 {
				m[c] = b.playing_field[i][j]
				c++
			}
			j++
		}
		b.playing_field[i] = m
	}
}
func (b *Board) moveRight() {
	b.rotateMatrix()
	b.rotateMatrix()
	b.moveLeft()
	b.rotateMatrix()
	b.rotateMatrix()
}
func (b *Board) moveDown() {
	b.rotateMatrix()
	b.moveLeft()
	b.rotateMatrix()
	b.rotateMatrix()
	b.rotateMatrix()
}
func (b *Board) moveUp() {
	b.rotateMatrix()
	b.rotateMatrix()
	b.rotateMatrix()
	b.moveLeft()
	b.rotateMatrix()
}
func (b *Board) MakeANewNum() {
	r := rand.New(rand.NewSource(99))
	slice := make([][]int, 0, 15)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if b.playing_field[i][j] == 0 {
				slice = append(slice, []int{i, j})
			}
		}
	}
	n := r.Intn(len(slice))
	num := r.Intn(10)
	if num == 0 {
		b.playing_field[slice[n][0]][slice[n][1]] = 2
	} else {
		b.playing_field[slice[n][0]][slice[n][1]] = 1
	}

}

func main() {
	brd := Board{}
	var x string
	var ans bool
	brd.MakeANewNum()
	for {
		ans = true
		brd.outputField()
		fmt.Scan(&x)
		if x == "r" {
			brd.moveRight()
		} else if x == "l" {
			brd.moveLeft()
		} else if x == "u" {
			brd.moveUp()
		} else if x == "d" {
			brd.moveDown()
		} else if x == "e" {
			fmt.Println("Спасибо за игру!")
			break
		} else {
			fmt.Println("Неверная команда")
			ans = false
		}
		if ans {
			brd.MakeANewNum()
		}
	}
}
