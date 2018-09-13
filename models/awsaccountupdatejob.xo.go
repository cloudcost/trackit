// Package models contains the types for schema 'trackit'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// AwsAccountUpdateJob represents a row from 'trackit.aws_account_update_job'.
type AwsAccountUpdateJob struct {
	ID              int       `json:"id"`              // id
	AwsAccountID    int       `json:"aws_account_id"`  // aws_account_id
	Completed       time.Time `json:"completed"`       // completed
	WorkerID        string    `json:"worker_id"`       // worker_id
	Joberror        string    `json:"jobError"`        // jobError
	Rdserror        string    `json:"rdsError"`        // rdsError
	Ec2error        string    `json:"ec2Error"`        // ec2Error
	Rdshistoryerror string    `json:"rdsHistoryError"` // rdsHistoryError
	Ec2historyerror string    `json:"ec2HistoryError"` // ec2HistoryError

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AwsAccountUpdateJob exists in the database.
func (aauj *AwsAccountUpdateJob) Exists() bool {
	return aauj._exists
}

// Deleted provides information if the AwsAccountUpdateJob has been deleted from the database.
func (aauj *AwsAccountUpdateJob) Deleted() bool {
	return aauj._deleted
}

// Insert inserts the AwsAccountUpdateJob to the database.
func (aauj *AwsAccountUpdateJob) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if aauj._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO trackit.aws_account_update_job (` +
		`aws_account_id, completed, worker_id, jobError, rdsError, ec2Error, rdsHistoryError, ec2HistoryError` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, aauj.AwsAccountID, aauj.Completed, aauj.WorkerID, aauj.Joberror, aauj.Rdserror, aauj.Ec2error, aauj.Rdshistoryerror, aauj.Ec2historyerror)
	res, err := db.Exec(sqlstr, aauj.AwsAccountID, aauj.Completed, aauj.WorkerID, aauj.Joberror, aauj.Rdserror, aauj.Ec2error, aauj.Rdshistoryerror, aauj.Ec2historyerror)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	aauj.ID = int(id)
	aauj._exists = true

	return nil
}

// Update updates the AwsAccountUpdateJob in the database.
func (aauj *AwsAccountUpdateJob) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !aauj._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if aauj._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE trackit.aws_account_update_job SET ` +
		`aws_account_id = ?, completed = ?, worker_id = ?, jobError = ?, rdsError = ?, ec2Error = ?, rdsHistoryError = ?, ec2HistoryError = ?` +
		` WHERE id = ?`

	// run query
	XOLog(sqlstr, aauj.AwsAccountID, aauj.Completed, aauj.WorkerID, aauj.Joberror, aauj.Rdserror, aauj.Ec2error, aauj.Rdshistoryerror, aauj.Ec2historyerror, aauj.ID)
	_, err = db.Exec(sqlstr, aauj.AwsAccountID, aauj.Completed, aauj.WorkerID, aauj.Joberror, aauj.Rdserror, aauj.Ec2error, aauj.Rdshistoryerror, aauj.Ec2historyerror, aauj.ID)
	return err
}

// Save saves the AwsAccountUpdateJob to the database.
func (aauj *AwsAccountUpdateJob) Save(db XODB) error {
	if aauj.Exists() {
		return aauj.Update(db)
	}

	return aauj.Insert(db)
}

// Delete deletes the AwsAccountUpdateJob from the database.
func (aauj *AwsAccountUpdateJob) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !aauj._exists {
		return nil
	}

	// if deleted, bail
	if aauj._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM trackit.aws_account_update_job WHERE id = ?`

	// run query
	XOLog(sqlstr, aauj.ID)
	_, err = db.Exec(sqlstr, aauj.ID)
	if err != nil {
		return err
	}

	// set deleted
	aauj._deleted = true

	return nil
}

// AwsAccount returns the AwsAccount associated with the AwsAccountUpdateJob's AwsAccountID (aws_account_id).
//
// Generated from foreign key 'aws_account_update_job_ibfk_1'.
func (aauj *AwsAccountUpdateJob) AwsAccount(db XODB) (*AwsAccount, error) {
	return AwsAccountByID(db, aauj.AwsAccountID)
}

// AwsAccountUpdateJobByID retrieves a row from 'trackit.aws_account_update_job' as a AwsAccountUpdateJob.
//
// Generated from index 'aws_account_update_job_id_pkey'.
func AwsAccountUpdateJobByID(db XODB, id int) (*AwsAccountUpdateJob, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, aws_account_id, completed, worker_id, jobError, rdsError, ec2Error, rdsHistoryError, ec2HistoryError ` +
		`FROM trackit.aws_account_update_job ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	aauj := AwsAccountUpdateJob{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&aauj.ID, &aauj.AwsAccountID, &aauj.Completed, &aauj.WorkerID, &aauj.Joberror, &aauj.Rdserror, &aauj.Ec2error, &aauj.Rdshistoryerror, &aauj.Ec2historyerror)
	if err != nil {
		return nil, err
	}

	return &aauj, nil
}

// AwsAccountUpdateJobsByAwsAccountID retrieves a row from 'trackit.aws_account_update_job' as a AwsAccountUpdateJob.
//
// Generated from index 'foreign_aws_account'.
func AwsAccountUpdateJobsByAwsAccountID(db XODB, awsAccountID int) ([]*AwsAccountUpdateJob, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, aws_account_id, completed, worker_id, jobError, rdsError, ec2Error, rdsHistoryError, ec2HistoryError ` +
		`FROM trackit.aws_account_update_job ` +
		`WHERE aws_account_id = ?`

	// run query
	XOLog(sqlstr, awsAccountID)
	q, err := db.Query(sqlstr, awsAccountID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*AwsAccountUpdateJob{}
	for q.Next() {
		aauj := AwsAccountUpdateJob{
			_exists: true,
		}

		// scan
		err = q.Scan(&aauj.ID, &aauj.AwsAccountID, &aauj.Completed, &aauj.WorkerID, &aauj.Joberror, &aauj.Rdserror, &aauj.Ec2error, &aauj.Rdshistoryerror, &aauj.Ec2historyerror)
		if err != nil {
			return nil, err
		}

		res = append(res, &aauj)
	}

	return res, nil
}
