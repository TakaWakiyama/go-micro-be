package common

type PrimaryKey = string

type UserInput struct {
	PK           *string
	Email        *string
	Name         *string
	AvatarSource *string
}
