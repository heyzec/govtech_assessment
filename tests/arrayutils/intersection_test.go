package handlers

import (
	"testing"

	"github.com/heyzec/govtech-assignment/internal/helpers/arrayutils"
	"github.com/stretchr/testify/require"
)

func TestIntersection1(t *testing.T) {
	nums1 := []uint{1, 2, 3, 4}
	nums2 := []uint{3, 4, 5, 6}
	expected := []uint{3, 4}
	output := arrayutils.Intersection(nums1, nums2)
	require.Equal(t, expected, output)
}

func TestIntersectionEmpty(t *testing.T) {
	nums1 := []uint{}
	nums2 := []uint{}
	expected := []uint{}
	output := arrayutils.Intersection(nums1, nums2)
	require.Equal(t, expected, output)
}

func TestIntersectionDisjoint(t *testing.T) {
	nums1 := []uint{3, 1, 4, 2}
	nums2 := []uint{8, 6, 7, 0}
	expected := []uint{}
	output := arrayutils.Intersection(nums1, nums2)
	require.Equal(t, expected, output)
}
