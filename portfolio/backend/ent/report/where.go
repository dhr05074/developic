// Code generated by ent, DO NOT EDIT.

package report

import (
	"portfolio/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldID, id))
}

// RequestID applies equality check predicate on the "request_id" field. It's identical to RequestIDEQ.
func RequestID(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldRequestID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldStatus, v))
}

// RequestIDEQ applies the EQ predicate on the "request_id" field.
func RequestIDEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldRequestID, v))
}

// RequestIDNEQ applies the NEQ predicate on the "request_id" field.
func RequestIDNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldRequestID, v))
}

// RequestIDIn applies the In predicate on the "request_id" field.
func RequestIDIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldRequestID, vs...))
}

// RequestIDNotIn applies the NotIn predicate on the "request_id" field.
func RequestIDNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldRequestID, vs...))
}

// RequestIDGT applies the GT predicate on the "request_id" field.
func RequestIDGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldRequestID, v))
}

// RequestIDGTE applies the GTE predicate on the "request_id" field.
func RequestIDGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldRequestID, v))
}

// RequestIDLT applies the LT predicate on the "request_id" field.
func RequestIDLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldRequestID, v))
}

// RequestIDLTE applies the LTE predicate on the "request_id" field.
func RequestIDLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldRequestID, v))
}

// RequestIDContains applies the Contains predicate on the "request_id" field.
func RequestIDContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldRequestID, v))
}

// RequestIDHasPrefix applies the HasPrefix predicate on the "request_id" field.
func RequestIDHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldRequestID, v))
}

// RequestIDHasSuffix applies the HasSuffix predicate on the "request_id" field.
func RequestIDHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldRequestID, v))
}

// RequestIDEqualFold applies the EqualFold predicate on the "request_id" field.
func RequestIDEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldRequestID, v))
}

// RequestIDContainsFold applies the ContainsFold predicate on the "request_id" field.
func RequestIDContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldRequestID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Report {
	return predicate.Report(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Report {
	return predicate.Report(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Report {
	return predicate.Report(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Report {
	return predicate.Report(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Report {
	return predicate.Report(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Report {
	return predicate.Report(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Report {
	return predicate.Report(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Report {
	return predicate.Report(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Report {
	return predicate.Report(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Report {
	return predicate.Report(sql.FieldContainsFold(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
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
func Not(p predicate.Report) predicate.Report {
	return predicate.Report(func(s *sql.Selector) {
		p(s.Not())
	})
}
