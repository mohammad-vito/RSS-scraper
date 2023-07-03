package user

import (
	"RssReader/bussiness/data/store/user"
	"RssReader/bussiness/data/tests"
	"RssReader/foundation/docker"
	"context"
	"fmt"
	"testing"
	"time"
)

var dc *docker.Container

func TestMain(m *testing.M) {
	var err error
	dc, err = tests.StartDB()
	time.Sleep(time.Duration(time.Second * 2))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tests.StopDB(dc)
	m.Run()

}

func TestCore_Create(t *testing.T) {
	log, db, teardown := tests.NewUnit(t, dc)
	core := NewCore(log, db)
	t.Cleanup(teardown)
	t.Log("Given the new user entity")
	{
		u := user.NewUser{
			Name:            "testi",
			Email:           "test@gmail.com",
			Roles:           nil,
			Password:        "test1",
			PasswordConfirm: "test1",
		}
		t.Logf("When user's email doesn't exist")
		ctx := context.Background()
		_, err := core.Create(ctx, u, time.Now())
		if err != nil {
			t.Fatalf("\t%s\tTest :\tShould be able to create user :", err)
		}
		s := user.NewStore(log, db)
		usr, err := s.QueryByEmail(ctx, u.Email)

		if usr.Email != u.Email || usr.Name != u.Name {
			t.Fatalf("Test: The email and name Should be the same as the input value ")
		}
		t.Logf("Test: The email and name Should be the same as the input value ")
	}
}

func TestCore_Authenticate(t *testing.T) {

}
