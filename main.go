package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{"Status": "OK"})
}

type RandomPoint struct {
	X            float32
	Y            float32
	InsideCircle bool
}

func randomPoint() RandomPoint {
	x := rand.Float32()*2.0 - 1.0
	y := rand.Float32()*2.0 - 1.0
	return RandomPoint{
		X:            x,
		Y:            y,
		InsideCircle: (x*x + y*y) < 1.0,
	}
}

func randomPoints(count int) []RandomPoint {
	points := make([]RandomPoint, count)
	for i := 0; i < count; i++ {
		points[i] = randomPoint()
	}
	return points
}

func estimatePi(points []RandomPoint) float32 {
	numberInsideCircle := 0
	for _, e := range points {
		if e.InsideCircle == true {
			numberInsideCircle += 1
		}
	}
	return float32(numberInsideCircle) / float32(len(points)) * 4.0
}

func pi(c *gin.Context) {
	count := 1000
	countString := c.Query("count")
	if countString != "" {
		customCount, err := strconv.Atoi(countString)
		if err == nil {
			count = customCount
		}
	}
	points := randomPoints(count)
	c.JSON(200, gin.H{"Points": points, "Pi": estimatePi(points)})
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	r := gin.Default()
        r.StaticFile("/index", "./public/index.html") // for local testing, served by CloudFront when deployed
	r.GET("/api", healthCheck)
	r.GET("/api/health", healthCheck)
	r.GET("/api/pi", pi)
	r.Run()
}
