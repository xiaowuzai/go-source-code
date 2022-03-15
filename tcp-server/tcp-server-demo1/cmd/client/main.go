package main

import "sync"
func main(){
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		go func(i int){
			defer wg.Done()
			startClient(i)
		}(i)
	}

	wg.Wait()
}

func startClient(i int) {

}



