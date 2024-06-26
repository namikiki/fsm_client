// Code generated by ent, DO NOT EDIT.

package file

import (
	"fsm_client/pkg/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.File {
	return predicate.File(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.File {
	return predicate.File(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldID, id))
}

// SyncID applies equality check predicate on the "sync_id" field. It's identical to SyncIDEQ.
func SyncID(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSyncID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldName, v))
}

// ParentDirID applies equality check predicate on the "parent_dir_id" field. It's identical to ParentDirIDEQ.
func ParentDirID(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldParentDirID, v))
}

// Level applies equality check predicate on the "level" field. It's identical to LevelEQ.
func Level(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldLevel, v))
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldHash, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSize, v))
}

// Deleted applies equality check predicate on the "deleted" field. It's identical to DeletedEQ.
func Deleted(v bool) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDeleted, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v int64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreateTime, v))
}

// ModTime applies equality check predicate on the "mod_time" field. It's identical to ModTimeEQ.
func ModTime(v int64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldModTime, v))
}

// SyncIDEQ applies the EQ predicate on the "sync_id" field.
func SyncIDEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSyncID, v))
}

// SyncIDNEQ applies the NEQ predicate on the "sync_id" field.
func SyncIDNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldSyncID, v))
}

// SyncIDIn applies the In predicate on the "sync_id" field.
func SyncIDIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldSyncID, vs...))
}

// SyncIDNotIn applies the NotIn predicate on the "sync_id" field.
func SyncIDNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldSyncID, vs...))
}

// SyncIDGT applies the GT predicate on the "sync_id" field.
func SyncIDGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldSyncID, v))
}

// SyncIDGTE applies the GTE predicate on the "sync_id" field.
func SyncIDGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldSyncID, v))
}

// SyncIDLT applies the LT predicate on the "sync_id" field.
func SyncIDLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldSyncID, v))
}

// SyncIDLTE applies the LTE predicate on the "sync_id" field.
func SyncIDLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldSyncID, v))
}

// SyncIDContains applies the Contains predicate on the "sync_id" field.
func SyncIDContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldSyncID, v))
}

// SyncIDHasPrefix applies the HasPrefix predicate on the "sync_id" field.
func SyncIDHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldSyncID, v))
}

// SyncIDHasSuffix applies the HasSuffix predicate on the "sync_id" field.
func SyncIDHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldSyncID, v))
}

// SyncIDEqualFold applies the EqualFold predicate on the "sync_id" field.
func SyncIDEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldSyncID, v))
}

// SyncIDContainsFold applies the ContainsFold predicate on the "sync_id" field.
func SyncIDContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldSyncID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldName, v))
}

// ParentDirIDEQ applies the EQ predicate on the "parent_dir_id" field.
func ParentDirIDEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldParentDirID, v))
}

// ParentDirIDNEQ applies the NEQ predicate on the "parent_dir_id" field.
func ParentDirIDNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldParentDirID, v))
}

// ParentDirIDIn applies the In predicate on the "parent_dir_id" field.
func ParentDirIDIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldParentDirID, vs...))
}

// ParentDirIDNotIn applies the NotIn predicate on the "parent_dir_id" field.
func ParentDirIDNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldParentDirID, vs...))
}

// ParentDirIDGT applies the GT predicate on the "parent_dir_id" field.
func ParentDirIDGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldParentDirID, v))
}

// ParentDirIDGTE applies the GTE predicate on the "parent_dir_id" field.
func ParentDirIDGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldParentDirID, v))
}

// ParentDirIDLT applies the LT predicate on the "parent_dir_id" field.
func ParentDirIDLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldParentDirID, v))
}

// ParentDirIDLTE applies the LTE predicate on the "parent_dir_id" field.
func ParentDirIDLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldParentDirID, v))
}

// ParentDirIDContains applies the Contains predicate on the "parent_dir_id" field.
func ParentDirIDContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldParentDirID, v))
}

// ParentDirIDHasPrefix applies the HasPrefix predicate on the "parent_dir_id" field.
func ParentDirIDHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldParentDirID, v))
}

// ParentDirIDHasSuffix applies the HasSuffix predicate on the "parent_dir_id" field.
func ParentDirIDHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldParentDirID, v))
}

