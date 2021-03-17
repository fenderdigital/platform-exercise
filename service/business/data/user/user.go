// Package user contains user related CRUD functionality.
package user

import (
	"context"
	"fmt"

	//"database/sql"
	"time"

	"platform-exercise/service/business/auth"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrNotFound is used when a specific User is requested but does not exist.
	ErrNotFound = errors.New("not found")

	// ErrInvalidID occurs when an ID is not in a valid form.
	ErrInvalidID = errors.New("ID is not in its proper form")

	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("authentication failed")

	// ErrForbidden occurs when a user tries to do something that is forbidden to them according to our access control policies.
	ErrForbidden = errors.New("attempted action is not allowed")

	userDB map[string]User
)

func init() {
	userDB = make(map[string]User)
}

// Create inserts a new user into the database.
func Create(ctx context.Context, db *sqlx.DB, nu NewUser, now time.Time) (User, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.user.create")
	defer span.End()

	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, errors.Wrap(err, "generating password hash")
	}

	u := User{
		ID:           uuid.New().String(),
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: hash,
		Roles:        nu.Roles,
		DateCreated:  now.UTC(),
		DateUpdated:  now.UTC(),
	}

	userDB[u.ID] = u

	fmt.Printf("Len of UserDB : %d\n", len(userDB))
	fmt.Printf("UserDB : %+v\n", userDB)

	return u, nil
}

// Update replaces a user document in the database.
func Update(ctx context.Context, claims auth.Claims, db *sqlx.DB, id string, uu UpdateUser, now time.Time) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.user.update")
	defer span.End()

	u, err := One(ctx, claims, db, id)
	if err != nil {
		return err
	}

	if uu.Name != nil {
		u.Name = *uu.Name
	}
	if uu.Email != nil {
		u.Email = *uu.Email
	}
	if uu.Roles != nil {
		u.Roles = uu.Roles
	}
	if uu.Password != nil {
		pw, err := bcrypt.GenerateFromPassword([]byte(*uu.Password), bcrypt.DefaultCost)
		if err != nil {
			return errors.Wrap(err, "generating password hash")
		}
		u.PasswordHash = pw
	}
	u.DateUpdated = now

	userDB[id] = u

	fmt.Printf("Len of UserDB : %d\n", len(userDB))
	fmt.Printf("UserDB : %+v\n", userDB)

	return nil
}

// Delete removes a user from the database.
func Delete(ctx context.Context, db *sqlx.DB, id string) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.user.delete")
	defer span.End()

	if _, err := uuid.Parse(id); err != nil {
		return ErrInvalidID
	}

	delete(userDB, id)

	fmt.Printf("Len of UserDB : %d\n", len(userDB))
	fmt.Printf("UserDB : %+v\n", userDB)

	return nil
}

// List retrieves a list of existing users from the database.
func List(ctx context.Context, db *sqlx.DB) ([]User, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.user.list")
	defer span.End()

	users := make([]User, 0, 1)
	for _, v := range userDB {
		users = append(users, v)
	}

	return users, nil
}

// One gets the specified user from the database.
func One(ctx context.Context, claims auth.Claims, db *sqlx.DB, userID string) (User, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.user.one")
	defer span.End()

	if _, err := uuid.Parse(userID); err != nil {
		return User{}, ErrInvalidID
	}

	// If you are not an admin and looking to retrieve someone other than yourself.
	if !claims.HasRole(auth.RoleAdmin) && claims.Subject != userID {
		return User{}, ErrForbidden
	}

	var u User
	u = userDB[userID]

	return u, nil
}

// Authenticate finds a user by their email and verifies their password. On
// success it returns a Claims value representing this user. The claims can be
// used to generate a token for future authentication.
func Authenticate(ctx context.Context, db *sqlx.DB, now time.Time, email, password string) (auth.Claims, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.user.authenticate")
	defer span.End()

	var u User
	found := false
	for _, v := range userDB {
		if v.Email == email {
			u = v
			found = true
			break
		}
	}
	if found == false {
		fmt.Println("User not found")
		return auth.Claims{}, ErrAuthenticationFailure
	}

	// Compare the provided password with the saved hash. Use the bcrypt
	// comparison function so it is cryptographically secure.
	if err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)); err != nil {
		fmt.Println("User password failed")
		return auth.Claims{}, ErrAuthenticationFailure
	}

	// If we are this far the request is valid. Create some claims for the user
	// and generate their token.
	claims := auth.Claims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "test team",
			Subject:   u.ID,
			Audience:  "test customer",
			ExpiresAt: now.Add(time.Hour).Unix(),
			IssuedAt:  now.Unix(),
		},
		Roles: u.Roles,
	}

	return claims, nil
}
