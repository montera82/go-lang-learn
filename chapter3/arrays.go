package main

import "fmt"

func main() {

	// var x [5]float64
	//
	// x[0] = 98
	// x[1] = 93
	// x[2] = 77
	// x[3] = 82
	// x[4] = 83

	// x := []float64{ 98, 93, 77, 82, 83, }
	//
	// y := make([]float64, 5)
	//
	// y = append (y, 9,9)
	//
	// fmt.Println(y)
	//
	// var total float64 = 0
	//
	// // for i := 0; i < len(x); i++ {
	// // 	total += x[i]
	// // }
	//
	// for _, value := range x {
	// 	total += value;
	// }
	//
	// fmt.Println(total/float64(len(x)))

	//MAPS

	// 	elements := map[string]map[string]string{
	// 		"P": map[string]string{
	// 			"name":  "Potassium",
	// 			"state": "Solid",
	// 		},
	// 		"H": map[string]string{
	// 			"name": "Hydrogen",
	//
	// 			"state": "Gas",
	// 		},
	// 		"He": map[string]string{
	// 			"name":  "Helium",
	// 			"state": "gas",
	// 		},
	// 	}
	//
	// 	if name, ok := elements["P"]; ok {
	// 	fmt.Println(name, ok)
	// }

	x := []int{48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	var smallest int

	for i := 0; i < len(x); i++ {

		if i == len(x)-1 {
			break
		}

		if x[i] < x[i+1] {
			smallest = x[i]
		}
	}

	fmt.Println(smallest)
}
