package pocket2rm

import (
	"database/sql"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

// NullTime is used to indicate an unset time.
var NullTime = time.Unix(0, 0)

func (p *Pocket2RM) openDatabase() {
	database, err := sql.Open("sqlite3", DatastoreFile)
	if err != nil {
		panic(err)
	}

	p.db = database
	statement, _ := database.Prepare(
`
CREATE TABLE IF NOT EXISTS articles
(
   ItemID         TEXT PRIMARY KEY,
   ResolvedURL    TEXT,
   ResolvedTitle  TEXT,
   DateFromPocket TEXT,
   DatePushedToRM TEXT,
   MercuryData    TEXT
)
`,
	)

	statement.Exec()
}

func (p *Pocket2RM) isArticleKnown(itemID string) time.Time {

	count := p.db.QueryRow("SELECT COUNT(*) FROM articles WHERE ItemID = ?", itemID)

	var itemTime time.Time
	err := count.Scan(&itemTime)
	if err == sql.ErrNoRows {
		return NullTime
	}

	var cnt int
	count.Scan(&cnt)
	if cnt == 0 {
		return NullTime
	}
	
	return itemTime
}

func (p *Pocket2RM) listAllArticles() ([]string, error) {
	rows, err := p.db.Query("SELECT ItemID FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string

	for rows.Next() {
		var itemID string
		err = rows.Scan(&itemID)
		if err == nil {
			items = append(items, itemID)
		}
	}

	return items, nil
}
