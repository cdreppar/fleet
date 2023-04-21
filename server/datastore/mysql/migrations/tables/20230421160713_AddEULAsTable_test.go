package tables

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUp_20230421160713(t *testing.T) {
	db := applyUpToPrev(t)
	applyNext(t, db)

	insertStmt := `
          INSERT INTO eulas (token, name, bytes)
	  VALUES (?, ?)
	`

	selectStmt := `
	  SELECT token, bytes 
	  FROM eulas
	  WHERE token = ?
	`

	_, err := db.Exec(insertStmt, "ABC-DEF", "eula.pdf", []byte("eula"))
	require.NoError(t, err)

	_, err = db.Exec(insertStmt, "ABC-DEF", "eula_2.pdf", []byte("eula_2"))
	require.ErrorContains(t, err, "Error 1062")

	_, err = db.Exec(insertStmt, "GHI-JLK", "eula_2.pdf", []byte("eula_2"))
	require.NoError(t, err)

	var (
		token string
		name  string
		bytes []byte
	)

	err = db.QueryRow(selectStmt, 1).Scan(&name, &bytes, &token)
	require.NoError(t, err)
	require.Equal(t, "ABC-DEF", token)
	require.Equal(t, "eula.pdf", name)
	require.Equal(t, []byte("eula"), bytes)
}
