package repository

import (
	"flashlight-go/internal/models"

	"gorm.io/gorm"
)

// Simple repositories using base repository

type MembershipTypeRepository struct {
	*BaseRepository[models.MembershipType]
}

func NewMembershipTypeRepository(db *gorm.DB) *MembershipTypeRepository {
	return &MembershipTypeRepository{
		BaseRepository: NewBaseRepository[models.MembershipType](db),
	}
}

type DeviceFCMTokenRepository struct {
	*BaseRepository[models.DeviceFCMToken]
}

func NewDeviceFCMTokenRepository(db *gorm.DB) *DeviceFCMTokenRepository {
	return &DeviceFCMTokenRepository{
		BaseRepository: NewBaseRepository[models.DeviceFCMToken](db),
	}
}

type VehicleRepository struct {
	*BaseRepository[models.Vehicle]
}

func NewVehicleRepository(db *gorm.DB) *VehicleRepository {
	return &VehicleRepository{
		BaseRepository: NewBaseRepository[models.Vehicle](db),
	}
}

type CustomerVehicleRepository struct {
	*BaseRepository[models.CustomerVehicle]
}

func NewCustomerVehicleRepository(db *gorm.DB) *CustomerVehicleRepository {
	return &CustomerVehicleRepository{
		BaseRepository: NewBaseRepository[models.CustomerVehicle](db),
	}
}

type ProductCategoryRepository struct {
	*BaseRepository[models.ProductCategory]
}

func NewProductCategoryRepository(db *gorm.DB) *ProductCategoryRepository {
	return &ProductCategoryRepository{
		BaseRepository: NewBaseRepository[models.ProductCategory](db),
	}
}

type ProductRepository struct {
	*BaseRepository[models.Product]
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		BaseRepository: NewBaseRepository[models.Product](db),
	}
}

type WorkOrderItemRepository struct {
	*BaseRepository[models.WorkOrderItem]
}

func NewWorkOrderItemRepository(db *gorm.DB) *WorkOrderItemRepository {
	return &WorkOrderItemRepository{
		BaseRepository: NewBaseRepository[models.WorkOrderItem](db),
	}
}
