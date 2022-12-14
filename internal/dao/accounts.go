// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"election/internal/dao/internal"
)

// internalAccountsDao is internal type for wrapping internal DAO implements.
type internalAccountsDao = *internal.AccountsDao

// accountsDao is the data access object for table accounts.
// You can define custom methods on it to extend its functionality as you wish.
type accountsDao struct {
	internalAccountsDao
}

var (
	// Accounts is globally public accessible object for table accounts operations.
	Accounts = accountsDao{
		internal.NewAccountsDao(),
	}
)

// Fill with you ideas below.


