package main

import (
	"flag"
	"github.com/yvasiyarov/gorelic"
)

var newrelicLicense = flag.String("newrelic-license", "", "Newrelic license")

func allocateAndSum(arraySize int) int {
	arr := make([]int, arraySize, arraySize)
	for i, _ := range arr {
		arr[i] = rand.Int()
	}
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	result := 0
	for _, v := range arr {
		result += v
	}
	log.Printf("Array size is: %d, sum is: %d\n", arraySize, result)
	return result
}

func doSomeJob(numRoutines int) {
	for i := 0; i < numRoutines; i++ {
		go allocateAndSum(rand.Intn(1024) * 1024)
	}
}

func main() {

	flag.Parse()
	if *newrelicLicense == "" {
		log.Fatalf("Please, pass a valid newrelic license key.\n Use --help to get more information about available options\n")
	}
	agent := gorelic.NewAgent()
	agent.Verbose = true
	agent.NewrelicLicense = *newrelicLicense
	agent.Run()

	/*
		doSomeJob(100)
		log.Println("All routines started\n")
		select {}
	*/
}
