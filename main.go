package main

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Country struct {
	idx  int
	name string
}

func (c *Country) SetName(name string) {
	c.name = name
}

var dsn = "landmark:landmark@csmsu@tcp(202.28.34.197:3306)/landmark"

func getConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()
	println("connection successful")
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	return db, err
}

// CRUD
func GetCountries() ([]Country, error) {
	db, err := getConnection()
	if err != nil {
		panic(err.Error())
	}
	query := "select * from country"
	rows, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	countries := []Country{}

	for rows.Next() {
		// idx := 0
		// name := ""
		country := Country{}
		err = rows.Scan(&country.idx, &country.name)
		if err != nil {
			panic(err.Error())
		}
		countries = append(countries, country)
	}
	return countries, err
}

func GetCountryById(idx int) (Country, error) {
	db, err := getConnection()
	if err != nil {
		panic(err.Error())
	}
	query := "select * from country where idx = ?"
	row := db.QueryRow(query, idx)
	var country Country
	err = row.Scan(&country.idx, &country.name)
	if err != nil {
		panic(err.Error())
	}
	return country, nil

}

func GetCountryByName(name string) ([]Country, error) {
	db, err := getConnection()
	if err != nil {
		panic(err.Error())
	}
	query := "select * from country where name like ?"
	rows, err := db.Query(query, "%"+name+"%")

	if err != nil {
		panic(err.Error())
	}

	countries := []Country{}

	for rows.Next() {
		// idx := 0
		// name := ""
		country := Country{}
		err = rows.Scan(&country.idx, &country.name)
		if err != nil {
			panic(err.Error())
		}
		countries = append(countries, country)
	}
	return countries, err
}

func AddCountry(country Country) (int64, int64, error) {
	db, err := getConnection()
	if err != nil {
		return -1, -1, err
	}
	query := "insert into country(name) values(?)"
	result, err := db.Exec(query, country.name)
	if err != nil {
		return -1, -1, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, -1, err
	}

	return affected, id, nil
}

func UpdateCountry(country Country) (int64, error) {
	db, err := getConnection()
	if err != nil {
		return -1, err
	}
	query := "update country set name = ? where idx = ?"
	result, err := db.Exec(query, country.name, country.idx)
	if err != nil {
		return -1, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return affected, nil
}

func DeleteCountry(idx int) (int64, error) {
	db, err := getConnection()
	if err != nil {
		return -1, err
	}
	query := "delete from country where idx = ?"
	result, err := db.Exec(query, idx)
	if err != nil {
		return -1, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return -1, err
	}

	return affected, nil
}

func main() {
	// countries, err := GetCountries()
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(countries)

	// country, err := GetCountryById(30)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(country)

	// countries, err := GetCountryByName("อ")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// result := countries
	// fmt.Println(result)

	// country := Country{}
	// country.SetName("New Country")
	// affected, idx, err := AddCountry(country)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(affected, idx)

	// country := Country{idx: 175, name: "Super New Country"}
	// affected, err := UpdateCountry(country)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(affected)

	affected, err := DeleteCountry(175)
	if err != nil {
		panic(err.Error())
	}
	println(affected)
}
