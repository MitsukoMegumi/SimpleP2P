package cli

import (
	"bufio"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	serverTypes "github.com/mitsukomegumi/Go-Rpcify/networking/types"
)

// Terminal - absctract container holding set of variable with values (runtime only)
type Terminal struct {
	Variables []Variable
}

// Variable - container holding variable values
type Variable struct {
	VariableName string      `json:"name"`
	VariableData interface{} `json:"data"`
	VariableType string      `json:"type"`
}

// NewTerminal - attempts to start io handler for term commands
func NewTerminal() error {
	server := serverTypes.NewServer("server") // Initialize server

	setupCalls(server) // Set calls

	go server.StartServer() // Start server

	term := Terminal{Variables: []Variable{}}

	reader := bufio.NewReader(os.Stdin) // Init reader

	for {
		fmt.Print("\n> ")

		input, err := reader.ReadString('\n') // Search for user input

		if err != nil {
			panic(err) // Panic
		}

		sendCommand(string(input)) // Send to server

		term.HandleCommand(string(input)) // Handle specified command
	}
}

func setupCalls(server *serverTypes.Server) {
	//call, _ := types.NewCall(node.NewNode, "node.NewNode()")
}

func sendCommand(command string) {
	http.Get("http://localhost:8080/call/" + command) // Send command
}

// AddVariable - attempt to append specified variable to terminal variable list
func (term *Terminal) AddVariable(variableName string, variableData interface{}, variableType string) error {
	variable := Variable{VariableName: variableName, VariableData: variableData, VariableType: variableType}

	if reflect.ValueOf(term).IsNil() { // Check for nil variable
		return errors.New("nil terminal found") // Return error
	}

	if len(term.Variables) == 0 { // Check for uninitialized variable array
		term.Variables = []Variable{variable} // Initialize with variable

		return nil // No error occurred, return nil
	}

	term.Variables = append(term.Variables, variable) // Append to array

	return nil // No error occurred, return nil
}

// ReplaceVariable - attempt to replace value at index with specified variable
func (term *Terminal) ReplaceVariable(variableIndex int, variableData interface{}) error {
	if reflect.ValueOf(term).IsNil() { // Check for nil variable
		return errors.New("nil terminal found") // Return error
	}

	if len(term.Variables) == 0 { // Check for uninitialized variable array
		return errors.New("empty terminal environment") // Return found error
	}

	(*term).Variables[variableIndex].VariableData = variableData // Replace value

	return nil
}

// QueryType - attempt to fetch index of variable with matching type
func (term *Terminal) QueryType(variableType string) (uint, error) {
	if variableType == "" { // Check for nil parameter
		return 0, errors.New("invalid type") // Return found error
	}

	if len(term.Variables) == 0 { // Check that terminal environment is not nil
		return 0, errors.New("empty terminal environment") // Return found error
	}

	for x := 0; x != len(term.Variables); x++ { // Declare, increment iterator
		if term.Variables[x].VariableType == variableType { // Check for match
			return uint(x), nil // Return result
		}
	}

	return 0, errors.New("couldn't find matching variable") // Return error
}

// hasVariableSet - checks if specified command sets a variable
func hasVariableSet(command string) bool {
	if strings.HasPrefix(strings.ToLower(command), "var") { // Check for prefix
		return true
	}

	return false
}
