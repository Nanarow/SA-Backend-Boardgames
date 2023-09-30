package entity

type MemberBill struct {
	Bill         Bill `gorm:"embedded"`
	MemberTypeID uint
}
