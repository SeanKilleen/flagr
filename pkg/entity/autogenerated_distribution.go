package entity

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set DistributionQuerySet

// DistributionQuerySet is an queryset type for Distribution
type DistributionQuerySet struct {
	db *gorm.DB
}

// NewDistributionQuerySet constructs new DistributionQuerySet
func NewDistributionQuerySet(db *gorm.DB) DistributionQuerySet {
	return DistributionQuerySet{
		db: db.Model(&Distribution{}),
	}
}

func (qs DistributionQuerySet) w(db *gorm.DB) DistributionQuerySet {
	return NewDistributionQuerySet(db)
}

// All is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) All(ret *[]Distribution) error {
	return qs.db.Find(ret).Error
}

// BitmapEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) BitmapEq(bitmap string) DistributionQuerySet {
	return qs.w(qs.db.Where("bitmap = ?", bitmap))
}

// BitmapIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) BitmapIn(bitmap string, bitmapRest ...string) DistributionQuerySet {
	iArgs := []interface{}{bitmap}
	for _, arg := range bitmapRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("bitmap IN (?)", iArgs))
}

// BitmapNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) BitmapNe(bitmap string) DistributionQuerySet {
	return qs.w(qs.db.Where("bitmap != ?", bitmap))
}

// BitmapNotIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) BitmapNotIn(bitmap string, bitmapRest ...string) DistributionQuerySet {
	iArgs := []interface{}{bitmap}
	for _, arg := range bitmapRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("bitmap NOT IN (?)", iArgs))
}

// Count is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// Create is an autogenerated method
// nolint: dupl
func (o *Distribution) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// CreatedAtEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) CreatedAtEq(createdAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("created_at = ?", createdAt))
}

// CreatedAtGt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) CreatedAtGt(createdAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("created_at > ?", createdAt))
}

// CreatedAtGte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) CreatedAtGte(createdAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("created_at >= ?", createdAt))
}

// CreatedAtLt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) CreatedAtLt(createdAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("created_at < ?", createdAt))
}

// CreatedAtLte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) CreatedAtLte(createdAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("created_at <= ?", createdAt))
}

// CreatedAtNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) CreatedAtNe(createdAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("created_at != ?", createdAt))
}

// Delete is an autogenerated method
// nolint: dupl
func (o *Distribution) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) Delete() error {
	return qs.db.Delete(Distribution{}).Error
}

// DeletedAtEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) DeletedAtEq(deletedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("deleted_at = ?", deletedAt))
}

// DeletedAtGt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) DeletedAtGt(deletedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("deleted_at > ?", deletedAt))
}

// DeletedAtGte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) DeletedAtGte(deletedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("deleted_at >= ?", deletedAt))
}

// DeletedAtIsNotNull is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) DeletedAtIsNotNull() DistributionQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NOT NULL"))
}

// DeletedAtIsNull is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) DeletedAtIsNull() DistributionQuerySet {
	return qs.w(qs.db.Where("deleted_at IS NULL"))
}

// DeletedAtLt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) DeletedAtLt(deletedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("deleted_at < ?", deletedAt))
}

// DeletedAtLte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) DeletedAtLte(deletedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("deleted_at <= ?", deletedAt))
}

// DeletedAtNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) DeletedAtNe(deletedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("deleted_at != ?", deletedAt))
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) GetUpdater() DistributionUpdater {
	return NewDistributionUpdater(qs.db)
}

// IDEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) IDEq(ID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("id = ?", ID))
}

// IDGt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) IDGt(ID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("id > ?", ID))
}

// IDGte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) IDGte(ID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("id >= ?", ID))
}

// IDIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) IDIn(ID uint, IDRest ...uint) DistributionQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id IN (?)", iArgs))
}

// IDLt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) IDLt(ID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("id < ?", ID))
}

// IDLte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) IDLte(ID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("id <= ?", ID))
}

// IDNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) IDNe(ID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("id != ?", ID))
}

// IDNotIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) IDNotIn(ID uint, IDRest ...uint) DistributionQuerySet {
	iArgs := []interface{}{ID}
	for _, arg := range IDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("id NOT IN (?)", iArgs))
}

// Limit is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) Limit(limit int) DistributionQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs DistributionQuerySet) One(ret *Distribution) error {
	return qs.db.First(ret).Error
}

// OrderAscByCreatedAt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderAscByCreatedAt() DistributionQuerySet {
	return qs.w(qs.db.Order("created_at ASC"))
}

// OrderAscByDeletedAt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderAscByDeletedAt() DistributionQuerySet {
	return qs.w(qs.db.Order("deleted_at ASC"))
}

// OrderAscByID is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderAscByID() DistributionQuerySet {
	return qs.w(qs.db.Order("id ASC"))
}

// OrderAscByPercent is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderAscByPercent() DistributionQuerySet {
	return qs.w(qs.db.Order("percent ASC"))
}

// OrderAscBySegmentID is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderAscBySegmentID() DistributionQuerySet {
	return qs.w(qs.db.Order("segment_id ASC"))
}

// OrderAscByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderAscByUpdatedAt() DistributionQuerySet {
	return qs.w(qs.db.Order("updated_at ASC"))
}

// OrderAscByVariantID is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderAscByVariantID() DistributionQuerySet {
	return qs.w(qs.db.Order("variant_id ASC"))
}

// OrderDescByCreatedAt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderDescByCreatedAt() DistributionQuerySet {
	return qs.w(qs.db.Order("created_at DESC"))
}

// OrderDescByDeletedAt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderDescByDeletedAt() DistributionQuerySet {
	return qs.w(qs.db.Order("deleted_at DESC"))
}

// OrderDescByID is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderDescByID() DistributionQuerySet {
	return qs.w(qs.db.Order("id DESC"))
}

// OrderDescByPercent is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderDescByPercent() DistributionQuerySet {
	return qs.w(qs.db.Order("percent DESC"))
}

// OrderDescBySegmentID is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderDescBySegmentID() DistributionQuerySet {
	return qs.w(qs.db.Order("segment_id DESC"))
}

// OrderDescByUpdatedAt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderDescByUpdatedAt() DistributionQuerySet {
	return qs.w(qs.db.Order("updated_at DESC"))
}

// OrderDescByVariantID is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) OrderDescByVariantID() DistributionQuerySet {
	return qs.w(qs.db.Order("variant_id DESC"))
}

// PercentEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) PercentEq(percent uint) DistributionQuerySet {
	return qs.w(qs.db.Where("percent = ?", percent))
}

// PercentGt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) PercentGt(percent uint) DistributionQuerySet {
	return qs.w(qs.db.Where("percent > ?", percent))
}

// PercentGte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) PercentGte(percent uint) DistributionQuerySet {
	return qs.w(qs.db.Where("percent >= ?", percent))
}

// PercentIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) PercentIn(percent uint, percentRest ...uint) DistributionQuerySet {
	iArgs := []interface{}{percent}
	for _, arg := range percentRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("percent IN (?)", iArgs))
}

// PercentLt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) PercentLt(percent uint) DistributionQuerySet {
	return qs.w(qs.db.Where("percent < ?", percent))
}

// PercentLte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) PercentLte(percent uint) DistributionQuerySet {
	return qs.w(qs.db.Where("percent <= ?", percent))
}

// PercentNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) PercentNe(percent uint) DistributionQuerySet {
	return qs.w(qs.db.Where("percent != ?", percent))
}

// PercentNotIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) PercentNotIn(percent uint, percentRest ...uint) DistributionQuerySet {
	iArgs := []interface{}{percent}
	for _, arg := range percentRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("percent NOT IN (?)", iArgs))
}

// SegmentIDEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) SegmentIDEq(segmentID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("segment_id = ?", segmentID))
}

// SegmentIDGt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) SegmentIDGt(segmentID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("segment_id > ?", segmentID))
}

// SegmentIDGte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) SegmentIDGte(segmentID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("segment_id >= ?", segmentID))
}

// SegmentIDIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) SegmentIDIn(segmentID uint, segmentIDRest ...uint) DistributionQuerySet {
	iArgs := []interface{}{segmentID}
	for _, arg := range segmentIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("segment_id IN (?)", iArgs))
}

// SegmentIDLt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) SegmentIDLt(segmentID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("segment_id < ?", segmentID))
}

// SegmentIDLte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) SegmentIDLte(segmentID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("segment_id <= ?", segmentID))
}

// SegmentIDNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) SegmentIDNe(segmentID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("segment_id != ?", segmentID))
}

// SegmentIDNotIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) SegmentIDNotIn(segmentID uint, segmentIDRest ...uint) DistributionQuerySet {
	iArgs := []interface{}{segmentID}
	for _, arg := range segmentIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("segment_id NOT IN (?)", iArgs))
}

// SetBitmap is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) SetBitmap(bitmap string) DistributionUpdater {
	u.fields[string(DistributionDBSchema.Bitmap)] = bitmap
	return u
}

// SetCreatedAt is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) SetCreatedAt(createdAt time.Time) DistributionUpdater {
	u.fields[string(DistributionDBSchema.CreatedAt)] = createdAt
	return u
}

// SetID is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) SetID(ID uint) DistributionUpdater {
	u.fields[string(DistributionDBSchema.ID)] = ID
	return u
}

// SetPercent is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) SetPercent(percent uint) DistributionUpdater {
	u.fields[string(DistributionDBSchema.Percent)] = percent
	return u
}

// SetSegmentID is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) SetSegmentID(segmentID uint) DistributionUpdater {
	u.fields[string(DistributionDBSchema.SegmentID)] = segmentID
	return u
}

// SetUpdatedAt is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) SetUpdatedAt(updatedAt time.Time) DistributionUpdater {
	u.fields[string(DistributionDBSchema.UpdatedAt)] = updatedAt
	return u
}

// SetVariantID is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) SetVariantID(variantID uint) DistributionUpdater {
	u.fields[string(DistributionDBSchema.VariantID)] = variantID
	return u
}

