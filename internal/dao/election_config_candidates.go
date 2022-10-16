// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"election/internal/dao/internal"
)

// internalElectionConfigCandidatesDao is internal type for wrapping internal DAO implements.
type internalElectionConfigCandidatesDao = *internal.ElectionConfigCandidatesDao

// electionConfigCandidatesDao is the data access object for table election_config_candidates.
// You can define custom methods on it to extend its functionality as you wish.
type electionConfigCandidatesDao struct {
	internalElectionConfigCandidatesDao
}

var (
	// ElectionConfigCandidates is globally public accessible object for table election_config_candidates operations.
	ElectionConfigCandidates = electionConfigCandidatesDao{
		internal.NewElectionConfigCandidatesDao(),
	}
)

// Fill with you ideas below.
