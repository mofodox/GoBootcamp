package tutorials

import "log"

func TutorialEight() {
	sl1 := make([]string, 6)
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	log.Println("Slice size:", len(sl1))

	for i := 0; i < len(sl1); i++ {
		log.Println(sl1[i])
	}

	for _, x := range sl1 {
		log.Println(x)
	}

	sl2 := []int{12, 21, 1974}
	log.Println(sl2)

	sArr := [5]int{
		1, 2, 3, 4, 5,
	}

	sl3 := sArr[0:2]
	log.Println("First 3:", sl3)
	
	// sl3 := sArr[:3]
	// log.Println("First 3:", sl3)

	// sl3 = sArr[2:]
	// log.Println("Last 3:", sl3)

	sArr[0] = 10
	log.Println("sl3:", sl3)

	sl3 = append(sl3, 12)
	log.Println("sl3:", sl3)
	log.Println("sArr:", sArr)

	sl4 := make([]string, 6)
	log.Println("sl4:", sl4)
	log.Println("sl4[0]:", sl4[0])
}