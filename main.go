package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"time"
)

func draw_guy(x int, y int) {
	termbox.SetCell(x, y, '☺', termbox.Attribute(15), termbox.Attribute(7))
	termbox.SetCell(x, y+1, '✝', termbox.Attribute(15), termbox.Attribute(7))
	termbox.SetCell(x, y+2, '∩', termbox.Attribute(15), termbox.Attribute(7))
}

func draw_world() {
	world := [30]string{}

	world[0] = "                                                                                "
	world[1] = "                                                                                "
	world[2] = "                                                                                "
	world[3] = "                                                                               !"
	world[4] = "                                                                      ----------"
	world[5] = "                                                                     |          "
	world[6] = "                                                                     |          "
	world[7] = "                                                                     |          "
	world[8] = "                                                               -----            "
	world[9] = "                                                             |                  "
	world[10] = "                                                            |                   "
	world[11] = "                                                            |                   "
	world[12] = "                                                       -----                    "
	world[13] = "                                                      |                         "
	world[14] = "                                                      |                         "
	world[15] = "                                                      |                         "
	world[16] = "                                                -----                           "
	world[17] = "                                              |                                 "
	world[18] = "                                              |                                 "
	world[19] = "                                              |                                 "
	world[20] = "                                              |                                 "
	world[21] = "                                           ---                                  "
	world[22] = "                           |              |                                     "
	world[23] = "                           |              |                                     "
	world[24] = "                           |              |                                     "
	world[25] = "            ------------------------------                                      "
	world[26] = "                                                                                "
	world[27] = "                                                                                "
	world[28] = "                                                                                "
	world[29] = "--------------------------------------------------------------------------------"

	colour := func(x int, y int, world [30]string) int {
		switch world[y][x] {
		case ' ':
			return 7
		case '-':
			return 3
		case '|':
			return 3
		case '!':
			return 2
		}
		return 0
	}
	colour(0, 0, world)
	width, height := 80, 30
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			termbox.SetCell(x, y, rune(world[y][x]), termbox.Attribute(colour(x, y, world)), termbox.Attribute(colour(x, y, world)))
		}
	}
}

func draw_all(guy_x int, guy_y int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	draw_world()
	draw_guy(guy_x, guy_y)

	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	jump_cycles := 0
	guy_x := 10
	guy_y := 0

	ticker := time.NewTicker(100 * time.Millisecond)
	quit := make(chan struct{})

	event_queue := make(chan termbox.Event)
	go func() {
		for {
			event_queue <- termbox.PollEvent()
		}
	}()

	go func(t *time.Ticker, j *int, gx *int, gy *int) {
		for {
			select {
			case <-t.C:
				if *j == 0 {
					w, h := termbox.Size()
					if *gx+(*gy+3)*w < w*h {
						c := termbox.CellBuffer()[*gx+(*gy+3)*w]
						if c.Ch == 32 {
							*gy++
							draw_all(*gx, *gy)
						}
					}
				} else if *j > 0 {
					w, _ := termbox.Size()
					if *gx+(*gy-1)*w > 0 {
						c := termbox.CellBuffer()[*gx+(*gy-1)*w]
						if c.Ch == 32 {
							*gy--
							*j--
							draw_all(*gx, *gy)
						}
					}
				}
			case <-quit:
				t.Stop()
				return
			}
		}
		time.Sleep(100 * time.Millisecond)
	}(ticker, &jump_cycles, &guy_x, &guy_y)

	draw_all(guy_x, guy_y)
loop:
	for {
		select {
		case ev := <-event_queue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyEsc:
					termbox.Close()
					break loop
				case termbox.KeyArrowLeft:
					w, h := 80, 30
					if guy_x-1+(guy_y+2)*guy_x < w*h {
						c1 := termbox.CellBuffer()[guy_x-1+guy_y*w]
						c2 := termbox.CellBuffer()[guy_x-1+(guy_y+1)*w]
						c3 := termbox.CellBuffer()[guy_x-1+(guy_y+2)*w]

						if guy_x > 0 && c1.Ch == 32 && c2.Ch == 32 && c3.Ch == 32 {
							guy_x--
						}

						draw_all(guy_x, guy_y)
					}
				case termbox.KeyArrowRight:
					w, h := termbox.Size()
					if guy_x-1+(guy_y+2)*guy_x < w*h {
						c1 := termbox.CellBuffer()[guy_x+1+guy_y*w]
						c2 := termbox.CellBuffer()[guy_x+1+(guy_y+1)*w]
						c3 := termbox.CellBuffer()[guy_x+1+(guy_y+2)*w]

						if guy_x < 79 && c1.Ch == 32 && c2.Ch == 32 && (c3.Ch == 32 || c3.Ch == 33) {
							guy_x++
						}

						draw_all(guy_x, guy_y)
					}
				case termbox.KeyArrowUp:
					w, h := termbox.Size()
					if guy_x+(guy_y+3)*w < w*h {
						c := termbox.CellBuffer()[guy_x+(guy_y+3)*w]
						if c.Ch != 32 && jump_cycles == 0 {
							jump_cycles = 5
						}
					}
				}
			}
		}
		if guy_y == 1 && guy_x == 79 {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			termbox.Close()
			fmt.Println("YOU WIN!")
			fmt.Println()
			break loop
		}
		time.Sleep(10 * time.Millisecond)
	}
	close(quit)
}
