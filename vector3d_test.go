package vector3d

import (
	"math"
	"testing"
)

func TestVector3D_DotProduct(t *testing.T) {
	tt := []struct {
		name     string
		v1, v2   *Vector3D
		expected float64
	}{
		{
			name:     "Zero vector dot Zero vector",
			v1:       New(0, 0, 0),
			v2:       New(0, 0, 0),
			expected: 0,
		},
		{
			name:     "Zero vector dot Positive vector",
			v1:       New(1, 1, 1),
			v2:       New(0, 0, 0),
			expected: 0,
		},
		{
			name:     "Zero vector dot Negative vector",
			v1:       New(1, 1, 1),
			v2:       New(0, 0, 0),
			expected: 0,
		},
		{
			name:     "Positive vector dot Positive vector",
			v1:       New(1, 1, 1),
			v2:       New(1, 1, 1),
			expected: 3,
		},
		{
			name:     "Negative vector dot Negative vector",
			v1:       New(-1, -1, -1),
			v2:       New(-1, -1, -1),
			expected: 3,
		},
		{
			name:     "Positive vector dot Negative vector",
			v1:       New(1, 1, 1),
			v2:       New(-1, -1, -1),
			expected: -3,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.v1.DotProduct(tc.v2)
			if !equalFloat64(got, tc.expected) {
				t.Fatalf("unexpected error: got=%v but expected=%v", got, tc.expected)
			}
		})
	}
}

func TestVector3D_Norm(t *testing.T) {
	tt := []struct {
		name     string
		vec      *Vector3D
		expected float64
	}{
		{
			name:     "Zero vector",
			vec:      New(0, 0, 0),
			expected: 0,
		},
		{
			name:     "Positive vector",
			vec:      New(1, 1, 1),
			expected: math.Sqrt(3),
		},
		{
			name:     "Negative vector",
			vec:      New(-1, -1, -1),
			expected: math.Sqrt(3),
		},
		{
			name:     "2D vector",
			vec:      New(1, 1, 0),
			expected: math.Sqrt(2),
		},
		{
			name:     "1D vector",
			vec:      New(1, 0, 0),
			expected: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.vec.Norm()
			if !equalFloat64(got, tc.expected) {
				t.Fatalf("unexpected error: got=%v but expected=%v", got, tc.expected)
			}
		})
	}
}

func TestVector3D_Normalize(t *testing.T) {
	tt := []struct {
		name          string
		vec           *Vector3D
		expectedErr   bool
		expectedValue float64
	}{
		{
			name:          "Zero vector",
			vec:           New(0, 0, 0),
			expectedValue: 0,
		},
		{
			name:          "Positive vector",
			vec:           New(1, 1, 1),
			expectedValue: 1,
		},
		{
			name:          "Negative vector",
			vec:           New(-1, -1, -1),
			expectedValue: 1,
		},
		{
			name:          "2D vector",
			vec:           New(1, 1, 0),
			expectedValue: 1,
		},
		{
			name:          "1D vector",
			vec:           New(1, 0, 0),
			expectedValue: 1,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.vec.Normalize()

			if !equalFloat64(got.Norm(), tc.expectedValue) {
				t.Fatalf("unexpected error: got=%v but expected=%v", got.Norm(), tc.expectedValue)
			}
		})
	}
}

func TestVector3D_IsEqual(t *testing.T) {
	tt := []struct {
		name     string
		v1       *Vector3D
		v2       *Vector3D
		expected bool
	}{
		{
			name:     "Zero vector",
			v1:       New(0, 0, 0),
			v2:       New(1, 1, 1),
			expected: false,
		},
		{
			name:     "Reversed vector",
			v1:       New(1, 2, 3),
			v2:       New(3, 2, 1),
			expected: false,
		},
		{
			name:     "Equal vectors",
			v1:       New(1, 2, 3),
			v2:       New(1, 2, 3),
			expected: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.v1.IsEqual(tc.v2)

			if got != tc.expected {
				t.Fatalf("unexpected error: got=%v but expected=%v", got, tc.expected)
			}
		})
	}
}
func TestVector3D_CrossProduct(t *testing.T) {
	tt := []struct {
		name     string
		v1       *Vector3D
		v2       *Vector3D
		expected *Vector3D
	}{
		{
			name:     "Zero vector",
			v1:       New(0, 0, 0),
			v2:       New(0, 0, 0),
			expected: New(0, 0, 0),
		},
		{
			name:     "Zero vector with Positive vector",
			v1:       New(0, 0, 0),
			v2:       New(1, 2, 3),
			expected: New(0, 0, 0),
		},
		{
			name:     "Random vector",
			v1:       New(1, 3, 3),
			v2:       New(7, 1, 0),
			expected: New(-3, 21, -20),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.v1.CrossProduct(tc.v2)

			if !got.IsEqual(tc.expected) {
				t.Fatalf("unexpected error: got=%v but expected=%v", got, tc.expected)
			}
		})
	}
}
