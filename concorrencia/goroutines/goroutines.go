package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	res, err := calculoETA()
	if err == nil {
		fmt.Println("Funcionou", res)
	}else {
		fmt.Println("Deu merda", err.Error())
	}
}

func calculoETA() ([]string, error) {
	start := time.Now()
	wg := sync.WaitGroup{}
	inputs := []int{1,3,5,7,9,2,11,13,15}
	results := make([]string, 0, len(inputs))
	var errors []error

	for i, v := range inputs {
		wg.Add(1)
		go func(index, num int) {
			r, err := gmaps(num)
			if err == nil {
				results = append(results, r)
			}else{
				errors = append(errors, err)
			}
			defer wg.Done()
		}(i,v)
	}
	wg.Wait()
	totalTimeInMS := float64(time.Since(start) / time.Millisecond)
	fmt.Printf("CalculateETA GMaps total time to finish: %vms\n", totalTimeInMS)

	if len(errors) > 0 {
		return nil, errors[0]
	}
	for i, s := range results {
		fmt.Println("Indice: ", i, "Value: ", s)
	}
	return results, nil
}

func gmaps(num int) (string, error) {
	time.Sleep( (time.Duration(num) * time.Second))
	if num % 2 ==0 {
		return "", fmt.Errorf("Número %d par não permitido", num);
	}
	return fmt.Sprintf("Bem vindo número %d impar!", num), nil
}


