package database_test

import (
	"backend/app/shared/services/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func 

func TestDatabaseConnection(){
	  assert := assert.New(t)
	assert.Equal(true,database.ConnectPostgresSQL() "you have a connection")
	
}