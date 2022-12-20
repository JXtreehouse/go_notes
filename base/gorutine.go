package main

// 运行结果 : fatal error: concurrent map read and map write
func main() {
	mymap := make(map[int]int)
	go func() {
		for {
			_ = mymap[1]
		}

	}()
	go func() {
		for {
			mymap[2] = 2
		}
	}()
	select {}
}
