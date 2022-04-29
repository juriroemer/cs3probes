package outlier

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)


// Takes percentile and warn flags, db connection, log information and checks for outliers in each metric
func HasOutlier(p float64, db *sql.DB, logTimes map[string]int, probe string, warnLimit int, targetId int) map[string]int {

	// Holds the number of entries for a given probe and a given target
	count := countTableRows(db, "probe_" + probe,  targetId)

	// Returns if the number of metrics is less than the provided warnLimit commandline flag
	if (count < warnLimit) {
		return nil
	}

	// Holds index for p-percentile
	pIndex := percentile(count, p)

	// Holds detected outliers
	var outliers map[string]int

	// For each test, check if metrics provided is outlier
	for test, time := range logTimes {
		rows, _ := db.Query(fmt.Sprintf("select %s from %s where targetId = %d order by %s asc limit 1 offset %d-1;", test, "probe_" + probe, targetId, test, pIndex))

		// Holds columns from database probe-table
		cols, err := rows.Columns()
		if err != nil {
			log.Fatal(err.Error())
		}

		// Holds number of columns
		colLen := len(cols)

		// Make and fill hashmap with values of rows query
		vals := make([]interface{}, colLen)
		for rows.Next() {
			for i := 0; i < colLen; i++ {
				vals[i] = new(string)
			}
			err := rows.Scan(vals...)
			if err != nil {
				log.Fatal(err.Error())
			}

			// get p-percentile as cutoff value for each metric
			for i,_ := range vals {
				cutoff, _ := strconv.Atoi(*(vals[i].(*string)))

				// Add metric to outliers, if it is greater than the p-percentile
				if (time > cutoff) {
					if outliers == nil {
						outliers = make(map[string]int)
					}
					outliers[test] = time
				}
			}
		}
	}

	return outliers
}

// Calculate percentile index for number of entries (n) and set percentile value (p)
func percentile(n int, p float64) int {
	index := float64(n) * p
	if index == float64(int(index)) {
		return int(0.5 * (index*2 + 1))
	}
	return int(index) + 1
}

// Count number of rows for a given database table and a given target
func countTableRows(db *sql.DB, tablename string, targetId int) int {
	rows, _ := db.Query(fmt.Sprintf("SELECT count(*) from %s where targetId = %d", tablename, targetId))
	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			fmt.Println(err)
		}
	}
	return count
}