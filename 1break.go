// break keyword : stop the nearest loop

func exampleBreak(){
	for i :=1; i<=10; i++{
		if i == 5 {
			break
		}
		fmt.Println(i)

	}
}