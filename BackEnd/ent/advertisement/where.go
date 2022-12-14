// Code generated by entc, DO NOT EDIT.

package advertisement

import (
	"Backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Date applies equality check predicate on the "Date" field. It's identical to DateEQ.
func Date(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// Contrat applies equality check predicate on the "Contrat" field. It's identical to ContratEQ.
func Contrat(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContrat), v))
	})
}

// Entreprise applies equality check predicate on the "Entreprise" field. It's identical to EntrepriseEQ.
func Entreprise(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntreprise), v))
	})
}

// Image applies equality check predicate on the "Image" field. It's identical to ImageEQ.
func Image(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// Localisation applies equality check predicate on the "Localisation" field. It's identical to LocalisationEQ.
func Localisation(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocalisation), v))
	})
}

// Name applies equality check predicate on the "Name" field. It's identical to NameEQ.
func Name(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Remuneration applies equality check predicate on the "Remuneration" field. It's identical to RemunerationEQ.
func Remuneration(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemuneration), v))
	})
}

// Sector applies equality check predicate on the "Sector" field. It's identical to SectorEQ.
func Sector(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSector), v))
	})
}

// DateEQ applies the EQ predicate on the "Date" field.
func DateEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "Date" field.
func DateNEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "Date" field.
func DateIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "Date" field.
func DateNotIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "Date" field.
func DateGT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "Date" field.
func DateGTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "Date" field.
func DateLT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "Date" field.
func DateLTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// DateContains applies the Contains predicate on the "Date" field.
func DateContains(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDate), v))
	})
}

// DateHasPrefix applies the HasPrefix predicate on the "Date" field.
func DateHasPrefix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDate), v))
	})
}

// DateHasSuffix applies the HasSuffix predicate on the "Date" field.
func DateHasSuffix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDate), v))
	})
}

// DateEqualFold applies the EqualFold predicate on the "Date" field.
func DateEqualFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDate), v))
	})
}

// DateContainsFold applies the ContainsFold predicate on the "Date" field.
func DateContainsFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDate), v))
	})
}

// ContratEQ applies the EQ predicate on the "Contrat" field.
func ContratEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContrat), v))
	})
}

// ContratNEQ applies the NEQ predicate on the "Contrat" field.
func ContratNEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContrat), v))
	})
}

// ContratIn applies the In predicate on the "Contrat" field.
func ContratIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldContrat), v...))
	})
}

// ContratNotIn applies the NotIn predicate on the "Contrat" field.
func ContratNotIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldContrat), v...))
	})
}

// ContratGT applies the GT predicate on the "Contrat" field.
func ContratGT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContrat), v))
	})
}

// ContratGTE applies the GTE predicate on the "Contrat" field.
func ContratGTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContrat), v))
	})
}

// ContratLT applies the LT predicate on the "Contrat" field.
func ContratLT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContrat), v))
	})
}

// ContratLTE applies the LTE predicate on the "Contrat" field.
func ContratLTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContrat), v))
	})
}

// ContratContains applies the Contains predicate on the "Contrat" field.
func ContratContains(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContrat), v))
	})
}

// ContratHasPrefix applies the HasPrefix predicate on the "Contrat" field.
func ContratHasPrefix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContrat), v))
	})
}

// ContratHasSuffix applies the HasSuffix predicate on the "Contrat" field.
func ContratHasSuffix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContrat), v))
	})
}

// ContratEqualFold applies the EqualFold predicate on the "Contrat" field.
func ContratEqualFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContrat), v))
	})
}

// ContratContainsFold applies the ContainsFold predicate on the "Contrat" field.
func ContratContainsFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContrat), v))
	})
}

// EntrepriseEQ applies the EQ predicate on the "Entreprise" field.
func EntrepriseEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntreprise), v))
	})
}

// EntrepriseNEQ applies the NEQ predicate on the "Entreprise" field.
func EntrepriseNEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntreprise), v))
	})
}

// EntrepriseIn applies the In predicate on the "Entreprise" field.
func EntrepriseIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEntreprise), v...))
	})
}

// EntrepriseNotIn applies the NotIn predicate on the "Entreprise" field.
func EntrepriseNotIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEntreprise), v...))
	})
}

// EntrepriseGT applies the GT predicate on the "Entreprise" field.
func EntrepriseGT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntreprise), v))
	})
}

