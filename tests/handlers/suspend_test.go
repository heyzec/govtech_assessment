package handlers_test

import (
	"testing"

	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/handlers"
	"github.com/heyzec/govtech-assignment/internal/helpers/config"
	"github.com/heyzec/govtech-assignment/internal/helpers/database"
	"github.com/heyzec/govtech-assignment/internal/params"
	"github.com/stretchr/testify/require"
)

func TestSuspendSuccess(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.SuspendParams{
		StudentEmail: "studentmary@gmail.com",
	}

	res, err := handlers.HandleSuspend(params, db)
	require.Nil(t, res)
	require.Nil(t, err)
}

func TestSuspendNotFound(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.SuspendParams{
		StudentEmail: "nosuchstudent@gmail.com",
	}

	res, err := handlers.HandleSuspend(params, db)
	require.Nil(t, res)
	var customErr *errors.NotFoundError
	require.ErrorAs(t, err, &customErr)
}
