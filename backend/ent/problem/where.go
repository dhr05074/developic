// Code generated by ent, DO NOT EDIT.

package problem

import (
	"code-connect/ent/predicate"
	"code-connect/gateway"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldUUID, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldCode, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldTitle, v))
}

// Language applies equality check predicate on the "language" field. It's identical to LanguageEQ.
func Language(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldEQ(FieldLanguage, vc))
}

// Difficulty applies equality check predicate on the "difficulty" field. It's identical to DifficultyEQ.
func Difficulty(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldDifficulty, v))
}

// Readability applies equality check predicate on the "readability" field. It's identical to ReadabilityEQ.
func Readability(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldReadability, v))
}

// Modularity applies equality check predicate on the "modularity" field. It's identical to ModularityEQ.
func Modularity(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldModularity, v))
}

// Efficiency applies equality check predicate on the "efficiency" field. It's identical to EfficiencyEQ.
func Efficiency(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldEfficiency, v))
}

// Testability applies equality check predicate on the "testability" field. It's identical to TestabilityEQ.
func Testability(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldTestability, v))
}

// Maintainablity applies equality check predicate on the "maintainablity" field. It's identical to MaintainablityEQ.
func Maintainablity(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldMaintainablity, v))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContainsFold(FieldUUID, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContainsFold(FieldCode, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Problem {
	return predicate.Problem(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Problem {
	return predicate.Problem(sql.FieldContainsFold(FieldTitle, v))
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldEQ(FieldLanguage, vc))
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldNEQ(FieldLanguage, vc))
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...gateway.ProgrammingLanguage) predicate.Problem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Problem(sql.FieldIn(FieldLanguage, v...))
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...gateway.ProgrammingLanguage) predicate.Problem {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = string(vs[i])
	}
	return predicate.Problem(sql.FieldNotIn(FieldLanguage, v...))
}

// LanguageGT applies the GT predicate on the "language" field.
func LanguageGT(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldGT(FieldLanguage, vc))
}

// LanguageGTE applies the GTE predicate on the "language" field.
func LanguageGTE(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldGTE(FieldLanguage, vc))
}

// LanguageLT applies the LT predicate on the "language" field.
func LanguageLT(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldLT(FieldLanguage, vc))
}

// LanguageLTE applies the LTE predicate on the "language" field.
func LanguageLTE(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldLTE(FieldLanguage, vc))
}

// LanguageContains applies the Contains predicate on the "language" field.
func LanguageContains(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldContains(FieldLanguage, vc))
}

// LanguageHasPrefix applies the HasPrefix predicate on the "language" field.
func LanguageHasPrefix(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldHasPrefix(FieldLanguage, vc))
}

// LanguageHasSuffix applies the HasSuffix predicate on the "language" field.
func LanguageHasSuffix(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldHasSuffix(FieldLanguage, vc))
}

// LanguageEqualFold applies the EqualFold predicate on the "language" field.
func LanguageEqualFold(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldEqualFold(FieldLanguage, vc))
}

// LanguageContainsFold applies the ContainsFold predicate on the "language" field.
func LanguageContainsFold(v gateway.ProgrammingLanguage) predicate.Problem {
	vc := string(v)
	return predicate.Problem(sql.FieldContainsFold(FieldLanguage, vc))
}

// DifficultyEQ applies the EQ predicate on the "difficulty" field.
func DifficultyEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldDifficulty, v))
}

// DifficultyNEQ applies the NEQ predicate on the "difficulty" field.
func DifficultyNEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldDifficulty, v))
}

// DifficultyIn applies the In predicate on the "difficulty" field.
func DifficultyIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldDifficulty, vs...))
}

// DifficultyNotIn applies the NotIn predicate on the "difficulty" field.
func DifficultyNotIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldDifficulty, vs...))
}

// DifficultyGT applies the GT predicate on the "difficulty" field.
func DifficultyGT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldDifficulty, v))
}

// DifficultyGTE applies the GTE predicate on the "difficulty" field.
func DifficultyGTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldDifficulty, v))
}

// DifficultyLT applies the LT predicate on the "difficulty" field.
func DifficultyLT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldDifficulty, v))
}

// DifficultyLTE applies the LTE predicate on the "difficulty" field.
func DifficultyLTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldDifficulty, v))
}

// ReadabilityEQ applies the EQ predicate on the "readability" field.
func ReadabilityEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldReadability, v))
}

// ReadabilityNEQ applies the NEQ predicate on the "readability" field.
func ReadabilityNEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldReadability, v))
}

// ReadabilityIn applies the In predicate on the "readability" field.
func ReadabilityIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldReadability, vs...))
}

// ReadabilityNotIn applies the NotIn predicate on the "readability" field.
func ReadabilityNotIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldReadability, vs...))
}

// ReadabilityGT applies the GT predicate on the "readability" field.
func ReadabilityGT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldReadability, v))
}

// ReadabilityGTE applies the GTE predicate on the "readability" field.
func ReadabilityGTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldReadability, v))
}