// SetVariantKey is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) SetVariantKey(variantKey string) DistributionUpdater {
	u.fields[string(DistributionDBSchema.VariantKey)] = variantKey
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u DistributionUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// UpdatedAtEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) UpdatedAtEq(updatedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("updated_at = ?", updatedAt))
}

// UpdatedAtGt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) UpdatedAtGt(updatedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("updated_at > ?", updatedAt))
}

// UpdatedAtGte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) UpdatedAtGte(updatedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("updated_at >= ?", updatedAt))
}

// UpdatedAtLt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) UpdatedAtLt(updatedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("updated_at < ?", updatedAt))
}

// UpdatedAtLte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) UpdatedAtLte(updatedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("updated_at <= ?", updatedAt))
}

// UpdatedAtNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) UpdatedAtNe(updatedAt time.Time) DistributionQuerySet {
	return qs.w(qs.db.Where("updated_at != ?", updatedAt))
}

// VariantIDEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantIDEq(variantID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("variant_id = ?", variantID))
}

// VariantIDGt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantIDGt(variantID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("variant_id > ?", variantID))
}

// VariantIDGte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantIDGte(variantID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("variant_id >= ?", variantID))
}

// VariantIDIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantIDIn(variantID uint, variantIDRest ...uint) DistributionQuerySet {
	iArgs := []interface{}{variantID}
	for _, arg := range variantIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("variant_id IN (?)", iArgs))
}

// VariantIDLt is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantIDLt(variantID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("variant_id < ?", variantID))
}

// VariantIDLte is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantIDLte(variantID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("variant_id <= ?", variantID))
}

// VariantIDNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantIDNe(variantID uint) DistributionQuerySet {
	return qs.w(qs.db.Where("variant_id != ?", variantID))
}

// VariantIDNotIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantIDNotIn(variantID uint, variantIDRest ...uint) DistributionQuerySet {
	iArgs := []interface{}{variantID}
	for _, arg := range variantIDRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("variant_id NOT IN (?)", iArgs))
}

// VariantKeyEq is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantKeyEq(variantKey string) DistributionQuerySet {
	return qs.w(qs.db.Where("variant_key = ?", variantKey))
}

// VariantKeyIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantKeyIn(variantKey string, variantKeyRest ...string) DistributionQuerySet {
	iArgs := []interface{}{variantKey}
	for _, arg := range variantKeyRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("variant_key IN (?)", iArgs))
}

// VariantKeyNe is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantKeyNe(variantKey string) DistributionQuerySet {
	return qs.w(qs.db.Where("variant_key != ?", variantKey))
}

// VariantKeyNotIn is an autogenerated method
// nolint: dupl
func (qs DistributionQuerySet) VariantKeyNotIn(variantKey string, variantKeyRest ...string) DistributionQuerySet {
	iArgs := []interface{}{variantKey}
	for _, arg := range variantKeyRest {
		iArgs = append(iArgs, arg)
	}
	return qs.w(qs.db.Where("variant_key NOT IN (?)", iArgs))
}

// ===== END of query set DistributionQuerySet

// ===== BEGIN of Distribution modifiers

type distributionDBSchemaField string

// DistributionDBSchema stores db field names of Distribution
var DistributionDBSchema = struct {
	ID         distributionDBSchemaField
	CreatedAt  distributionDBSchemaField
	UpdatedAt  distributionDBSchemaField
	DeletedAt  distributionDBSchemaField
	SegmentID  distributionDBSchemaField
	VariantID  distributionDBSchemaField
	VariantKey distributionDBSchemaField
	Percent    distributionDBSchemaField
	Bitmap     distributionDBSchemaField
}{

	ID:         distributionDBSchemaField("id"),
	CreatedAt:  distributionDBSchemaField("created_at"),
	UpdatedAt:  distributionDBSchemaField("updated_at"),
	DeletedAt:  distributionDBSchemaField("deleted_at"),
	SegmentID:  distributionDBSchemaField("segment_id"),
	VariantID:  distributionDBSchemaField("variant_id"),
	VariantKey: distributionDBSchemaField("variant_key"),
	Percent:    distributionDBSchemaField("percent"),
	Bitmap:     distributionDBSchemaField("bitmap"),
}

// Update updates Distribution fields by primary key
func (o *Distribution) Update(db *gorm.DB, fields ...distributionDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"id":          o.ID,
		"created_at":  o.CreatedAt,
		"updated_at":  o.UpdatedAt,
		"deleted_at":  o.DeletedAt,
		"segment_id":  o.SegmentID,
		"variant_id":  o.VariantID,
		"variant_key": o.VariantKey,
		"percent":     o.Percent,
		"bitmap":      o.Bitmap,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := string(f)
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update Distribution %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// DistributionUpdater is an Distribution updates manager
type DistributionUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewDistributionUpdater creates new Distribution updater
func NewDistributionUpdater(db *gorm.DB) DistributionUpdater {
	return DistributionUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&Distribution{}),
	}
}

// ===== END of Distribution modifiers

// ===== END of all query sets
