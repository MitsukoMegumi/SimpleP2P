package shard

import (
	"testing"

	"github.com/mitsukomegumi/GoP2P/common"
	"github.com/mitsukomegumi/GoP2P/types/environment"
	"github.com/mitsukomegumi/GoP2P/types/node"
)

/*
	BEGIN EXPORTED METHODS
*/

// TestNewShard - test functionality of shard initializer
func TestNewShard(t *testing.T) {
	node, err := newNodeSafe() // Initialize shard node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShard(node) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("Initialized new shard with address %s", shard.Address) // Log new shard
}

// TestNewShardWithNodes - test functionality of shard initializer
func TestNewShardWithNodes(t *testing.T) {
	localNode, err := newNodeSafe() // Initialize shard node

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	shard, err := NewShardWithNodes(&[]node.Node{*localNode, *localNode}) // Init shard

	if err != nil { // Check for errors
		t.Errorf(err.Error()) // Log found error
		t.FailNow()           // Panic
	}

	t.Logf("Initialized new shard with address %s", shard.Address) // Log new shard
}

/*
	END EXPORTED METHODS
*/

/*
	BEGIN INTERNAL METHODS
*/

func newNodeSafe() (*node.Node, error) {
	ip, err := common.GetExtIPAddrWithoutUPnP() // Fetch IP address

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	environment, _ := environment.NewEnvironment() // Create new environment

	if err != nil { // Check for errors
		return &node.Node{}, err // Return found error
	}

	node := node.Node{Address: ip, Reputation: 0, IsBootstrap: false, Environment: environment} // Creates new node instance with specified address

	return &node, nil // Return initialized node
}

/*
	END INTERNAL METHODS
*/
