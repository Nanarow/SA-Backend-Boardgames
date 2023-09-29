package entity

type MemberBill struct {
	// Bill Bill `gorm:"embedded"`
	BillID *uint `gorm:"primaryKey"`

	MemberTypeID *uint
	MemberType   MemberType `gorm:"foreignKey:MemberTypeID"`
}
