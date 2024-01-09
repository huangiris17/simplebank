package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	db "github.com/huangiris17/simplebank/db/sqlc"
	"github.com/huangiris17/simplebank/util"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)
	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

//Stub is replacement for some dependency in your code that will be used during test execution.
//It is typically built for one particular test and unlikely can be reused for another because
//it has hardcoded expectations and assumptions.

//Mock takes stubs to next level. It adds means for configuration,
//so you can set up different expectations for different tests.
//That makes mocks more complicated, but reusable for different tests.

// func TestCreateUserAPI(t *testing.T) {
// 	user, password := randomUser(t)

// 	testCases := []struct {
// 		name          string
// 		body          gin.H
// 		buildStubs    func(store *mockdb.MockStore)
// 		checkResponse func(recoder *httptest.ResponseRecorder)
// 	}{
// 		{
// 			name: "OK",
// 			body: gin.H{
// 				"username":  user.Username,
// 				"password":  password,
// 				"full_name": user.FullName,
// 				"email":     user.Email,
// 			},
// 			buildStubs: func(store *mockdb.MockStore) {
// 				arg := db.CreateUserParams{
// 					Username: user.Username,
// 					FullName: user.FullName,
// 					Email:    user.Email,
// 				}
// 				store.EXPECT().
// 					CreateUser(gomock.Any(), EqCreateUserParams(arg, password)).
// 					Times(1).
// 					Return(user, nil)
// 			},
// 			checkResponse: func(recorder *httptest.ResponseRecorder) {
// 				require.Equal(t, http.StatusOK, recorder.Code)
// 				requireBodyMatchUser(t, recorder.Body, user)
// 			},
// 		},
// 	}
// }

func randomUser(t *testing.T) (user db.Users, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = db.Users{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	return
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.Users) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.Users
	err = json.Unmarshal(data, &gotUser)

	require.NoError(t, err)
	require.Equal(t, user.Username, gotUser.Username)
	require.Equal(t, user.FullName, gotUser.FullName)
	require.Equal(t, user.Email, gotUser.Email)
	require.Empty(t, gotUser.HashedPassword)
}
