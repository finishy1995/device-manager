package logic

import (
	"math"
	"math/rand"
)

var (
	GradualDriftEuclideanDistance   []float64
	GradualDriftAngleBetweenVectors []float64
	FixedEuclideanDistance          []float64
	FixedAngleBetweenVectors        []float64
	RandomEuclideanDistance         []float64
	RandomAngleBetweenVectors       []float64
)

func init() {
	GradualDriftEuclideanDistance, GradualDriftAngleBetweenVectors = generateGradualDriftVector()
	FixedEuclideanDistance, FixedAngleBetweenVectors = generateFixedVector()
	RandomEuclideanDistance, RandomAngleBetweenVectors = generateRandomVector()
}

var MaxDistance = 180 * math.Sqrt(3)

// angleDifference calculates the smallest difference between two angles, considering the periodicity of angles
func angleDifference(a, b float64) float64 {
	diff := math.Abs(a - b)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff
}

// EuclideanDistanceWithNormalize calculates the Euclidean distance between two points in 3D space, then normalizes it to [0, 1]
func EuclideanDistanceWithNormalize(x1, y1, z1, x2, y2, z2 float64) float64 {
	// Calculate the periodic differences for each dimension
	dx := angleDifference(x1, x2)
	dy := angleDifference(y1, y2)
	dz := angleDifference(z1, z2)

	// Compute the Euclidean distance using the periodic differences
	distance := math.Sqrt(dx*dx + dy*dy + dz*dz)

	// Normalize the distance
	return distance / MaxDistance
}

func AngleBetweenVectorsWithNormalize(x1, y1, z1, x2, y2, z2 float64) float64 {
	// Compute the dot product of the two vectors
	dotProduct := x1*x2 + y1*y2 + z1*z2

	// Compute the magnitudes of the two vectors
	magnitude1 := math.Sqrt(x1*x1 + y1*y1 + z1*z1)
	magnitude2 := math.Sqrt(x2*x2 + y2*y2 + z2*z2)

	// Avoid division by zero
	if magnitude1 == 0 || magnitude2 == 0 {
		return 0
	}

	// Compute the cosine of the angle
	cosTheta := dotProduct / (magnitude1 * magnitude2)

	// Clamp cosTheta to the range [-1, 1] to avoid numerical issues with arccos
	cosTheta = math.Max(-1, math.Min(1, cosTheta))

	// Compute the angle in radians
	return math.Acos(cosTheta) / math.Pi
}

func generateGradualDriftVector() ([]float64, []float64) {
	euclideanDistance := EuclideanDistanceWithNormalize(0, 0, 0, 0.3, 0.3, 0.3)
	angleBetweenVectors := AngleBetweenVectorsWithNormalize(0, 0, 0, 0.3, 0.3, 0.3)
	euclideanDistanceVector := make([]float64, 0, 600)
	angleBetweenVectorsVector := make([]float64, 0, 600)
	for i := 0; i < 600; i++ {
		euclideanDistanceVector = append(euclideanDistanceVector, euclideanDistance)
		angleBetweenVectorsVector = append(angleBetweenVectorsVector, angleBetweenVectors)
	}
	return euclideanDistanceVector, angleBetweenVectorsVector
}

func generateFixedVector() ([]float64, []float64) {
	euclideanDistance := EuclideanDistanceWithNormalize(1, 1, 1, 1, 1, 1)
	angleBetweenVectors := AngleBetweenVectorsWithNormalize(1, 1, 1, 1, 1, 1)
	euclideanDistanceVector := make([]float64, 0, 600)
	angleBetweenVectorsVector := make([]float64, 0, 600)
	for i := 0; i < 600; i++ {
		euclideanDistanceVector = append(euclideanDistanceVector, euclideanDistance)
		angleBetweenVectorsVector = append(angleBetweenVectorsVector, angleBetweenVectors)
	}
	return euclideanDistanceVector, angleBetweenVectorsVector
}

func generateRandomVector() ([]float64, []float64) {
	euclideanDistanceVector := make([]float64, 0, 600)
	angleBetweenVectorsVector := make([]float64, 0, 600)
	for i := 0; i < 600; i++ {
		x1 := float64(rand.Intn(36000)) / 100.0
		y1 := float64(rand.Intn(36000)) / 100.0
		z1 := float64(rand.Intn(36000)) / 100.0
		x2 := float64(rand.Intn(36000)) / 100.0
		y2 := float64(rand.Intn(36000)) / 100.0
		z2 := float64(rand.Intn(36000)) / 100.0
		euclideanDistanceVector = append(euclideanDistanceVector, EuclideanDistanceWithNormalize(x1, y1, z1, x2, y2, z2))
		angleBetweenVectorsVector = append(angleBetweenVectorsVector, AngleBetweenVectorsWithNormalize(x1, y1, z1, x2, y2, z2))
	}
	return euclideanDistanceVector, angleBetweenVectorsVector
}