// ParentDirIDEqualFold applies the EqualFold predicate on the "parent_dir_id" field.
func ParentDirIDEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldParentDirID, v))
}

// ParentDirIDContainsFold applies the ContainsFold predicate on the "parent_dir_id" field.
func ParentDirIDContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldParentDirID, v))
}

// LevelEQ applies the EQ predicate on the "level" field.
func LevelEQ(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldLevel, v))
}

// LevelNEQ applies the NEQ predicate on the "level" field.
func LevelNEQ(v int) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldLevel, v))
}

// LevelIn applies the In predicate on the "level" field.
func LevelIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldIn(FieldLevel, vs...))
}

// LevelNotIn applies the NotIn predicate on the "level" field.
func LevelNotIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldLevel, vs...))
}

// LevelGT applies the GT predicate on the "level" field.
func LevelGT(v int) predicate.File {
	return predicate.File(sql.FieldGT(FieldLevel, v))
}

// LevelGTE applies the GTE predicate on the "level" field.
func LevelGTE(v int) predicate.File {
	return predicate.File(sql.FieldGTE(FieldLevel, v))
}

// LevelLT applies the LT predicate on the "level" field.
func LevelLT(v int) predicate.File {
	return predicate.File(sql.FieldLT(FieldLevel, v))
}

// LevelLTE applies the LTE predicate on the "level" field.
func LevelLTE(v int) predicate.File {
	return predicate.File(sql.FieldLTE(FieldLevel, v))
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldHash, v))
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldHash, v))
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldHash, vs...))
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldHash, vs...))
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldHash, v))
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldHash, v))
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldHash, v))
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldHash, v))
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldHash, v))
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldHash, v))
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldHash, v))
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldHash, v))
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldHash, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int64) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int64) predicate.File {
	return predicate.File(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int64) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int64) predicate.File {
	return predicate.File(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int64) predicate.File {
	return predicate.File(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int64) predicate.File {
	return predicate.File(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int64) predicate.File {
	return predicate.File(sql.FieldLTE(FieldSize, v))
}

// DeletedEQ applies the EQ predicate on the "deleted" field.
func DeletedEQ(v bool) predicate.File {
	return predicate.File(sql.FieldEQ(FieldDeleted, v))
}

// DeletedNEQ applies the NEQ predicate on the "deleted" field.
func DeletedNEQ(v bool) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldDeleted, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v int64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v int64) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...int64) predicate.File {
	return predicate.File(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...int64) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v int64) predicate.File {
	return predicate.File(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v int64) predicate.File {
	return predicate.File(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v int64) predicate.File {
	return predicate.File(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v int64) predicate.File {
	return predicate.File(sql.FieldLTE(FieldCreateTime, v))
}

// ModTimeEQ applies the EQ predicate on the "mod_time" field.
func ModTimeEQ(v int64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldModTime, v))
}

// ModTimeNEQ applies the NEQ predicate on the "mod_time" field.
func ModTimeNEQ(v int64) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldModTime, v))
}

// ModTimeIn applies the In predicate on the "mod_time" field.
func ModTimeIn(vs ...int64) predicate.File {
	return predicate.File(sql.FieldIn(FieldModTime, vs...))
}

// ModTimeNotIn applies the NotIn predicate on the "mod_time" field.
func ModTimeNotIn(vs ...int64) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldModTime, vs...))
}

// ModTimeGT applies the GT predicate on the "mod_time" field.
func ModTimeGT(v int64) predicate.File {
	return predicate.File(sql.FieldGT(FieldModTime, v))
}

// ModTimeGTE applies the GTE predicate on the "mod_time" field.
func ModTimeGTE(v int64) predicate.File {
	return predicate.File(sql.FieldGTE(FieldModTime, v))
}

// ModTimeLT applies the LT predicate on the "mod_time" field.
func ModTimeLT(v int64) predicate.File {
	return predicate.File(sql.FieldLT(FieldModTime, v))
}

// ModTimeLTE applies the LTE predicate on the "mod_time" field.
func ModTimeLTE(v int64) predicate.File {
	return predicate.File(sql.FieldLTE(FieldModTime, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		p(s.Not())
	})
}
