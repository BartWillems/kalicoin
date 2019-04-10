package models

import (
	"errors"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
)

// Wallet is a chat member's wallet 🤔
type Wallet struct {
	ID        int64     `json:"id" db:"id"`
	OwnerID   int       `json:"owner_id" db:"owner_id" uri:"owner_id" binding:"required"`
	GroupID   int64     `json:"group_id" db:"group_id" uri:"group_id" binding:"required"`
	Capital   uint32    `json:"capital" db:"capital"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// StarterCapital is the capital that a new wallet starts with
const StarterCapital uint32 = 100

// Wallets is a collection of wallets
type Wallets []Wallet

// Create will create a user's wallet with the default StarterCapital
func (w *Wallet) Create(tx *pop.Connection) error {
	w.Capital = StarterCapital

	return tx.Create(w)
}

// Get fetches the user's wallet and creates it if it doesn't exist
func (w *Wallet) Get(tx *pop.Connection, GroupID int64, UserID nulls.Int) error {
	if !UserID.Valid {
		return errors.New("Empty UserID provided")
	}

	return tx.Where("group_id = ?", GroupID).
		Where("owner_id = ?", UserID).
		First(w)
}

func (w *Wallet) pay(tx *pop.Connection, amount uint32) error {
	return tx.RawQuery("UPDATE wallets SET capital = capital - ? WHERE group_id = ? AND owner_id = ? RETURNING capital", amount, w.GroupID, w.OwnerID).Exec()
}

func (w *Wallet) reward(tx *pop.Connection, amount uint32) error {
	return tx.RawQuery("UPDATE wallets SET capital = capital + ? WHERE group_id = ? AND owner_id = ? RETURNING capital", amount, w.GroupID, w.OwnerID).Exec()
}
