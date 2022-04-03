package model

type InternalRole string

const (
	RoleSuperAdmin InternalRole = "superadmin"
	RoleAdmin      InternalRole = "admin"
	RoleUser       InternalRole = "user"
)
