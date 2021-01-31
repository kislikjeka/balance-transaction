package domain

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid entity")

//ErrInvalidBalance invalid balance value
var ErrInvalidBalance= errors.New("Invalid balance")