package utils

import (
	"context"

	"fmt"
	"github.com/gopalrg310/bitburst/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

var Log = logrus.New()

func ResponseHandler(rw http.ResponseWriter, reqLog *logrus.Entry, serviceName string, contentType interface{}, status interface{}, response string, code int, loggerInfo string, loggerInfoObj interface{}, loggerErr string, err interface{}) {

	defer RecoverFunc("ResponseHandler")
	switch k := contentType.(type) {
	case string:
		if k != "" {
			rw.Header().Set("Content-Type", "application/"+k)
		}
	default:

	}
	switch k := status.(type) {
	case int:
		if k != 0 {
			rw.WriteHeader(k)
		}
	default:

	}

	io.WriteString(rw, response)

	if reqLog != nil {
		if err != nil && loggerErr == "" {
			reqLog.Error(err)
		}
		if loggerInfoObj != nil && loggerInfo == "" {
			reqLog.Info(loggerInfoObj)
		}
		if loggerInfo != "" {
			if loggerInfoObj != nil {
				reqLog.Info(loggerInfo, loggerInfoObj)
			} else {
				reqLog.Info(loggerInfo)
			}
		}
		if loggerErr != "" {
			if err != nil {
				reqLog.Errorf("%v %v", loggerErr, err)
			} else {
				reqLog.Error(loggerErr)
			}
		}
	}
}
func RecoverFunc(f string) {
	requestLogger := Log.WithFields(logrus.Fields{"func": f})
	if r := recover(); r != nil {
		requestLogger.Error("Recovered : ", r)
		requestLogger.Error("Tracing Issue : ", string(debug.Stack()))
	}
}
func InitialRecover(rw http.ResponseWriter, submodule, serviceName string) {
	requestLogger := Log.WithFields(logrus.Fields{"func": serviceName, "module": submodule})
	if r := recover(); r != nil {
		rw.WriteHeader(http.StatusNotFound)
		rs := fmt.Sprintf("%s", r)
		io.WriteString(rw, rs)
		requestLogger.Error("Recovered in : ", r)
		requestLogger.Error("Tracing Issue : ", string(debug.Stack()))
	}
}
func FormJsonOutput(output string) string {
	return fmt.Sprintf("{\"Message\":\"%v\"}", output)
}
func AddTransaction(db *pgxpool.Pool, transaction models.UserTransaction) error {
	statement := `INSERT INTO transactions (user_id, amount, transaction_id, Created_At) 
                  VALUES ($1, $2, $3, $4) ON CONFLICT (user_id, transaction_id) DO NOTHING`

	// Execute the SQL statement
	_, err := db.Exec(context.Background(), statement, transaction.UserID, transaction.Amount, transaction.TransactionID, time.Now().UTC())
	if err != nil {
		return err
	}
	return nil
}
func ConnectDB(nooftry int, databaseUrl string) (*pgxpool.Pool, context.Context) {
	nooftry += 1
	ctx := context.Background()
	dbPool, err := pgxpool.Connect(ctx, databaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		time.Sleep(1 * time.Minute)
		return ConnectDB(nooftry, databaseUrl)
	}
	if nooftry > 5 {
		os.Exit(1)
	}
	return dbPool, ctx
}