// ReadabilityLT applies the LT predicate on the "readability" field.
func ReadabilityLT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldReadability, v))
}

// ReadabilityLTE applies the LTE predicate on the "readability" field.
func ReadabilityLTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldReadability, v))
}

// ModularityEQ applies the EQ predicate on the "modularity" field.
func ModularityEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldModularity, v))
}

// ModularityNEQ applies the NEQ predicate on the "modularity" field.
func ModularityNEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldModularity, v))
}

// ModularityIn applies the In predicate on the "modularity" field.
func ModularityIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldModularity, vs...))
}

// ModularityNotIn applies the NotIn predicate on the "modularity" field.
func ModularityNotIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldModularity, vs...))
}

// ModularityGT applies the GT predicate on the "modularity" field.
func ModularityGT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldModularity, v))
}

// ModularityGTE applies the GTE predicate on the "modularity" field.
func ModularityGTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldModularity, v))
}

// ModularityLT applies the LT predicate on the "modularity" field.
func ModularityLT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldModularity, v))
}

// ModularityLTE applies the LTE predicate on the "modularity" field.
func ModularityLTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldModularity, v))
}

// EfficiencyEQ applies the EQ predicate on the "efficiency" field.
func EfficiencyEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldEfficiency, v))
}

// EfficiencyNEQ applies the NEQ predicate on the "efficiency" field.
func EfficiencyNEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldEfficiency, v))
}

// EfficiencyIn applies the In predicate on the "efficiency" field.
func EfficiencyIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldEfficiency, vs...))
}

// EfficiencyNotIn applies the NotIn predicate on the "efficiency" field.
func EfficiencyNotIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldEfficiency, vs...))
}

// EfficiencyGT applies the GT predicate on the "efficiency" field.
func EfficiencyGT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldEfficiency, v))
}

// EfficiencyGTE applies the GTE predicate on the "efficiency" field.
func EfficiencyGTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldEfficiency, v))
}

// EfficiencyLT applies the LT predicate on the "efficiency" field.
func EfficiencyLT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldEfficiency, v))
}

// EfficiencyLTE applies the LTE predicate on the "efficiency" field.
func EfficiencyLTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldEfficiency, v))
}

// TestabilityEQ applies the EQ predicate on the "testability" field.
func TestabilityEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldTestability, v))
}

// TestabilityNEQ applies the NEQ predicate on the "testability" field.
func TestabilityNEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldTestability, v))
}

// TestabilityIn applies the In predicate on the "testability" field.
func TestabilityIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldTestability, vs...))
}

// TestabilityNotIn applies the NotIn predicate on the "testability" field.
func TestabilityNotIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldTestability, vs...))
}

// TestabilityGT applies the GT predicate on the "testability" field.
func TestabilityGT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldTestability, v))
}

// TestabilityGTE applies the GTE predicate on the "testability" field.
func TestabilityGTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldTestability, v))
}

// TestabilityLT applies the LT predicate on the "testability" field.
func TestabilityLT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldTestability, v))
}

// TestabilityLTE applies the LTE predicate on the "testability" field.
func TestabilityLTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldTestability, v))
}

// MaintainablityEQ applies the EQ predicate on the "maintainablity" field.
func MaintainablityEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldEQ(FieldMaintainablity, v))
}

// MaintainablityNEQ applies the NEQ predicate on the "maintainablity" field.
func MaintainablityNEQ(v int) predicate.Problem {
	return predicate.Problem(sql.FieldNEQ(FieldMaintainablity, v))
}

// MaintainablityIn applies the In predicate on the "maintainablity" field.
func MaintainablityIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldIn(FieldMaintainablity, vs...))
}

// MaintainablityNotIn applies the NotIn predicate on the "maintainablity" field.
func MaintainablityNotIn(vs ...int) predicate.Problem {
	return predicate.Problem(sql.FieldNotIn(FieldMaintainablity, vs...))
}

// MaintainablityGT applies the GT predicate on the "maintainablity" field.
func MaintainablityGT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGT(FieldMaintainablity, v))
}

// MaintainablityGTE applies the GTE predicate on the "maintainablity" field.
func MaintainablityGTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldGTE(FieldMaintainablity, v))
}

// MaintainablityLT applies the LT predicate on the "maintainablity" field.
func MaintainablityLT(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLT(FieldMaintainablity, v))
}

// MaintainablityLTE applies the LTE predicate on the "maintainablity" field.
func MaintainablityLTE(v int) predicate.Problem {
	return predicate.Problem(sql.FieldLTE(FieldMaintainablity, v))
}

// HasRecords applies the HasEdge predicate on the "records" edge.
func HasRecords() predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, RecordsTable, RecordsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRecordsWith applies the HasEdge predicate on the "records" edge with a given conditions (other predicates).
func HasRecordsWith(preds ...predicate.Record) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		step := newRecordsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Problem) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Problem) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
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
func Not(p predicate.Problem) predicate.Problem {
	return predicate.Problem(func(s *sql.Selector) {
		p(s.Not())
	})
}