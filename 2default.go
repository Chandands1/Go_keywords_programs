//default : execute when no case matches in a switch

func exampleDefault(){
	x := 10
	switch x{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")	
	default:
		fmt.Println("Number doesn't match")	
	}

}