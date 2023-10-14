package handlers

import (
	"testing"

	"github.com/heyzec/govtech-assignment/internal/handlers"
	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	"github.com/heyzec/govtech-assignment/internal/helpers/database"
	"github.com/heyzec/govtech-assignment/internal/params"
	"github.com/heyzec/govtech-assignment/internal/views"
	"github.com/stretchr/testify/require"
)

func TestIntersection(t *testing.T) {
	nums1 := []uint{1, 2, 3, 4}
	nums2 := []uint{3, 4, 5, 6}
	expected := []uint{3, 4}
	output := handlers.Intersection(nums1, nums2)
	require.Equal(t, expected, output)
}

func TestSuccess1(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.CommonStudentsParams{
		TeacherEmails: []string{"teacherken@gmail.com"},
	}

	output, err := handlers.HandleCommonStudents(params, db)
	expected := &views.CommonStudentsView{
		StudentEmails: []string{
			"commonstudent1@gmail.com",
			"commonstudent2@gmail.com",
			"student_only_under_teacher_ken@gmail.com",
		},
	}

	require.Nil(t, err)
	require.Equal(t, output.JSONPayload, expected)
}

func TestSuccess2(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.CommonStudentsParams{
		TeacherEmails: []string{
			"teacherken@gmail.com",
			"teacherjoe@gmail.com",
		},
	}

	output, err := handlers.HandleCommonStudents(params, db)
	expected := &views.CommonStudentsView{
		StudentEmails: []string{
			"commonstudent1@gmail.com",
			"commonstudent2@gmail.com",
		},
	}

	require.Nil(t, err)
	require.Equal(t, output.JSONPayload, expected)
}
