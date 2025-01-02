package main

func generateVertices(pattern [][]int) []float32 {
	var vertices []float32
	maxY := len(pattern)
	maxX := len(pattern[0])

	// Calculate the block size so that the grid fits within the [-1, 1] range
	blockSizeX := 2.0 / float32(maxX)
	blockSizeY := 2.0 / float32(maxY)

	for y, row := range pattern {
		for x, cell := range row {
			if cell == 1 {
				startX := float32(x)*blockSizeX - 1.0
				startY := 1.0 - float32(y+1)*blockSizeY

				vertices = append(vertices,
					startX, startY, 0.0,
					startX, startY+blockSizeY, 0.0,
					startX+blockSizeX, startY+blockSizeY, 0.0,
				)
				// Top-right triangle of the quad
				vertices = append(vertices,
					startX, startY, 0.0,
					startX+blockSizeX, startY+blockSizeY, 0.0,
					startX+blockSizeX, startY, 0.0,
				)
			}
		}
	}
	return vertices
}
