package main

import (
	"fmt"
	"sync"
	"time"
)

// 組合(不分順序)
func Combination() {

	var bb = &BallPick{}

	//bb.getBall2(11)
	//bb.getBall3(30)
	//bb.getBall4(11)
	//bb.getBall5(11)
	s := time.Now()
	//bb.getBall6(49)
	bb.getBall7(49)
	fmt.Println("花費時間:", time.Since(s).Seconds(), " 秒")

	// 取7數量: 85900584
	// 花費時間: 39.971854635  秒

}

// BallPick 取5顆球
type BallPick struct {
	wg     *sync.WaitGroup
	locker *sync.RWMutex
	result [][]int
}

func (b *BallPick) setResult(pos [][]int) {
	b.locker.Lock()
	b.result = append(b.result, pos...)
	b.locker.Unlock()
}

func (b *BallPick) getBall2(maxBall int) {

	b.result = [][]int{}
	b.wg = new(sync.WaitGroup)
	b.locker = new(sync.RWMutex)

	var n = 2
	for i := 1; i <= (maxBall - n + 1); i++ {
		b.wg.Add(1)
		go b.getBall2sub(i, maxBall)
	}
	b.wg.Wait()

	fmt.Println(b.result)
	fmt.Println("取2數量:", len(b.result))
}

func (b *BallPick) getBall2sub(p1, maxBall int) {

	var (
		position [][]int
		m        = maxBall
	)
	// 取2
	for p2 := p1 + 1; p2 <= m; p2++ {
		position = append(position, []int{p2, p1})
	}

	b.setResult(position)
	b.wg.Done()

}

func (b *BallPick) getBall3(maxBall int) {

	b.result = [][]int{}
	b.wg = new(sync.WaitGroup)
	b.locker = new(sync.RWMutex)

	var n = 3
	for i := 1; i <= (maxBall - n + 1); i++ {
		b.wg.Add(1)
		go b.getBall3sub(i, maxBall)
	}
	b.wg.Wait()

	fmt.Println(b.result)
	fmt.Println("取3數量:", len(b.result))
}

func (b *BallPick) getBall3sub(p1, maxBall int) {

	var (
		position [][]int
		m        = maxBall
	)
	// 取3
	for p2 := p1 + 1; p2 <= m; p2++ {
		for p3 := p2 + 1; p3 <= m; p3++ {
			position = append(position, []int{p3, p2, p1})
		}
	}

	b.setResult(position)
	b.wg.Done()

}

func (b *BallPick) getBall4(maxBall int) {

	b.result = [][]int{}
	b.wg = new(sync.WaitGroup)
	b.locker = new(sync.RWMutex)

	var n = 4
	for i := 1; i <= (maxBall - n + 1); i++ {
		b.wg.Add(1)
		go b.getBall4sub(i, maxBall)
	}
	b.wg.Wait()

	fmt.Println(b.result)
	fmt.Println("取4數量:", len(b.result))
}

func (b *BallPick) getBall4sub(p1, maxBall int) {

	var (
		position [][]int
		m        = maxBall
	)
	// 取3
	for p2 := p1 + 1; p2 <= m; p2++ {
		for p3 := p2 + 1; p3 <= m; p3++ {
			for p4 := p3 + 1; p4 <= m; p4++ {
				position = append(position, []int{p4, p3, p2, p1})
			}
		}
	}

	b.setResult(position)
	b.wg.Done()

}

func (b *BallPick) getBall5(maxBall int) {

	b.result = [][]int{}
	b.wg = new(sync.WaitGroup)
	b.locker = new(sync.RWMutex)

	var n = 5
	for i := 1; i <= (maxBall - n + 1); i++ {
		b.wg.Add(1)
		go b.getBall5sub(i, maxBall)
	}
	b.wg.Wait()

	fmt.Println(b.result)
	fmt.Println("取5數量:", len(b.result))
}

func (b *BallPick) getBall5sub(p1, maxBall int) {

	var (
		position [][]int
		m        = maxBall
	)
	// 取5
	for p2 := p1 + 1; p2 <= m; p2++ {
		for p3 := p2 + 1; p3 <= m; p3++ {
			for p4 := p3 + 1; p4 <= m; p4++ {
				for p5 := p4 + 1; p5 <= m; p5++ {
					position = append(position, []int{p5, p4, p3, p2, p1})
				}
			}
		}
	}

	b.setResult(position)
	b.wg.Done()

}

func (b *BallPick) getBall6(maxBall int) {

	b.result = [][]int{}
	b.wg = new(sync.WaitGroup)
	b.locker = new(sync.RWMutex)

	var n = 6
	for i := 1; i <= (maxBall - n + 1); i++ {
		b.wg.Add(1)
		go b.getBall6sub(i, maxBall)
	}
	b.wg.Wait()

	//fmt.Println(b.result)
	fmt.Println("取6數量:", len(b.result))
}

func (b *BallPick) getBall6sub(p1, maxBall int) {

	var (
		position [][]int
		m        = maxBall
	)
	// 取5
	for p2 := p1 + 1; p2 <= m; p2++ {
		for p3 := p2 + 1; p3 <= m; p3++ {
			for p4 := p3 + 1; p4 <= m; p4++ {
				for p5 := p4 + 1; p5 <= m; p5++ {
					for p6 := p5 + 1; p6 <= m; p6++ {
						position = append(position, []int{p6, p5, p4, p3, p2, p1})
					}
				}
			}
		}
	}

	b.setResult(position)
	b.wg.Done()

}

func (b *BallPick) getBall7(maxBall int) {

	b.result = [][]int{}
	b.wg = new(sync.WaitGroup)
	b.locker = new(sync.RWMutex)

	var n = 7
	for i := 1; i <= (maxBall - n + 1); i++ {
		b.wg.Add(1)
		go b.getBall7sub(i, maxBall)
	}
	b.wg.Wait()

	//fmt.Println(b.result)
	fmt.Println("取7數量:", len(b.result))
}

func (b *BallPick) getBall7sub(p1, maxBall int) {

	var (
		position [][]int
		m        = maxBall
	)
	// 取5
	for p2 := p1 + 1; p2 <= m; p2++ {
		for p3 := p2 + 1; p3 <= m; p3++ {
			for p4 := p3 + 1; p4 <= m; p4++ {
				for p5 := p4 + 1; p5 <= m; p5++ {
					for p6 := p5 + 1; p6 <= m; p6++ {
						for p7 := p6 + 1; p7 <= m; p7++ {
							position = append(position, []int{p7, p6, p5, p4, p3, p2, p1})
						}
					}
				}
			}
		}
	}

	b.setResult(position)
	b.wg.Done()

}
