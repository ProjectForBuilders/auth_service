package models

import (
	"time"
)

type User struct {
	ID           string  `json:"id"`
	Email        string  `json:"email"`
	Phone        *string `json:"phone"`
	PasswordHash string  `json:"-"`

	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	AvatarUrl *string `json:"avatar_url"`

	IsActive   bool       `json:"is_active"`
	IsVerified bool       `json:"is_verified"`
	LastLogin  *time.Time `json:"last_login"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Company struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	LegalName *string `json:"legal_name"`
	INN       *string `json:"inn"`
	OGRN      *string `json:"ogrn"`
	KPP       *string `json:"kpp"`
	Address   *string `json:"address"`
	Website   *string `json:"website"`
	LogoUrl   *string `json:"logo_url"`

	ContactEmail *string `json:"contact_email"`
	ContactPhone *string `json:"contact_phone"`

	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Level       string    `json:"level"`
	IsSystem    bool      `json:"is_system"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserRole struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	CompanyID string `json:"company_id"`
	RoleID    int    `json:"role_id"`

	Position   *string `json:"position"`
	Department *string `json:"department"`
	IsPrimary  bool    `json:"is_primary"`

	Status    string     `json:"status"`
	InviteBy  *string    `json:"invite_by"`
	InvitedAt *time.Time `json:"invited_at"`
	JoinedAt  *time.Time `json:"joined_at"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User    *User    `json:"user"`
	Company *Company `json:"company"`
	Role    *Role    `json:"role"`
}

type Permission struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Category    *string   `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
}

type RolePermission struct {
	RoleID       int `json:"role_id"`
	PermissionID int `json:"permission_id"`
}

type Invition struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	CompanyID string `json:"company_id"`
	RoleID    int    `json:"role_id"`
	InvitedBy string `json:"invited_by"`

	Status    string    `json:"status"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`

	Position   *string `json:"position"`
	Department *string `json:"department"`
	Message    *string `json:"message"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Session struct {
	ID           string  `json:"id"`
	UserID       string  `json:"user_id"`
	RefreshToken *string `json:"refresh_token"`
	UserAgent    *string `json:"user_agent"`
	IPAddress    *string `json:"ip_address"`
	DeviceInfo   *JSONB  `json:"device_info"`

	IsActive  bool  `json:"is_active"`
	ExpiresAt int64 `json:"expires_at"`

	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
}

type LoginHistory struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Email         string    `json:"email"`
	IPAddress     *string   `json:"ip_address"`
	UserAgent     *string   `json:"user_agent"`
	DeviceInfo    *JSONB    `json:"device_info"`
	Status        string    `json:"status"`
	FailureReason *string   `json:"failure_reason"`
	CreatedAt     time.Time `json:"created_at"`
}

type PasswordResetToken struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Token     string     `json:"token"`
	ExpiresAt time.Time  `json:"expires_at"`
	IsUsed    bool       `json:"is_used"`
	UsedAt    *time.Time `json:"used_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type TwoFactorAuth struct {
	UserID      string     `json:"user_id"`
	Secret      string     `json:"secret"`
	IsEnabled   bool       `json:"is_enabled"`
	BackupCodes *JSONB     `json:"backup_codes"`
	VerifiedAt  *time.Time `json:"verified_at"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type JSONB map[string]interface{}
