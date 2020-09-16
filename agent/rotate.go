package agent

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/moco"
)

func (a *Agent) RotateLog(w http.ResponseWriter, r *http.Request) {
	errFile := filepath.Join(moco.VarLogPath, moco.MySQLErrorLogName)
	_, err := os.Stat(errFile)
	if err == nil {
		err := os.Rename(errFile, errFile+".0")
		if err != nil {
			log.Error("failed to rotate err log file", map[string]interface{}{
				log.FnError: err,
			})
			internalServerError(w, fmt.Errorf("failed to rotate err log file: %w", err))
			return
		}
	} else if !os.IsNotExist(err) {
		log.Error("failed to stat err log file", map[string]interface{}{
			log.FnError: err,
		})
		internalServerError(w, fmt.Errorf("failed to stat err log file: %w", err))
		return
	}

	slowFile := filepath.Join(moco.VarLogPath, moco.MySQLSlowLogName)
	_, err = os.Stat(slowFile)
	if err == nil {
		err := os.Rename(slowFile, slowFile+".0")
		if err != nil {
			log.Error("failed to rotate slow query log file", map[string]interface{}{
				log.FnError: err,
			})
			internalServerError(w, fmt.Errorf("failed to rotate slow query log file: %w", err))
			return
		}
	} else if !os.IsNotExist(err) {
		log.Error("failed to stat slow query log file", map[string]interface{}{
			log.FnError: err,
		})
		internalServerError(w, fmt.Errorf("failed to stat slow query log file: %w", err))
		return
	}

	buf, err := ioutil.ReadFile(moco.MiscPasswordPath)
	if err != nil {
		internalServerError(w, fmt.Errorf("failed to read misc passsword file: %w", err))
		return
	}
	password := strings.TrimSpace(string(buf))

	podName := os.Getenv(moco.PodNameEnvName)

	db, err := a.acc.Get(fmt.Sprintf("%s:%d", podName, moco.MySQLAdminPort), moco.MiscUser, password)
	if err != nil {
		internalServerError(w, fmt.Errorf("failed to get database: %w", err))
		return
	}

	if _, err := db.ExecContext(r.Context(), "FLUSH LOCAL ERROR LOGS;\nFLUSH LOCAL SLOW LOGS;\n"); err != nil {
		internalServerError(w, fmt.Errorf("failed to exec mysql FLUSH: %w", err))
		return
	}
}
