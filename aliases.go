package rocketchat

import (
	"github.com/davesavic/rocketchat/internal/data"
	"github.com/davesavic/rocketchat/internal/model"
)

// User related aliases
type CreateUser = data.CreateUser
type LoginUser = data.LoginUser
type User = model.User
type AuthenticatedUser = model.AuthenticatedUser

// Room related aliases
type Room = model.Room

// InstantMessage related aliases
type CreateInstantMessage = data.CreateInstantMessage
type ListInstantMessageMessages = data.ListInstantMessageMessages

// Message related aliases
type MessageList = model.MessageList
type Message = model.Message
