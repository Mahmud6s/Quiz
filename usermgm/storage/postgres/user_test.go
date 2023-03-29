package postgres

import (
	"quiz/usermgm/storage"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestRegisterUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "CREATE_USER_SUCESS",
			in: storage.User{
				FirstName: "mahmud",
				LastName:  "mahudd",
				Email:     "mahmud@gmail.com",
				Username:  "mahmudtry",
				Password:  "123456",
			},
			want: &storage.User{
				FirstName: "mahmud",
				LastName:  "mahudd",
				Email:     "mahmud@gmail.com",
				Username:  "mahmudtry",
				Password:  "123456",
				IsAdmin:   false,
				IsActive:  true,
			},
		},
		{
			name: "CREATE_USER_EMAIL_UNIQUE_FAILD",
			in: storage.User{
				FirstName: "mahmud",
				LastName:  "mahudd",
				Email:     "mahmud@gmail.com",
				Username:  "mahmutry",
				Password:  "123456",
			},
			wantErr: true,
		},
		{
			name: "CREATE_USER_USERNAME_UNIQUE_FAILD",
			in: storage.User{
				FirstName: "first3",
				LastName:  "last3",
				Email:     "test3@example.com",
				Username:  "mahmudtry",
				Password:  "161252",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RegisterUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.RegisterUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestRegisterAdmin(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})

	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "CREATE_ADMIN_SUCESS",
			in: storage.User{
				FirstName: "sayeem",
				LastName:  "mahud",
				Email:     "sayeem@gmail.com",
				Username:  "sayeem",
				Password:  "123456",
			},
			want: &storage.User{
				FirstName: "sayeem",
				LastName:  "mahud",
				Email:     "sayeem@gmail.com",
				Username:  "sayeem",
				Password:  "123456",
				IsAdmin:   true,
				IsActive:  true,
			},
		},
		{
			name: "CREATE_ADMIN_EMAIL_UNIQUE_FAILD",
			in: storage.User{
				FirstName: "sayeem1",
				LastName:  "mahud",
				Email:     "sayeem@gmail.com",
				Username:  "sayeem1",
				Password:  "123456",
			},
			wantErr: true,
		},
		{
			name: "CREATE_ADMIN_USERNAME_UNIQUE_FAILD",
			in: storage.User{
				FirstName: "first3",
				LastName:  "last3",
				Email:     "sayeem13@gmail.com",
				Username:  "sayeem",
				Password:  "161252",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.RegisterAdmin(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.RegisterAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}
			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.RegisterAdmin() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newUser := storage.User{
		FirstName: "rahim",
		LastName:  "khan",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "161252",
	}
	tests := []struct {
		name    string
		in      storage.User
		want    *storage.User
		wantErr bool
	}{
		{
			name: "Update_USER_SUCESS",
			in: storage.User{
				FirstName: "rahimupdate",
				LastName:  "khanupdate",
				IsAdmin:   false,
				IsActive:  true,
			},
			want: &storage.User{
				FirstName: "rahimupdate",
				LastName:  "khanupdate",
				Email:     "rahim@example.com",
				Username:  "rahim",
				Password:  "161252",
				IsAdmin:   false,
				IsActive:  true,
			},
		},
	}
	user, err := s.RegisterUser(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.ID = user.ID
			got, err := s.UpdateUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Email", "Username", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UpdateUser() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newUser := storage.User{
		FirstName: "rahim",
		LastName:  "khan",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "161252",
	}
	user, err := s.RegisterUser(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	id := int32(user.ID)
	tests := []struct {
		name    string
		in      int32
		wantErr bool
	}{
		{
			name: "DELETE_USER_SUCESS",
			in:   id,
		},
		{
			name:    "DELETE_USER_FAILD",
			in:      id,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := s.DeleteUser(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.DeleteUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	newUser := storage.User{
		FirstName: "rahimer",
		LastName:  "khanfg",
		Email:     "rahim@example.com",
		Username:  "rahim",
		Password:  "161252",
	}
	user, err := s.RegisterUser(newUser)
	if err != nil {
		t.Fatalf("PostgresStorage.RegisterUser() error = %v", err)
	}
	username := user.Username
	tests := []struct {
		name    string
		in      string
		want    *storage.User
		wantErr bool
	}{
		{
			name: "GET_USER_BY_USERNAME_SUCESS",
			in:   username,
			want: &storage.User{
				FirstName: "rahimer",
				LastName:  "khanfg",
				Email:     "rahim@example.com",
				Username:  "rahim",
				IsAdmin:   false,
				IsActive:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.GetUserByUsername(tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.GetUserByUsername() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}

func TestUserList(t *testing.T) {
	s, tr := NewTestStorage(getDBConnectionString(), getMigrationDir())
	t.Parallel()

	t.Cleanup(func() {
		tr()
	})
	users := []storage.User{
		{
			FirstName: "jabbar",
			LastName:  "khan",
			Email:     "jabbar@example.com",
			Username:  "jabbar",
			Password:  "12345678",
		},
		{
			FirstName: "ratul",
			LastName:  "khan",
			Email:     "ratul@example.com",
			Username:  "ratul",
			Password:  "12345678",
		},
		{
			FirstName: "pranto",
			LastName:  "khan",
			Email:     "pranto@example.com",
			Username:  "pranto",
			Password:  "12345678",
		},
	}
	for _, user := range users {
		_, err := s.RegisterUser(user)
		if err != nil {
			t.Fatalf("unable to create user for list user testing %v", err)
		}
	}
	tests := []struct {
		name    string
		in      storage.User
		want    []storage.User
		wantErr bool
	}{
		{
			name: "LIST_ALL_USER_SUCCESS",
			in:   storage.User{},
			want: []storage.User{
				{
					FirstName: "jabbar",
					LastName:  "khan",
					Email:     "jabbar@example.com",
					Username:  "jabbar",
					IsActive:  true,
					IsAdmin:   false,
				},
				{
					FirstName: "ratul",
					LastName:  "khan",
					Email:     "ratul@example.com",
					Username:  "ratul",
					IsActive:  true,
					IsAdmin:   false,
				},
				{
					FirstName: "pranto",
					LastName:  "khan",
					Email:     "pranto@example.com",
					Username:  "pranto",
					IsActive:  true,
					IsAdmin:   false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.UserList()
			if (err != nil) != tt.wantErr {
				t.Errorf("PostgresStorage.UserList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				cmpopts.IgnoreFields(storage.User{}, "ID", "Password", "CreatedAt", "UpdatedAt", "DeletedAt"),
			}

			sort.SliceStable(got, func(i, j int) bool {
				return got[i].ID < got[j].ID
			})

			if !cmp.Equal(got, tt.want, opts...) {
				t.Errorf("PostgresStorage.UserList() diff = %v", cmp.Diff(got, tt.want, opts...))
			}
		})
	}
}
