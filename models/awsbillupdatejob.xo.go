// Package models contains the types for schema 'trackit'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// AwsBillUpdateJob represents a row from 'trackit.aws_bill_update_job'.
type AwsBillUpdateJob struct {
	ID                  int       `json:"id"`                     // id
	AwsBillRepositoryID int       `json:"aws_bill_repository_id"` // aws_bill_repository_id
	Expired             time.Time `json:"expired"`                // expired
	Completed           time.Time `json:"completed"`              // completed
	WorkerID            string    `json:"worker_id"`              // worker_id
	Error               string    `json:"error"`                  // error

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AwsBillUpdateJob exists in the database.
func (abuj *AwsBillUpdateJob) Exists() bool {
	return abuj._exists
}

// Deleted provides information if the AwsBillUpdateJob has been deleted from the database.
func (abuj *AwsBillUpdateJob) Deleted() bool {
	return abuj._deleted
}

// Insert inserts the AwsBillUpdateJob to the database.
func (abuj *AwsBillUpdateJob) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if abuj._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO trackit.aws_bill_update_job (` +
		`aws_bill_repository_id, expired, completed, worker_id, error` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, abuj.AwsBillRepositoryID, abuj.Expired, abuj.Completed, abuj.WorkerID, abuj.Error)
	res, err := db.Exec(sqlstr, abuj.AwsBillRepositoryID, abuj.Expired, abuj.Completed, abuj.WorkerID, abuj.Error)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	abuj.ID = int(id)
	abuj._exists = true

	return nil
}

// Update updates the AwsBillUpdateJob in the database.
func (abuj *AwsBillUpdateJob) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !abuj._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if abuj._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE trackit.aws_bill_update_job SET ` +
		`aws_bill_repository_id = ?, expired = ?, completed = ?, worker_id = ?, error = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, abuj.AwsBillRepositoryID, abuj.Expired, abuj.Completed, abuj.WorkerID, abuj.Error, abuj.ID)
	_, err = db.Exec(sqlstr, abuj.AwsBillRepositoryID, abuj.Expired, abuj.Completed, abuj.WorkerID, abuj.Error, abuj.ID)
	return err
}

// Save saves the AwsBillUpdateJob to the database.
func (abuj *AwsBillUpdateJob) Save(db XODB) error {
	if abuj.Exists() {
		return abuj.Update(db)
	}

	return abuj.Insert(db)
}

// Delete deletes the AwsBillUpdateJob from the database.
func (abuj *AwsBillUpdateJob) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !abuj._exists {
		return nil
	}

	// if deleted, bail
	if abuj._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM trackit.aws_bill_update_job WHERE id = ?`

	// run query
	XOLog(sqlstr, abuj.ID)
	_, err = db.Exec(sqlstr, abuj.ID)
	if err != nil {
		return err
	}

	// set deleted
	abuj._deleted = true

	return nil
}

// AwsBillRepository returns the AwsBillRepository associated with the AwsBillUpdateJob's AwsBillRepositoryID (aws_bill_repository_id).
//
// Generated from foreign key 'aws_bill_update_job_ibfk_1'.
func (abuj *AwsBillUpdateJob) AwsBillRepository(db XODB) (*AwsBillRepository, error) {
	return AwsBillRepositoryByID(db, abuj.AwsBillRepositoryID)
}

// AwsBillUpdateJobByID retrieves a row from 'trackit.aws_bill_update_job' as a AwsBillUpdateJob.
//
// Generated from index 'aws_bill_update_job_id_pkey'.
func AwsBillUpdateJobByID(db XODB, id int) (*AwsBillUpdateJob, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, aws_bill_repository_id, expired, completed, worker_id, error ` +
		`FROM trackit.aws_bill_update_job ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	abuj := AwsBillUpdateJob{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&abuj.ID, &abuj.AwsBillRepositoryID, &abuj.Expired, &abuj.Completed, &abuj.WorkerID, &abuj.Error)
	if err != nil {
		return nil, err
	}

	return &abuj, nil
}

// AwsBillUpdateJobsByAwsBillRepositoryID retrieves a row from 'trackit.aws_bill_update_job' as a AwsBillUpdateJob.
//
// Generated from index 'foreign_bill_repository'.
func AwsBillUpdateJobsByAwsBillRepositoryID(db XODB, awsBillRepositoryID int) ([]*AwsBillUpdateJob, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, aws_bill_repository_id, expired, completed, worker_id, error ` +
		`FROM trackit.aws_bill_update_job ` +
		`WHERE aws_bill_repository_id = ?`

	// run query
	XOLog(sqlstr, awsBillRepositoryID)
	q, err := db.Query(sqlstr, awsBillRepositoryID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AwsBillUpdateJob{}
	for q.Next() {
		abuj := AwsBillUpdateJob{
			_exists: true,
		}

		// scan
		err = q.Scan(&abuj.ID, &abuj.AwsBillRepositoryID, &abuj.Expired, &abuj.Completed, &abuj.WorkerID, &abuj.Error)
		if err != nil {
			return nil, err
		}

		res = append(res, &abuj)
	}

	return res, nil
}
