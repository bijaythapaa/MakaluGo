package loops

func main() {

	// breaking out from outer level of loop
breakpoint:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j && j == 1 {
				break breakpoint
			}
		}
	}

	// same as above
	// b := false
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		if i == j && j == 1 {
	// 			b = true
	// 			break
	// 		}
	// 	}
	// 	if b == true {
	// 		break
	// 	}
	// }

	// forever-loop
	// for {
	// 	fmt.Println("will only execute this block once")
	// 	return
	// }
	// since, this return inside forever loop will returns the control out out the main function,
	// any codes below this forver-loop will not be executed.
}
