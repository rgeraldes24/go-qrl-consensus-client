// Copyright © 2022, 2023 Attestant Limited.
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

//nolint:revive
// Need to `go install github.com/ferranbt/fastssz/sszgen@latest` for this to work.
//go:generate rm -f beaconblockbody_ssz.go beaconblock_ssz.go beaconstate_ssz.go executionpayloadheader_ssz.go executionpayload_ssz.go historicalsummary_ssz.go signedbeaconblock_ssz.go withdrawal_ssz.go
//go:generate sszgen -suffix ssz -include ../phase0,../altair,../bellatrix -path . -objs BeaconBlockBody,BeaconBlock,BeaconState,ExecutionPayload,ExecutionPayloadHeader,HistoricalSummary,SignedBeaconBlock,Withdrawal
//go:generate goimports -w beaconblockbody_ssz.go beaconblock_ssz.go beaconstate_ssz.go executionpayloadheader_ssz.go executionpayload_ssz.go historicalsummary_ssz.go signedbeaconblock_ssz.go withdrawal_ssz.go

// ALTAIR
//go:generate rm -f beaconblock_ssz.go beaconblockbody_ssz.go beaconstate_ssz.go contributionandproof_ssz.go signedbeaconblock_ssz.go signedcontributionandproof_ssz.go syncaggregate_ssz.go syncaggregatorselectiondata_ssz.go synccommittee_ssz.go synccommitteecontribution_ssz.go synccommitteemessage_ssz.go
//go:generate sszgen -suffix ssz -include ../phase0 -path . -objs BeaconBlock,BeaconBlockBody,BeaconState,ContributionAndProof,SignedBeaconBlock,SignedContributionAndProof,SyncAggregate,SyncAggregatorSelectionData,SyncCommittee,SyncCommitteeContribution,SyncCommitteeMessage
//go:generate goimports -w beaconblock_ssz.go beaconblockbody_ssz.go beaconstate_ssz.go contributionandproof_ssz.go signedbeaconblock_ssz.go signedcontributionandproof_ssz.go syncaggregate_ssz.go syncaggregatorselectiondata_ssz.go synccommitteecontribution_ssz.go synccommitteemessage_ssz.go

// BELLATRIX
//go:generate rm -f beaconblock_ssz.go beaconblockbody_ssz.go beaconstate_ssz.go executionpayload_ssz.go executionpayloadheader_ssz.go signedbeaconblock_ssz.go
//go:generate sszgen -suffix ssz -include ../phase0,../altair -path . -objs BeaconBlock,BeaconBlockBody,BeaconState,ExecutionPayload,ExecutionPaylodHeader,SignedBeaconBlock
//go:generate goimports -w beaconblock_ssz.go beaconblockbody_ssz.go beaconstate_ssz.go executionpayload_ssz.go executionpayloadheader_ssz.go signedbeaconblock_ssz.go

// PHASE0
//go:generate rm -f aggregateandproof_ssz.go attestationdata_ssz.go attestation_ssz.go attesterslashing_ssz.go beaconblockbody_ssz.go beaconblock_ssz.go beaconblockheader_ssz.go beaconstate_ssz.go checkpoint_ssz.go depositdata_ssz.go deposit_ssz.go depositmessage_ssz.go executiondata_ssz.go forkdata_ssz.go fork_ssz.go indexedattestation_ssz.go pendingattestation_ssz.go proposerslashing_ssz.go signedaggregateandproof_ssz.go signedbeaconblock_ssz.go signedbeaconblockheader_ssz.go signedvoluntaryexit_ssz.go signingdata_ssz.go validator_ssz.go voluntaryexit_ssz.go
//go:generate sszgen -suffix ssz -path . --objs AggregateAndProof,AttestationData,Attestation,AttesterSlashing,BeaconBlockBody,BeaconBlock,BeaconBlockHeader,BeaconState,Checkpoint,Deposit,DepositData,DepositMessage,ExecutionData,Fork,ForkData,IndexedAttestation,PendingAttestation,ProposerSlashing,SignedAggregateAndProof,SignedBeaconBlock,SignedBeaconBlockHeader,SignedVoluntaryExit,SigningData,Validator,VoluntaryExit
//go:generate goimports -w aggregateandproof_ssz.go attestationdata_ssz.go attestation_ssz.go attesterslashing_ssz.go beaconblockbody_ssz.go beaconblock_ssz.go beaconblockheader_ssz.go beaconstate_ssz.go checkpoint_ssz.go depositdata_ssz.go deposit_ssz.go depositmessage_ssz.go executiondata_ssz.go forkdata_ssz.go fork_ssz.go indexedattestation_ssz.go pendingattestation_ssz.go proposerslashing_ssz.go signedaggregateandproof_ssz.go signedbeaconblock_ssz.go signedbeaconblockheader_ssz.go signedvoluntaryexit_ssz.go signingdata_ssz.go validator_ssz.go voluntaryexit_ssz.go
