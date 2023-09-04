package model

import "github.com/Kitsuya0828/gqlgen-ent-clean-architecture-boilerplate/ent"

// User is the model entity for the User schema.
type User = ent.User

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput = ent.CreateUserInput

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput = ent.UpdateUserInput

// UserConnection is the connection containing edges to User.
type UserConnection = ent.UserConnection

// UserOrder is the order direction for User.
type UserOrder = ent.UserOrder

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput = ent.UserWhereInput
