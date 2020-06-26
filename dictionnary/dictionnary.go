package dictionnary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

//Dictionnary => Dictionnary
type Dictionnary struct {
	db *badger.DB
}

//Entry => Entry
type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created) //Le nombre derrière le % indique le nombre de caractère alloué pour cette valeur, le \t signifie un dab = tabulation
}

//New => New
func New(dir string) (*Dictionnary, error) {
	opts := badger.DefaultOptions(dir)
	opts.Dir = dir
	opts.ValueDir = dir

	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	dict := &Dictionnary{
		db: db,
	}
	return dict, nil
}

//Close => Close
func (d *Dictionnary) Close() {
	d.db.Close()
}
