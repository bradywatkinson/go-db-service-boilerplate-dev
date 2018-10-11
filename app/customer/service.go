// customer/service.go

package customer

import (
  "database/sql"
  "net/http"
  "strconv"
  "encoding/json"
  "fmt"

  "github.com/gorilla/mux"

  "app/common"
)


func createCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    var c customer
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&c); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid request payload: %s", err))
      return
    }
    defer r.Body.Close()

    if err := c.createCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusCreated, c)
  }
}

func getCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
    }

    c := customer{ID: id}
    if err := c.getCustomer(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Customer not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    common.RespondWithJSON(w, http.StatusOK, c)
  }
}

func updateCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
      return
    }

    c := customer{ID: id}
    if err := c.getCustomer(a.DB); err != nil {
      switch err {
      case sql.ErrNoRows:
        common.RespondWithError(w, http.StatusNotFound, "Customer not found")
      default:
        common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      }
      return
    }

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&c); err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
      return
    }
    defer r.Body.Close()

    if err := c.updateCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, c)
  }
}

func deleteCustomer(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
      common.RespondWithError(w, http.StatusBadRequest, "Invalid Customer ID")
      return
    }

    c := customer{ID: id}
    if err := c.deleteCustomer(a.DB); err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
  }
}

func getCustomers(a *common.App) func(http.ResponseWriter, *http.Request) {
  return func(w http.ResponseWriter, r *http.Request) {
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
      count = 10
    }
    if start < 0 {
      start = 0
    }

    var c customer
    customers, err := c.getCustomers(a.DB, start, count)
    if err != nil {
      common.RespondWithError(w, http.StatusInternalServerError, err.Error())
      return
    }

    common.RespondWithJSON(w, http.StatusOK, customers)
  }
}