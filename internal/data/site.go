package data

import (
	"errors"
	"gorm.io/gorm"
	"svc/proxy-service/internal/common"
	"time"
)

type Site struct {
	SiteID                   int       `gorm:"primaryKey;autoIncrement;column:site_unique_id"`
	CustomerID               int       `gorm:"column:customer_unique_id"`
	SiteName                 string    `gorm:"column:site_name"`
	SiteGroup                string    `gorm:"column:site_group"`
	PoDate                   time.Time `gorm:"column:po_date"`
	PoNumber                 string    `gorm:"column:po_number"`
	ShippingDate             time.Time `gorm:"column:shipping_date"`
	TerastorPartNumber       string    `gorm:"column:terastor_part_number"`
	SiteAddr1                string    `gorm:"column:siteaddr1"`
	SiteAddr2                string    `gorm:"column:siteaddr2"`
	SiteCity                 string    `gorm:"column:sitecity;default:Hudson"`
	SiteStateProvince        string    `gorm:"column:sitestateprovince;default:MA"`
	SiteCountry              string    `gorm:"column:sitecountry"`
	SitePostalCode           string    `gorm:"column:sitepostalcode"`
	SiteLatitude             float64   `gorm:"column:sitelatitude"`
	SiteLongitude            float64   `gorm:"column:sitelongitude"`
	SiteElevation            float64   `gorm:"column:siteelevation;type:numeric(10,2)"`
	SiteManager              string    `gorm:"column:sitemanager"`
	SiteManagerContactNumber string    `gorm:"column:sitemanagercontactnumber"`
	SiteDescription          string    `gorm:"column:sitedescription;type:text"`
	CommissionDate           time.Time `gorm:"column:commissiondate;default:CURRENT_DATE"`
	NumTerastors             int       `gorm:"column:numterastors"`
	SiteAcVoltage            float64   `gorm:"column:siteacvoltage;type:numeric(10,2)"`
	SiteRatedPower           []int64   `gorm:"column:siteratedpower;type:integer[]"`
	SiteRatedReactivePower   float64   `gorm:"column:siteratedreactivepower;type:numeric(10,2)"`
	SiteRatedCapacityKwh     []int64   `gorm:"column:siteratedcapacitykwh;type:integer[]"`
	CreatedAt                time.Time `gorm:"column:created_at"`
	Ds                       time.Time `gorm:"column:ds"`
	Deleted                  bool      `gorm:"column:deleted;default:false;not null"`
	PrimaryContactEmail      string    `gorm:"column:primarycontactemail"`
	EmsVersionID             int       `gorm:"column:ems_version_id"`
	FirmwareEmsUpdatedAt     time.Time `gorm:"column:firmware_ems_updated_at;type:timestamp with time zone"`
	BmsVersionID             int       `gorm:"column:bms_version_id"`
	FirmwareBmsUpdatedAt     time.Time `gorm:"column:firmware_bms_updated_at;type:timestamp with time zone"`
}

func (Site) TableName() string {
	return "site_specifications"
}

func FindSite(condition interface{}) Site {
	db := DB
	var model Site
	err := db.Where(condition).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(common.ClientError{Code: 404, Message: "Site not found"})
	}
	return model
}

func FindSites(condition interface{}) []Site {
	db := DB
	var model []Site
	err := db.Where(condition).Find(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(common.ClientError{Code: 404, Message: "Sites not found"})
	}
	return model
}
