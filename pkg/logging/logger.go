package logger

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	// "reflect"

	"github.com/Daniel-WWU-IT/cs3probes/pkg/outlier"
	_ "github.com/mattn/go-sqlite3"
)

// Setup SQL statement constant to create target table

const (
	createTargetsTableStatement string = "CREATE TABLE IF NOT EXISTS targets (id integer primary key, host, port)"
)

// Define sqlLogger and log datatypes

type sqlLogger struct {
	db *sql.DB
}

type log struct {
	probe     string
	warnLimit int
	host      string
	port      int
	restimes  map[string]int
}

// Factory for type sqlLogger
func NewLogger() *sqlLogger {
	return &sqlLogger{}
}

// Factory for type log
func NewLog() *log {
	return &log{restimes: make(map[string]int)}
}

// Getters and Setters for type Log

func (l *log) SetProbeName(pn string) {
	l.probe = pn
}

func (l *log) SetWarnLimit(wl int) {
	l.warnLimit = wl
}

func (l *log) SetTarget(t string) {
	l.host = strings.Split(t, ":")[0]
	l.port, _ = strconv.Atoi(strings.Split(t, ":")[1])
}

func (l *log) AddMetric(name string, value int) {
	l.restimes[name] = value
}

func (l log) Probe() string {
	return l.probe
}

func (l log) Host() string {
	return l.host
}

// Opens a new database connection to log database
func (s *sqlLogger) connect() (*sql.DB, error) {
	_ = os.MkdirAll("./data/logs/", os.ModePerm)
	return sql.Open("sqlite3", "./data/logs/logs.db")
}

// Takes a log Type and inserts it into the database, also checks for outliers
func (s *sqlLogger) InsertLog(l *log, percentile int) map[string]int {
	// Holds primary key to target system
	var targetId int

	// Connect to database
	db, err := s.connect()
	if err != nil {
		fmt.Println(err)
	}

	// Create targets table, if it does not exists
	db.Exec(createTargetsTableStatement)

	// Get primary key for target system, creates db entry for target, if it doesn't exists
	selectTargetIdStatement := fmt.Sprintf("SELECT id FROM targets WHERE host = '%s' and port = %d ", l.host, l.port)
SELECTTARGET:
	rows, _ := db.Query(selectTargetIdStatement)

	for rows.Next() {
		rows.Scan(&targetId)
	}
	if targetId == 0 {
		insertTargetStatement := fmt.Sprintf("INSERT INTO targets (host, port) values ('%s', %d) ", l.host, l.port)
		db.Exec(insertTargetStatement)
		goto SELECTTARGET
	}

	// Generate SQL statements
	createProbesTableStatement, insertStatement := s.prepareSqlStatements(l.probe, l.restimes, targetId)

	// Create probe-table if it doesn't exist, insert log into database
	db.Exec(createProbesTableStatement)
	db.Exec(insertStatement)

	// check log for outliers and return them
	outliers, _ := outlier.HasOutlier(float64(percentile)/100, db, l.restimes, l.probe, l.warnLimit, targetId)
	return outliers
}

// Generates "CREATE TABLE IF NOT EXISTS" and "INSERT" sql-commands according to the probe and column names provided in the log
// Types in "CREATE TABLE" statements are optional in sqlite3
func (s *sqlLogger) prepareSqlStatements(probename string, restimes map[string]int, targetId int) (string, string) {
	var cols string
	var times string
	// Generates strings from testnames and testresults
	for col, t := range restimes {
		cols += ", " + col
		times += ", " + strconv.Itoa(t)
	}

	createProbesTableStatement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY, timestamp INTEGER, targetId%s)", "probe_"+probename, cols)
	insertStatement := fmt.Sprintf("INSERT INTO %s (timestamp,targetId%s) VALUES (%d, %d %s)", "probe_"+probename, cols, time.Now().Unix(), targetId, times)
	fmt.Println(insertStatement)
	return createProbesTableStatement, insertStatement
}