// EntrepriseGTE applies the GTE predicate on the "Entreprise" field.
func EntrepriseGTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntreprise), v))
	})
}

// EntrepriseLT applies the LT predicate on the "Entreprise" field.
func EntrepriseLT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntreprise), v))
	})
}

// EntrepriseLTE applies the LTE predicate on the "Entreprise" field.
func EntrepriseLTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntreprise), v))
	})
}

// EntrepriseContains applies the Contains predicate on the "Entreprise" field.
func EntrepriseContains(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEntreprise), v))
	})
}

// EntrepriseHasPrefix applies the HasPrefix predicate on the "Entreprise" field.
func EntrepriseHasPrefix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEntreprise), v))
	})
}

// EntrepriseHasSuffix applies the HasSuffix predicate on the "Entreprise" field.
func EntrepriseHasSuffix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEntreprise), v))
	})
}

// EntrepriseEqualFold applies the EqualFold predicate on the "Entreprise" field.
func EntrepriseEqualFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEntreprise), v))
	})
}

// EntrepriseContainsFold applies the ContainsFold predicate on the "Entreprise" field.
func EntrepriseContainsFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEntreprise), v))
	})
}

// ImageEQ applies the EQ predicate on the "Image" field.
func ImageEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldImage), v))
	})
}

// ImageNEQ applies the NEQ predicate on the "Image" field.
func ImageNEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldImage), v))
	})
}

// ImageIn applies the In predicate on the "Image" field.
func ImageIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldImage), v...))
	})
}

// ImageNotIn applies the NotIn predicate on the "Image" field.
func ImageNotIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldImage), v...))
	})
}

// ImageGT applies the GT predicate on the "Image" field.
func ImageGT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldImage), v))
	})
}

// ImageGTE applies the GTE predicate on the "Image" field.
func ImageGTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldImage), v))
	})
}

// ImageLT applies the LT predicate on the "Image" field.
func ImageLT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldImage), v))
	})
}

// ImageLTE applies the LTE predicate on the "Image" field.
func ImageLTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldImage), v))
	})
}

// ImageContains applies the Contains predicate on the "Image" field.
func ImageContains(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldImage), v))
	})
}

// ImageHasPrefix applies the HasPrefix predicate on the "Image" field.
func ImageHasPrefix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldImage), v))
	})
}

// ImageHasSuffix applies the HasSuffix predicate on the "Image" field.
func ImageHasSuffix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldImage), v))
	})
}

// ImageEqualFold applies the EqualFold predicate on the "Image" field.
func ImageEqualFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldImage), v))
	})
}

// ImageContainsFold applies the ContainsFold predicate on the "Image" field.
func ImageContainsFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldImage), v))
	})
}

// LocalisationEQ applies the EQ predicate on the "Localisation" field.
func LocalisationEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocalisation), v))
	})
}

// LocalisationNEQ applies the NEQ predicate on the "Localisation" field.
func LocalisationNEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocalisation), v))
	})
}

// LocalisationIn applies the In predicate on the "Localisation" field.
func LocalisationIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLocalisation), v...))
	})
}

// LocalisationNotIn applies the NotIn predicate on the "Localisation" field.
func LocalisationNotIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLocalisation), v...))
	})
}

// LocalisationGT applies the GT predicate on the "Localisation" field.
func LocalisationGT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocalisation), v))
	})
}

// LocalisationGTE applies the GTE predicate on the "Localisation" field.
func LocalisationGTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocalisation), v))
	})
}

// LocalisationLT applies the LT predicate on the "Localisation" field.
func LocalisationLT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocalisation), v))
	})
}

// LocalisationLTE applies the LTE predicate on the "Localisation" field.
func LocalisationLTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocalisation), v))
	})
}

// LocalisationContains applies the Contains predicate on the "Localisation" field.
func LocalisationContains(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocalisation), v))
	})
}

// LocalisationHasPrefix applies the HasPrefix predicate on the "Localisation" field.
func LocalisationHasPrefix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocalisation), v))
	})
}

// LocalisationHasSuffix applies the HasSuffix predicate on the "Localisation" field.
func LocalisationHasSuffix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocalisation), v))
	})
}

// LocalisationEqualFold applies the EqualFold predicate on the "Localisation" field.
func LocalisationEqualFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocalisation), v))
	})
}

