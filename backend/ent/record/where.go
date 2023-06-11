// Code generated by ent, DO NOT EDIT.

package record

import (
	"code-connect/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldUUID, v))
}

// UserUUID applies equality check predicate on the "user_uuid" field. It's identical to UserUUIDEQ.
func UserUUID(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldUserUUID, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldCode, v))
}

// Readability applies equality check predicate on the "readability" field. It's identical to ReadabilityEQ.
func Readability(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldReadability, v))
}

// Modularity applies equality check predicate on the "modularity" field. It's identical to ModularityEQ.
func Modularity(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldModularity, v))
}

// Efficiency applies equality check predicate on the "efficiency" field. It's identical to EfficiencyEQ.
func Efficiency(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldEfficiency, v))
}

// Testability applies equality check predicate on the "testability" field. It's identical to TestabilityEQ.
func Testability(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldTestability, v))
}

// Maintainablity applies equality check predicate on the "maintainablity" field. It's identical to MaintainablityEQ.
func Maintainablity(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldMaintainablity, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Record {
	return predicate.Record(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Record {
	return predicate.Record(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Record {
	return predicate.Record(sql.FieldContainsFold(FieldUUID, v))
}

// UserUUIDEQ applies the EQ predicate on the "user_uuid" field.
func UserUUIDEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldUserUUID, v))
}

// UserUUIDNEQ applies the NEQ predicate on the "user_uuid" field.
func UserUUIDNEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldUserUUID, v))
}

// UserUUIDIn applies the In predicate on the "user_uuid" field.
func UserUUIDIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldUserUUID, vs...))
}

// UserUUIDNotIn applies the NotIn predicate on the "user_uuid" field.
func UserUUIDNotIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldUserUUID, vs...))
}

// UserUUIDGT applies the GT predicate on the "user_uuid" field.
func UserUUIDGT(v string) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldUserUUID, v))
}

// UserUUIDGTE applies the GTE predicate on the "user_uuid" field.
func UserUUIDGTE(v string) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldUserUUID, v))
}

// UserUUIDLT applies the LT predicate on the "user_uuid" field.
func UserUUIDLT(v string) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldUserUUID, v))
}

// UserUUIDLTE applies the LTE predicate on the "user_uuid" field.
func UserUUIDLTE(v string) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldUserUUID, v))
}

// UserUUIDContains applies the Contains predicate on the "user_uuid" field.
func UserUUIDContains(v string) predicate.Record {
	return predicate.Record(sql.FieldContains(FieldUserUUID, v))
}

// UserUUIDHasPrefix applies the HasPrefix predicate on the "user_uuid" field.
func UserUUIDHasPrefix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasPrefix(FieldUserUUID, v))
}

// UserUUIDHasSuffix applies the HasSuffix predicate on the "user_uuid" field.
func UserUUIDHasSuffix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasSuffix(FieldUserUUID, v))
}

// UserUUIDEqualFold applies the EqualFold predicate on the "user_uuid" field.
func UserUUIDEqualFold(v string) predicate.Record {
	return predicate.Record(sql.FieldEqualFold(FieldUserUUID, v))
}

// UserUUIDContainsFold applies the ContainsFold predicate on the "user_uuid" field.
func UserUUIDContainsFold(v string) predicate.Record {
	return predicate.Record(sql.FieldContainsFold(FieldUserUUID, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Record {
	return predicate.Record(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Record {
	return predicate.Record(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Record {
	return predicate.Record(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Record {
	return predicate.Record(sql.FieldContainsFold(FieldCode, v))
}

// ReadabilityEQ applies the EQ predicate on the "readability" field.
func ReadabilityEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldReadability, v))
}

// ReadabilityNEQ applies the NEQ predicate on the "readability" field.
func ReadabilityNEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldReadability, v))
}

// ReadabilityIn applies the In predicate on the "readability" field.
func ReadabilityIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldReadability, vs...))
}

// ReadabilityNotIn applies the NotIn predicate on the "readability" field.
func ReadabilityNotIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldReadability, vs...))
}

// ReadabilityGT applies the GT predicate on the "readability" field.
func ReadabilityGT(v int) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldReadability, v))
}

// ReadabilityGTE applies the GTE predicate on the "readability" field.
func ReadabilityGTE(v int) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldReadability, v))
}

// ReadabilityLT applies the LT predicate on the "readability" field.
func ReadabilityLT(v int) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldReadability, v))
}

// ReadabilityLTE applies the LTE predicate on the "readability" field.
func ReadabilityLTE(v int) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldReadability, v))
}

// ModularityEQ applies the EQ predicate on the "modularity" field.
func ModularityEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldModularity, v))
}

// ModularityNEQ applies the NEQ predicate on the "modularity" field.
func ModularityNEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldModularity, v))
}

