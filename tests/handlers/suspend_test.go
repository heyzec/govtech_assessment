package handlers

import (
	"testing"

	"github.com/heyzec/govtech-assignment/internal/config"
	"github.com/heyzec/govtech-assignment/internal/database"
	"github.com/heyzec/govtech-assignment/internal/errors"
	"github.com/heyzec/govtech-assignment/internal/handlers"
	"github.com/heyzec/govtech-assignment/internal/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccess(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.SuspendParams{
		StudentEmail: "studentmary@gmail.com",
	}

	res, err := handlers.HandleSuspend(params, db)
	assert.Nil(t, res)
	assert.Nil(t, err)
}

func TestNotFound(t *testing.T) {
	cfg := config.LoadEnv()
	db := database.GetGormDB(cfg)

	params := &params.SuspendParams{
		StudentEmail: "nosuchstudent@gmail.com",
	}

	res, err := handlers.HandleSuspend(params, db)
	assert.Nil(t, res)
	var customErr *errors.NotFoundError
	require.ErrorAs(t, err, &customErr)
}
