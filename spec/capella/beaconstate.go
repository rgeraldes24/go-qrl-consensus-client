// Copyright © 2022 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package capella

import (
	"fmt"

	"github.com/goccy/go-yaml"
	bitfield "github.com/theQRL/go-bitfield"
	"github.com/theQRL/go-qrl-consensus-client/spec/capella"
)

// BeaconState represents a beacon state.
type BeaconState struct {
	GenesisTime                  uint64
	GenesisValidatorsRoot        capella.Root `ssz-size:"32"`
	Slot                         capella.Slot
	Fork                         *capella.Fork
	LatestBlockHeader            *capella.BeaconBlockHeader
	BlockRoots                   []capella.Root `dynssz-size:"SLOTS_PER_HISTORICAL_ROOT,32" ssz-size:"8192,32"`
	StateRoots                   []capella.Root `dynssz-size:"SLOTS_PER_HISTORICAL_ROOT,32" ssz-size:"8192,32"`
	HistoricalRoots              []capella.Root `dynssz-max:"HISTORICAL_ROOTS_LIMIT"        ssz-max:"16777216" ssz-size:"?,32"`
	ExecutionData                *capella.ExecutionData
	ExecutionDataVotes           []*capella.ExecutionData `dynssz-max:"EPOCHS_PER_ETH1_VOTING_PERIOD*SLOTS_PER_EPOCH" ssz-max:"2048"`
	ETH1DepositIndex             uint64
	Validators                   []*capella.Validator         `dynssz-max:"VALIDATOR_REGISTRY_LIMIT"         ssz-max:"1099511627776"`
	Balances                     []capella.Gwei               `dynssz-max:"VALIDATOR_REGISTRY_LIMIT"         ssz-max:"1099511627776"`
	RANDAOMixes                  []capella.Root               `dynssz-size:"EPOCHS_PER_HISTORICAL_VECTOR,32" ssz-size:"65536,32"`
	Slashings                    []capella.Gwei               `dynssz-size:"EPOCHS_PER_SLASHINGS_VECTOR"     ssz-size:"8192"`
	PreviousEpochParticipation   []capella.ParticipationFlags `dynssz-max:"VALIDATOR_REGISTRY_LIMIT"         ssz-max:"1099511627776"`
	CurrentEpochParticipation    []capella.ParticipationFlags `dynssz-max:"VALIDATOR_REGISTRY_LIMIT"         ssz-max:"1099511627776"`
	JustificationBits            bitfield.Bitvector4          `ssz-size:"1"`
	PreviousJustifiedCheckpoint  *capella.Checkpoint
	CurrentJustifiedCheckpoint   *capella.Checkpoint
	FinalizedCheckpoint          *capella.Checkpoint
	InactivityScores             []uint64 `dynssz-max:"VALIDATOR_REGISTRY_LIMIT" ssz-max:"1099511627776"`
	CurrentSyncCommittee         *capella.SyncCommittee
	NextSyncCommittee            *capella.SyncCommittee
	LatestExecutionPayloadHeader *ExecutionPayloadHeader
	NextWithdrawalIndex          WithdrawalIndex
	NextWithdrawalValidatorIndex capella.ValidatorIndex
	HistoricalSummaries          []*HistoricalSummary `dynssz-max:"HISTORICAL_ROOTS_LIMIT" ssz-max:"16777216"`
}

// String returns a string version of the structure.
func (s *BeaconState) String() string {
	data, err := yaml.Marshal(s)
	if err != nil {
		return fmt.Sprintf("ERR: %v", err)
	}

	return string(data)
}
