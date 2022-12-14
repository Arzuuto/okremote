// Code generated by entc, DO NOT EDIT.

package users

import (
	"Backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// FistName applies equality check predicate on the "FistName" field. It's identical to FistNameEQ.
func FistName(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFistName), v))
	})
}

// LastName applies equality check predicate on the "LastName" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// Email applies equality check predicate on the "Email" field. It's identical to EmailEQ.
func Email(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// Password applies equality check predicate on the "Password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// FistNameEQ applies the EQ predicate on the "FistName" field.
func FistNameEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFistName), v))
	})
}

// FistNameNEQ applies the NEQ predicate on the "FistName" field.
func FistNameNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFistName), v))
	})
}

// FistNameIn applies the In predicate on the "FistName" field.
func FistNameIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFistName), v...))
	})
}

// FistNameNotIn applies the NotIn predicate on the "FistName" field.
func FistNameNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFistName), v...))
	})
}

// FistNameGT applies the GT predicate on the "FistName" field.
func FistNameGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFistName), v))
	})
}

// FistNameGTE applies the GTE predicate on the "FistName" field.
func FistNameGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFistName), v))
	})
}

// FistNameLT applies the LT predicate on the "FistName" field.
func FistNameLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFistName), v))
	})
}

// FistNameLTE applies the LTE predicate on the "FistName" field.
func FistNameLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFistName), v))
	})
}

// FistNameContains applies the Contains predicate on the "FistName" field.
func FistNameContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFistName), v))
	})
}

// FistNameHasPrefix applies the HasPrefix predicate on the "FistName" field.
func FistNameHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFistName), v))
	})
}

// FistNameHasSuffix applies the HasSuffix predicate on the "FistName" field.
func FistNameHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFistName), v))
	})
}

// FistNameEqualFold applies the EqualFold predicate on the "FistName" field.
func FistNameEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFistName), v))
	})
}

// FistNameContainsFold applies the ContainsFold predicate on the "FistName" field.
func FistNameContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFistName), v))
	})
}

// LastNameEQ applies the EQ predicate on the "LastName" field.
func LastNameEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastName), v))
	})
}

// LastNameNEQ applies the NEQ predicate on the "LastName" field.
func LastNameNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastName), v))
	})
}

// LastNameIn applies the In predicate on the "LastName" field.
func LastNameIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLastName), v...))
	})
}

// LastNameNotIn applies the NotIn predicate on the "LastName" field.
func LastNameNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLastName), v...))
	})
}

// LastNameGT applies the GT predicate on the "LastName" field.
func LastNameGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastName), v))
	})
}

// LastNameGTE applies the GTE predicate on the "LastName" field.
func LastNameGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastName), v))
	})
}

// LastNameLT applies the LT predicate on the "LastName" field.
func LastNameLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastName), v))
	})
}

// LastNameLTE applies the LTE predicate on the "LastName" field.
func LastNameLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastName), v))
	})
}

// LastNameContains applies the Contains predicate on the "LastName" field.
func LastNameContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLastName), v))
	})
}

// LastNameHasPrefix applies the HasPrefix predicate on the "LastName" field.
func LastNameHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLastName), v))
	})
}

// LastNameHasSuffix applies the HasSuffix predicate on the "LastName" field.
func LastNameHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLastName), v))
	})
}

// LastNameEqualFold applies the EqualFold predicate on the "LastName" field.
func LastNameEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLastName), v))
	})
}

// LastNameContainsFold applies the ContainsFold predicate on the "LastName" field.
func LastNameContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLastName), v))
	})
}

// EmailEQ applies the EQ predicate on the "Email" field.
func EmailEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEmail), v))
	})
}

// EmailNEQ applies the NEQ predicate on the "Email" field.
func EmailNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEmail), v))
	})
}

// EmailIn applies the In predicate on the "Email" field.
func EmailIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEmail), v...))
	})
}

// EmailNotIn applies the NotIn predicate on the "Email" field.
func EmailNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEmail), v...))
	})
}

// EmailGT applies the GT predicate on the "Email" field.
func EmailGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEmail), v))
	})
}

// EmailGTE applies the GTE predicate on the "Email" field.
func EmailGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEmail), v))
	})
}

// EmailLT applies the LT predicate on the "Email" field.
func EmailLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEmail), v))
	})
}

// EmailLTE applies the LTE predicate on the "Email" field.
func EmailLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEmail), v))
	})
}

// EmailContains applies the Contains predicate on the "Email" field.
func EmailContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEmail), v))
	})
}

// EmailHasPrefix applies the HasPrefix predicate on the "Email" field.
func EmailHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEmail), v))
	})
}

// EmailHasSuffix applies the HasSuffix predicate on the "Email" field.
func EmailHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEmail), v))
	})
}

// EmailEqualFold applies the EqualFold predicate on the "Email" field.
func EmailEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEmail), v))
	})
}

// EmailContainsFold applies the ContainsFold predicate on the "Email" field.
func EmailContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEmail), v))
	})
}

// PasswordEQ applies the EQ predicate on the "Password" field.
func PasswordEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "Password" field.
func PasswordNEQ(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "Password" field.
func PasswordIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "Password" field.
func PasswordNotIn(vs ...string) predicate.Users {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Users(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "Password" field.
func PasswordGT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "Password" field.
func PasswordGTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "Password" field.
func PasswordLT(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "Password" field.
func PasswordLTE(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "Password" field.
func PasswordContains(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "Password" field.
func PasswordHasPrefix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "Password" field.
func PasswordHasSuffix(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "Password" field.
func PasswordEqualFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "Password" field.
func PasswordContainsFold(v string) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// HasAdvertisement applies the HasEdge predicate on the "advertisement" edge.
func HasAdvertisement() predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdvertisementTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AdvertisementTable, AdvertisementColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAdvertisementWith applies the HasEdge predicate on the "advertisement" edge with a given conditions (other predicates).
func HasAdvertisementWith(preds ...predicate.Advertisement) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AdvertisementInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AdvertisementTable, AdvertisementColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
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
func Not(p predicate.Users) predicate.Users {
	return predicate.Users(func(s *sql.Selector) {
		p(s.Not())
	})
}
