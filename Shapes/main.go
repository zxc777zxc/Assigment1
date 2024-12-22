package Shapes

func main() {
	shapes := []Shape{
		Rectangle{length: 10, width: 5},
		Circle{radius: 7},
		Square{length: 4},
		Triangle{sideA: 3, sideB: 4, sideC: 5},
	}

	for _, shape := range shapes {
		PrintShapeDetails(shape)
	}
}
