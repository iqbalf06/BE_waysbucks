package authdto

type RegisterResponse struct {
	Fullname  string `gorm:"type: varchar(255)" json:"fullname" validate:"required"`
}

type LoginResponse struct {
	ID       int    `json:"id"`
	Fullname  string `gorm:"type: varchar(255)" json:"fullname"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Token string `gorm:"type: varchar(255)" json:"token"`
	Role string `gorm:"type: varchar(50)" json:"role"`
}

type CheckAuthResponse struct {
	ID int `gorm:"type: int" json:"id"`
	Fullname string `gorm:"type:varchar(255)" json:"name"`
	Email string `gorm:"type: varchar(255)" json:"email"`
	Role string `gorm:"type: varchar(50)" json:"role"`
}
