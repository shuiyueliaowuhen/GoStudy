package main

func recv(c chan int) {
	ret := <-c
	println("接受成功", ret)

}

func main() {

	println("=====无缓冲通道=====")
	ch := make(chan int)
	go recv(ch)
	ch <- 10
	println("发送成功")

	println("=====有缓冲通道=====")
	ch1 := make(chan int, 1)
	ch1 <- 10
	println("发送成功")

}
