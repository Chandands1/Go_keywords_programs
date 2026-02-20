// interface : defines behavior(method that a type must implement)


type Speaker interface{
	speak()
}
type Person stuct{}


func (p Person)Speak(){
	fmt.Println("I am speaking")
}

func main() {
	var s Speaker = Person{}
	s.Speak()
}