package dur

import (
	"flag"
	"fmt"
	"time"
)

func HitungMundur() {

	duration := flag.Duration("durasi", time.Minute*25, "masukkan durasi")
	flag.Parse()

	zero, err := time.ParseDuration("1s")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("duration is", duration)
	for i := zero; i <= *duration; i = i + time.Second {
		fmt.Println("berjalan selama", i)
		time.Sleep(time.Second)
	}
}