// LocalisationContainsFold applies the ContainsFold predicate on the "Localisation" field.
func LocalisationContainsFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocalisation), v))
	})
}

// NameEQ applies the EQ predicate on the "Name" field.
func NameEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "Name" field.
func NameNEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "Name" field.
func NameIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "Name" field.
func NameNotIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "Name" field.
func NameGT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "Name" field.
func NameGTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "Name" field.
func NameLT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "Name" field.
func NameLTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "Name" field.
func NameContains(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "Name" field.
func NameHasPrefix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "Name" field.
func NameHasSuffix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "Name" field.
func NameEqualFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "Name" field.
func NameContainsFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// RemunerationEQ applies the EQ predicate on the "Remuneration" field.
func RemunerationEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRemuneration), v))
	})
}

// RemunerationNEQ applies the NEQ predicate on the "Remuneration" field.
func RemunerationNEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRemuneration), v))
	})
}

// RemunerationIn applies the In predicate on the "Remuneration" field.
func RemunerationIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRemuneration), v...))
	})
}

// RemunerationNotIn applies the NotIn predicate on the "Remuneration" field.
func RemunerationNotIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRemuneration), v...))
	})
}

// RemunerationGT applies the GT predicate on the "Remuneration" field.
func RemunerationGT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRemuneration), v))
	})
}

// RemunerationGTE applies the GTE predicate on the "Remuneration" field.
func RemunerationGTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRemuneration), v))
	})
}

// RemunerationLT applies the LT predicate on the "Remuneration" field.
func RemunerationLT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRemuneration), v))
	})
}

// RemunerationLTE applies the LTE predicate on the "Remuneration" field.
func RemunerationLTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRemuneration), v))
	})
}

// RemunerationContains applies the Contains predicate on the "Remuneration" field.
func RemunerationContains(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRemuneration), v))
	})
}

// RemunerationHasPrefix applies the HasPrefix predicate on the "Remuneration" field.
func RemunerationHasPrefix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRemuneration), v))
	})
}

// RemunerationHasSuffix applies the HasSuffix predicate on the "Remuneration" field.
func RemunerationHasSuffix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRemuneration), v))
	})
}

// RemunerationEqualFold applies the EqualFold predicate on the "Remuneration" field.
func RemunerationEqualFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRemuneration), v))
	})
}

// RemunerationContainsFold applies the ContainsFold predicate on the "Remuneration" field.
func RemunerationContainsFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRemuneration), v))
	})
}

// SectorEQ applies the EQ predicate on the "Sector" field.
func SectorEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSector), v))
	})
}

// SectorNEQ applies the NEQ predicate on the "Sector" field.
func SectorNEQ(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSector), v))
	})
}

// SectorIn applies the In predicate on the "Sector" field.
func SectorIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSector), v...))
	})
}

// SectorNotIn applies the NotIn predicate on the "Sector" field.
func SectorNotIn(vs ...string) predicate.Advertisement {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Advertisement(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSector), v...))
	})
}

// SectorGT applies the GT predicate on the "Sector" field.
func SectorGT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSector), v))
	})
}

// SectorGTE applies the GTE predicate on the "Sector" field.
func SectorGTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSector), v))
	})
}

// SectorLT applies the LT predicate on the "Sector" field.
func SectorLT(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSector), v))
	})
}

// SectorLTE applies the LTE predicate on the "Sector" field.
func SectorLTE(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSector), v))
	})
}

// SectorContains applies the Contains predicate on the "Sector" field.
func SectorContains(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSector), v))
	})
}

// SectorHasPrefix applies the HasPrefix predicate on the "Sector" field.
func SectorHasPrefix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSector), v))
	})
}

// SectorHasSuffix applies the HasSuffix predicate on the "Sector" field.
func SectorHasSuffix(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSector), v))
	})
}

// SectorEqualFold applies the EqualFold predicate on the "Sector" field.
func SectorEqualFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSector), v))
	})
}

// SectorContainsFold applies the ContainsFold predicate on the "Sector" field.
func SectorContainsFold(v string) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSector), v))
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.Users) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UsersInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Advertisement) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Advertisement) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
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
func Not(p predicate.Advertisement) predicate.Advertisement {
	return predicate.Advertisement(func(s *sql.Selector) {
		p(s.Not())
	})
}
