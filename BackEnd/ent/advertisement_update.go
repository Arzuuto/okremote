// Code generated by entc, DO NOT EDIT.

package ent

import (
	"Backend/ent/advertisement"
	"Backend/ent/predicate"
	"Backend/ent/users"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AdvertisementUpdate is the builder for updating Advertisement entities.
type AdvertisementUpdate struct {
	config
	hooks    []Hook
	mutation *AdvertisementMutation
}

// Where appends a list predicates to the AdvertisementUpdate builder.
func (au *AdvertisementUpdate) Where(ps ...predicate.Advertisement) *AdvertisementUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetDate sets the "Date" field.
func (au *AdvertisementUpdate) SetDate(s string) *AdvertisementUpdate {
	au.mutation.SetDate(s)
	return au
}

// SetContrat sets the "Contrat" field.
func (au *AdvertisementUpdate) SetContrat(s string) *AdvertisementUpdate {
	au.mutation.SetContrat(s)
	return au
}

// SetEntreprise sets the "Entreprise" field.
func (au *AdvertisementUpdate) SetEntreprise(s string) *AdvertisementUpdate {
	au.mutation.SetEntreprise(s)
	return au
}

// SetImage sets the "Image" field.
func (au *AdvertisementUpdate) SetImage(s string) *AdvertisementUpdate {
	au.mutation.SetImage(s)
	return au
}

// SetLocalisation sets the "Localisation" field.
func (au *AdvertisementUpdate) SetLocalisation(s string) *AdvertisementUpdate {
	au.mutation.SetLocalisation(s)
	return au
}

// SetName sets the "Name" field.
func (au *AdvertisementUpdate) SetName(s string) *AdvertisementUpdate {
	au.mutation.SetName(s)
	return au
}

// SetRemuneration sets the "Remuneration" field.
func (au *AdvertisementUpdate) SetRemuneration(s string) *AdvertisementUpdate {
	au.mutation.SetRemuneration(s)
	return au
}

// SetSector sets the "Sector" field.
func (au *AdvertisementUpdate) SetSector(s string) *AdvertisementUpdate {
	au.mutation.SetSector(s)
	return au
}

// SetUsersID sets the "users" edge to the Users entity by ID.
func (au *AdvertisementUpdate) SetUsersID(id int) *AdvertisementUpdate {
	au.mutation.SetUsersID(id)
	return au
}

// SetNillableUsersID sets the "users" edge to the Users entity by ID if the given value is not nil.
func (au *AdvertisementUpdate) SetNillableUsersID(id *int) *AdvertisementUpdate {
	if id != nil {
		au = au.SetUsersID(*id)
	}
	return au
}

// SetUsers sets the "users" edge to the Users entity.
func (au *AdvertisementUpdate) SetUsers(u *Users) *AdvertisementUpdate {
	return au.SetUsersID(u.ID)
}

// Mutation returns the AdvertisementMutation object of the builder.
func (au *AdvertisementUpdate) Mutation() *AdvertisementMutation {
	return au.mutation
}

// ClearUsers clears the "users" edge to the Users entity.
func (au *AdvertisementUpdate) ClearUsers() *AdvertisementUpdate {
	au.mutation.ClearUsers()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AdvertisementUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdvertisementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			if au.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AdvertisementUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AdvertisementUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AdvertisementUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AdvertisementUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   advertisement.Table,
			Columns: advertisement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: advertisement.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldDate,
		})
	}
	if value, ok := au.mutation.Contrat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldContrat,
		})
	}
	if value, ok := au.mutation.Entreprise(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldEntreprise,
		})
	}
	if value, ok := au.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldImage,
		})
	}
	if value, ok := au.mutation.Localisation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldLocalisation,
		})
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldName,
		})
	}
	if value, ok := au.mutation.Remuneration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldRemuneration,
		})
	}
	if value, ok := au.mutation.Sector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldSector,
		})
	}
	if au.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   advertisement.UsersTable,
			Columns: []string{advertisement.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: users.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   advertisement.UsersTable,
			Columns: []string{advertisement.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: users.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{advertisement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AdvertisementUpdateOne is the builder for updating a single Advertisement entity.
type AdvertisementUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AdvertisementMutation
}

// SetDate sets the "Date" field.
func (auo *AdvertisementUpdateOne) SetDate(s string) *AdvertisementUpdateOne {
	auo.mutation.SetDate(s)
	return auo
}

// SetContrat sets the "Contrat" field.
func (auo *AdvertisementUpdateOne) SetContrat(s string) *AdvertisementUpdateOne {
	auo.mutation.SetContrat(s)
	return auo
}

// SetEntreprise sets the "Entreprise" field.
func (auo *AdvertisementUpdateOne) SetEntreprise(s string) *AdvertisementUpdateOne {
	auo.mutation.SetEntreprise(s)
	return auo
}

// SetImage sets the "Image" field.
func (auo *AdvertisementUpdateOne) SetImage(s string) *AdvertisementUpdateOne {
	auo.mutation.SetImage(s)
	return auo
}

// SetLocalisation sets the "Localisation" field.
func (auo *AdvertisementUpdateOne) SetLocalisation(s string) *AdvertisementUpdateOne {
	auo.mutation.SetLocalisation(s)
	return auo
}

// SetName sets the "Name" field.
func (auo *AdvertisementUpdateOne) SetName(s string) *AdvertisementUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetRemuneration sets the "Remuneration" field.
func (auo *AdvertisementUpdateOne) SetRemuneration(s string) *AdvertisementUpdateOne {
	auo.mutation.SetRemuneration(s)
	return auo
}

// SetSector sets the "Sector" field.
func (auo *AdvertisementUpdateOne) SetSector(s string) *AdvertisementUpdateOne {
	auo.mutation.SetSector(s)
	return auo
}

// SetUsersID sets the "users" edge to the Users entity by ID.
func (auo *AdvertisementUpdateOne) SetUsersID(id int) *AdvertisementUpdateOne {
	auo.mutation.SetUsersID(id)
	return auo
}

// SetNillableUsersID sets the "users" edge to the Users entity by ID if the given value is not nil.
func (auo *AdvertisementUpdateOne) SetNillableUsersID(id *int) *AdvertisementUpdateOne {
	if id != nil {
		auo = auo.SetUsersID(*id)
	}
	return auo
}

// SetUsers sets the "users" edge to the Users entity.
func (auo *AdvertisementUpdateOne) SetUsers(u *Users) *AdvertisementUpdateOne {
	return auo.SetUsersID(u.ID)
}

// Mutation returns the AdvertisementMutation object of the builder.
func (auo *AdvertisementUpdateOne) Mutation() *AdvertisementMutation {
	return auo.mutation
}

// ClearUsers clears the "users" edge to the Users entity.
func (auo *AdvertisementUpdateOne) ClearUsers() *AdvertisementUpdateOne {
	auo.mutation.ClearUsers()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AdvertisementUpdateOne) Select(field string, fields ...string) *AdvertisementUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Advertisement entity.
func (auo *AdvertisementUpdateOne) Save(ctx context.Context) (*Advertisement, error) {
	var (
		err  error
		node *Advertisement
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AdvertisementMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			if auo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AdvertisementUpdateOne) SaveX(ctx context.Context) *Advertisement {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AdvertisementUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AdvertisementUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AdvertisementUpdateOne) sqlSave(ctx context.Context) (_node *Advertisement, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   advertisement.Table,
			Columns: advertisement.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: advertisement.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Advertisement.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, advertisement.FieldID)
		for _, f := range fields {
			if !advertisement.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != advertisement.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldDate,
		})
	}
	if value, ok := auo.mutation.Contrat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldContrat,
		})
	}
	if value, ok := auo.mutation.Entreprise(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldEntreprise,
		})
	}
	if value, ok := auo.mutation.Image(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldImage,
		})
	}
	if value, ok := auo.mutation.Localisation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldLocalisation,
		})
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldName,
		})
	}
	if value, ok := auo.mutation.Remuneration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldRemuneration,
		})
	}
	if value, ok := auo.mutation.Sector(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: advertisement.FieldSector,
		})
	}
	if auo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   advertisement.UsersTable,
			Columns: []string{advertisement.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: users.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   advertisement.UsersTable,
			Columns: []string{advertisement.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: users.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Advertisement{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{advertisement.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
