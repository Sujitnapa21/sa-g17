// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/Sujitnapa29/app/ent/adult"
	"github.com/Sujitnapa29/app/ent/books"
	"github.com/Sujitnapa29/app/ent/customer"
	"github.com/Sujitnapa29/app/ent/kid"
	"github.com/Sujitnapa29/app/ent/room"
	"github.com/Sujitnapa29/app/ent/roomamount"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Books is the model entity for the Books schema.
type Books struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Checkin holds the value of the "Checkin" field.
	Checkin time.Time `json:"Checkin,omitempty"`
	// Checkout holds the value of the "Checkout" field.
	Checkout time.Time `json:"Checkout,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BooksQuery when eager-loading is set.
	Edges            BooksEdges `json:"edges"`
	adult_books      *int
	customer_books   *int
	kid_books        *int
	room_books       *int
	roomamount_books *int
}

// BooksEdges holds the relations/edges for other nodes in the graph.
type BooksEdges struct {
	// Adult holds the value of the adult edge.
	Adult *Adult
	// Kid holds the value of the kid edge.
	Kid *Kid
	// Roomamount holds the value of the roomamount edge.
	Roomamount *Roomamount
	// Room holds the value of the room edge.
	Room *Room
	// Customer holds the value of the customer edge.
	Customer *Customer
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// AdultOrErr returns the Adult value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BooksEdges) AdultOrErr() (*Adult, error) {
	if e.loadedTypes[0] {
		if e.Adult == nil {
			// The edge adult was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: adult.Label}
		}
		return e.Adult, nil
	}
	return nil, &NotLoadedError{edge: "adult"}
}

// KidOrErr returns the Kid value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BooksEdges) KidOrErr() (*Kid, error) {
	if e.loadedTypes[1] {
		if e.Kid == nil {
			// The edge kid was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: kid.Label}
		}
		return e.Kid, nil
	}
	return nil, &NotLoadedError{edge: "kid"}
}

// RoomamountOrErr returns the Roomamount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BooksEdges) RoomamountOrErr() (*Roomamount, error) {
	if e.loadedTypes[2] {
		if e.Roomamount == nil {
			// The edge roomamount was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: roomamount.Label}
		}
		return e.Roomamount, nil
	}
	return nil, &NotLoadedError{edge: "roomamount"}
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BooksEdges) RoomOrErr() (*Room, error) {
	if e.loadedTypes[3] {
		if e.Room == nil {
			// The edge room was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: room.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// CustomerOrErr returns the Customer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BooksEdges) CustomerOrErr() (*Customer, error) {
	if e.loadedTypes[4] {
		if e.Customer == nil {
			// The edge customer was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: customer.Label}
		}
		return e.Customer, nil
	}
	return nil, &NotLoadedError{edge: "customer"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Books) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // Checkin
		&sql.NullTime{},  // Checkout
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Books) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // adult_books
		&sql.NullInt64{}, // customer_books
		&sql.NullInt64{}, // kid_books
		&sql.NullInt64{}, // room_books
		&sql.NullInt64{}, // roomamount_books
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Books fields.
func (b *Books) assignValues(values ...interface{}) error {
	if m, n := len(values), len(books.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field Checkin", values[0])
	} else if value.Valid {
		b.Checkin = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field Checkout", values[1])
	} else if value.Valid {
		b.Checkout = value.Time
	}
	values = values[2:]
	if len(values) == len(books.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field adult_books", value)
		} else if value.Valid {
			b.adult_books = new(int)
			*b.adult_books = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field customer_books", value)
		} else if value.Valid {
			b.customer_books = new(int)
			*b.customer_books = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field kid_books", value)
		} else if value.Valid {
			b.kid_books = new(int)
			*b.kid_books = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field room_books", value)
		} else if value.Valid {
			b.room_books = new(int)
			*b.room_books = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field roomamount_books", value)
		} else if value.Valid {
			b.roomamount_books = new(int)
			*b.roomamount_books = int(value.Int64)
		}
	}
	return nil
}

// QueryAdult queries the adult edge of the Books.
func (b *Books) QueryAdult() *AdultQuery {
	return (&BooksClient{config: b.config}).QueryAdult(b)
}

// QueryKid queries the kid edge of the Books.
func (b *Books) QueryKid() *KidQuery {
	return (&BooksClient{config: b.config}).QueryKid(b)
}

// QueryRoomamount queries the roomamount edge of the Books.
func (b *Books) QueryRoomamount() *RoomamountQuery {
	return (&BooksClient{config: b.config}).QueryRoomamount(b)
}

// QueryRoom queries the room edge of the Books.
func (b *Books) QueryRoom() *RoomQuery {
	return (&BooksClient{config: b.config}).QueryRoom(b)
}

// QueryCustomer queries the customer edge of the Books.
func (b *Books) QueryCustomer() *CustomerQuery {
	return (&BooksClient{config: b.config}).QueryCustomer(b)
}

// Update returns a builder for updating this Books.
// Note that, you need to call Books.Unwrap() before calling this method, if this Books
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Books) Update() *BooksUpdateOne {
	return (&BooksClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Books) Unwrap() *Books {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Books is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Books) String() string {
	var builder strings.Builder
	builder.WriteString("Books(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", Checkin=")
	builder.WriteString(b.Checkin.Format(time.ANSIC))
	builder.WriteString(", Checkout=")
	builder.WriteString(b.Checkout.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BooksSlice is a parsable slice of Books.
type BooksSlice []*Books

func (b BooksSlice) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
