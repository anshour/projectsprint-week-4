package constants

import (
	"errors"

	"github.com/lib/pq"
)

var ErrNotFound = errors.New("no row affected")
var ErrConflict = errors.New("error conflict")
var ErrInternalServer = errors.New("errror checking row affected")
var ErrWrongPassword = errors.New("wrong password verification")
var ErrEmptyUserId = errors.New("Error Empty user id")
var ErrUsernameAlreadyExist = errors.New("username or email already exist")
var ErrMissingMerchantItem = errors.New("one of merchant or item not exist")
var ErrTooFarLocation = errors.New("One of item has too far location")
var ErrNoRowsResultText = "sql: no rows in result set"
var ErrErrorParsing = "Error parsing lat, long"
var ErrEmptyStringUUID = errors.New("uuid cant be empty string")
var ErrInvalidUUID = errors.New("invalid uuid type")
var ErrStartingPoint = "Need one starting point"
var ErrNoRowsResult = errors.New(ErrNoRowsResultText)
var ErrNoLatLongPointProvided = errors.New("No latitude and longitude points provided")

const (
	UniqueViolationExistData    = pq.ErrorCode("23505") // 'unique_violation'
	UniqueViolationNotExistData = pq.ErrorCode("23503") // 'schema_and_data_statement_mixing_not_supported'
)
