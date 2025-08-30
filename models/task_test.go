package models

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestTaskBSONSerialization(t *testing.T) {
	task := Task{
		Id:          bson.NewObjectID(),
		Description: "Test task",
	}

	// Test BSON marshaling
	bsonData, err := bson.Marshal(task)
	require.NoError(t, err)

	// Test unmarshaling
	var unmarshaled Task
	err = bson.Unmarshal(bsonData, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, task.Id, unmarshaled.Id)
	assert.Equal(t, task.Description, unmarshaled.Description)
}

func TestTaskWithEmptyID(t *testing.T) {
	task := Task{
		Description: "Test task without ID",
	}

	// Test that empty ID is handled correctly
	assert.True(t, task.Id.IsZero())

	// Test BSON marshaling with omitempty
	bsonData, err := bson.Marshal(task)
	require.NoError(t, err)

	var bsonMap bson.M
	err = bson.Unmarshal(bsonData, &bsonMap)
	require.NoError(t, err)
	_, exists := bsonMap["_id"]
	assert.False(t, exists, "_id field shouldn't exist when ID is empty")

	// Unmarshal and check that ID is still zero
	var unmarshaled Task
	err = bson.Unmarshal(bsonData, &unmarshaled)
	assert.NoError(t, err)
	assert.True(t, unmarshaled.Id.IsZero())
	assert.Equal(t, task.Description, unmarshaled.Description)
}

func TestViewTaskSerialization(t *testing.T) {
	viewTask := ViewTask{
		Id:          "507f1f77bcf86cd799439011",
		Description: "Test view task",
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(viewTask)
	assert.NoError(t, err)
	assert.Contains(t, string(jsonData), "Test view task")
	assert.Contains(t, string(jsonData), "507f1f77bcf86cd799439011")

	// Test JSON unmarshaling
	var unmarshaled ViewTask
	err = json.Unmarshal(jsonData, &unmarshaled)
	assert.NoError(t, err)
	assert.Equal(t, viewTask.Id, unmarshaled.Id)
	assert.Equal(t, viewTask.Description, unmarshaled.Description)
}

func TestTaskToViewTaskConversion(t *testing.T) {
	task := Task{
		Id:          bson.NewObjectID(),
		Description: "Test task",
	}

	// Convert Task to ViewTask
	viewTask := ViewTask{
		Id:          task.Id.Hex(),
		Description: task.Description,
	}

	assert.Equal(t, task.Id.Hex(), viewTask.Id)
	assert.Equal(t, task.Description, viewTask.Description)
	assert.Len(t, viewTask.Id, 24) // ObjectID hex string length
}
