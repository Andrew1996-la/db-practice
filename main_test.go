package main

import (
	"database/sql"
	"db-practice/clients"
	"testing"

	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := 1

	// напиши тест здесь
	client, err := clients.SelectById(db, clientID)

	require.NotEmpty(t, err)

	require.Equal(t, client.Id, clientID)
	require.NotEmpty(t, client.FIO)
	require.NotEmpty(t, client.Login)
	require.NotEmpty(t, client.Birthday)
	require.NotEmpty(t, client.Email)
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	clientID := -1

	// напиши тест здесь
	client, err := clients.SelectById(db, clientID)
	require.Equal(t, sql.ErrNoRows, err)

	require.Empty(t, client.FIO)
	require.Empty(t, client.Login)
	require.Empty(t, client.Birthday)
	require.Empty(t, client.Email)
}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := clients.Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.Id, err = clients.Insert(db, cl)
	require.NoError(t, err)
	require.NotZero(t, cl.Id)

	client, err := clients.SelectById(db, cl.Id)
	require.NoError(t, err)

	require.Equal(t, cl.Id, client.Id)
	require.Equal(t, cl.Login, client.Login)
	require.Equal(t, cl.Email, client.Email)
	require.Equal(t, cl.Birthday, client.Birthday)
	require.Equal(t, cl.FIO, client.FIO)
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	require.NoError(t, err)
	defer db.Close()

	cl := clients.Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// напиши тест здесь
	cl.Id, err = clients.Insert(db, cl)
	require.NoError(t, err)
	require.NotZero(t, cl.Id)

	client, err := clients.SelectById(db, cl.Id)
	require.NoError(t, err)

	err = clients.DeleteClient(db, client.Id)
	require.NoError(t, err)

	client, err = clients.SelectById(db, cl.Id)
	require.Equal(t, sql.ErrNoRows, err)
}
