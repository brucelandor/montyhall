package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	rounds := int(1e6) // 游戏进行的场数
	type game struct {
		change  bool
		winRate float64
	}
	rc := make(chan game, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		winRate := montyhall(rounds, false)
		rc <- game{winRate: winRate, change: false}

	}()
	go func() {
		defer wg.Done()
		winRate := montyhall(rounds, true)
		rc <- game{winRate: winRate, change: true}

	}()
	wg.Wait()
	close(rc)
	for game := range rc {
		fmt.Printf("%+v\n", game)
	}
}

func montyhall(rounds int, change bool) float64 {
	n := 3
	wins := 0
	for i := 0; i < rounds; i++ {
		car := rand.Intn(n)
		guess := rand.Intn(n)
		if !change {
			// 不改变
			if guess == car {
				wins++
			}
		} else {
			// 改变，区分猜测是不是对的
			if guess == car {
				// 猜的是车，改变，输
				continue
			}
			// 猜的不是车，主持人打开另一个不是车的门, 改变，赢
			wins++
		}
	}
	return float64(wins) / float64(rounds)
}
