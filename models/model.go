package model

import(
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
	"github.com/gopalrg310/BitBurst/models"
)

type UserTransaction struct {
	ID            int `json:",omitempty"`
	UserID        string
	Amount        float64
	TransactionID string    `json:",omitempty"`
	Timestamp     time.Time `json:",omitempty"`
}