// ModularityIn applies the In predicate on the "modularity" field.
func ModularityIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldModularity, vs...))
}

// ModularityNotIn applies the NotIn predicate on the "modularity" field.
func ModularityNotIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldModularity, vs...))
}

// ModularityGT applies the GT predicate on the "modularity" field.
func ModularityGT(v int) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldModularity, v))
}

// ModularityGTE applies the GTE predicate on the "modularity" field.
func ModularityGTE(v int) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldModularity, v))
}

// ModularityLT applies the LT predicate on the "modularity" field.
func ModularityLT(v int) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldModularity, v))
}

// ModularityLTE applies the LTE predicate on the "modularity" field.
func ModularityLTE(v int) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldModularity, v))
}

// EfficiencyEQ applies the EQ predicate on the "efficiency" field.
func EfficiencyEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldEfficiency, v))
}

// EfficiencyNEQ applies the NEQ predicate on the "efficiency" field.
func EfficiencyNEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldEfficiency, v))
}

// EfficiencyIn applies the In predicate on the "efficiency" field.
func EfficiencyIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldEfficiency, vs...))
}

// EfficiencyNotIn applies the NotIn predicate on the "efficiency" field.
func EfficiencyNotIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldEfficiency, vs...))
}

// EfficiencyGT applies the GT predicate on the "efficiency" field.
func EfficiencyGT(v int) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldEfficiency, v))
}

// EfficiencyGTE applies the GTE predicate on the "efficiency" field.
func EfficiencyGTE(v int) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldEfficiency, v))
}

// EfficiencyLT applies the LT predicate on the "efficiency" field.
func EfficiencyLT(v int) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldEfficiency, v))
}

// EfficiencyLTE applies the LTE predicate on the "efficiency" field.
func EfficiencyLTE(v int) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldEfficiency, v))
}

// TestabilityEQ applies the EQ predicate on the "testability" field.
func TestabilityEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldTestability, v))
}

// TestabilityNEQ applies the NEQ predicate on the "testability" field.
func TestabilityNEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldTestability, v))
}

// TestabilityIn applies the In predicate on the "testability" field.
func TestabilityIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldTestability, vs...))
}

// TestabilityNotIn applies the NotIn predicate on the "testability" field.
func TestabilityNotIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldTestability, vs...))
}

// TestabilityGT applies the GT predicate on the "testability" field.
func TestabilityGT(v int) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldTestability, v))
}

// TestabilityGTE applies the GTE predicate on the "testability" field.
func TestabilityGTE(v int) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldTestability, v))
}

// TestabilityLT applies the LT predicate on the "testability" field.
func TestabilityLT(v int) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldTestability, v))
}

// TestabilityLTE applies the LTE predicate on the "testability" field.
func TestabilityLTE(v int) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldTestability, v))
}

// MaintainablityEQ applies the EQ predicate on the "maintainablity" field.
func MaintainablityEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldEQ(FieldMaintainablity, v))
}

// MaintainablityNEQ applies the NEQ predicate on the "maintainablity" field.
func MaintainablityNEQ(v int) predicate.Record {
	return predicate.Record(sql.FieldNEQ(FieldMaintainablity, v))
}

// MaintainablityIn applies the In predicate on the "maintainablity" field.
func MaintainablityIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldIn(FieldMaintainablity, vs...))
}

// MaintainablityNotIn applies the NotIn predicate on the "maintainablity" field.
func MaintainablityNotIn(vs ...int) predicate.Record {
	return predicate.Record(sql.FieldNotIn(FieldMaintainablity, vs...))
}

// MaintainablityGT applies the GT predicate on the "maintainablity" field.
func MaintainablityGT(v int) predicate.Record {
	return predicate.Record(sql.FieldGT(FieldMaintainablity, v))
}

// MaintainablityGTE applies the GTE predicate on the "maintainablity" field.
func MaintainablityGTE(v int) predicate.Record {
	return predicate.Record(sql.FieldGTE(FieldMaintainablity, v))
}

// MaintainablityLT applies the LT predicate on the "maintainablity" field.
func MaintainablityLT(v int) predicate.Record {
	return predicate.Record(sql.FieldLT(FieldMaintainablity, v))
}

// MaintainablityLTE applies the LTE predicate on the "maintainablity" field.
func MaintainablityLTE(v int) predicate.Record {
	return predicate.Record(sql.FieldLTE(FieldMaintainablity, v))
}

// HasProblem applies the HasEdge predicate on the "problem" edge.
func HasProblem() predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProblemTable, ProblemPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProblemWith applies the HasEdge predicate on the "problem" edge with a given conditions (other predicates).
func HasProblemWith(preds ...predicate.Problem) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		step := newProblemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
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
func Not(p predicate.Record) predicate.Record {
	return predicate.Record(func(s *sql.Selector) {
		p(s.Not())
	})
}
