package types

import (
"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TestDataList: []TestData{},
// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in testData
testDataIdMap := make(map[uint64]bool)
testDataCount := gs.GetTestDataCount()
for _, elem := range gs.TestDataList {
	if _, ok := testDataIdMap[elem.Id]; ok {
		return fmt.Errorf("duplicated id for testData")
	}
	if elem.Id >= testDataCount {
		return fmt.Errorf("testData id should be lower or equal than the last id")
	}
	testDataIdMap[elem.Id] = true
}
// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
