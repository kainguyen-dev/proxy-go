package data

import (
	"errors"
	"gorm.io/gorm"
	"svc/proxy-service/internal/common"
	"time"
)

type Customer struct {
	CustomerID          string    `gorm:"column:customer_id;type:varchar;not null"`
	Name                string    `gorm:"column:name;type:varchar;not null"`
	AdminFirstName      *string   `gorm:"column:admin_first_name;type:varchar"`
	AdminLastName       *string   `gorm:"column:admin_last_name;type:varchar"`
	CreatedAt           time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	UpdatedAt           time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	PrimaryContactEmail *string   `gorm:"column:primary_contact_email;type:varchar"`
	Address             string    `gorm:"column:address;type:varchar;not null"`
	NumberOfSites       int       `gorm:"column:number_of_sites;type:integer;default:0;not null"`
	Phone               *string   `gorm:"column:phone;type:varchar"`
	StartDate           time.Time `gorm:"column:start_date;type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	City                string    `gorm:"column:city;type:varchar;not null"`
	Country             string    `gorm:"column:country;type:varchar;not null"`
	StateProvince       string    `gorm:"column:state_province;type:varchar;not null"`
	ZipCode             string    `gorm:"column:zip_code;type:varchar;not null"`
	Deleted             bool      `gorm:"column:deleted;type:boolean;default:false;not null"`
	CustomerUniqueID    int       `gorm:"column:customer_unique_id;type:integer;default:nextval('customer_id_seq'::regclass);primaryKey;not null"`
	SubscriptionPlan    *string   `gorm:"column:subscription_plan;type:varchar(100)"`
}

func (Customer) TableName() string {
	return "customer"
}

func FindCustomer(condition interface{}) Customer {
	db := DB
	var model Customer
	err := db.Where(condition).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(common.ClientError{Code: 404, Message: "Customer not found"})
	}
	return model
}

func FindCustomers(condition interface{}) []Customer {
	db := DB
	var model []Customer
	err := db.Where(condition).Find(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(common.ClientError{Code: 404, Message: "Customer not found"})
	}
	return model
}
