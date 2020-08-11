package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"walmart/db"
	"walmart/model"
)

type App struct {
	d        db.DB
	handlers map[string]http.HandlerFunc
}

func NewApp(d db.DB, cors bool) App {
	app := App{
		d:        d,
		handlers: make(map[string]http.HandlerFunc),
	}

	productByIdHandler := app.GetProductById
	if !cors {
		productByIdHandler = disableCors(productByIdHandler)
	}

	productsByTokenHandler := app.GetProductsByToken
	if !cors {
		productsByTokenHandler = disableCors(productsByTokenHandler)
	}

	app.handlers["/api/byId/"] = productByIdHandler
	app.handlers["/api/byToken/"] = productsByTokenHandler
	app.handlers["/"] = http.FileServer(http.Dir("/webapp")).ServeHTTP
	return app
}

func (a *App) Serve() error {
	for path, handler := range a.handlers {
		http.Handle(path, handler)
	}
	log.Println("Web server is available on port 8080")
	return http.ListenAndServe(":8080", nil)
}

func (a *App) GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := a.d.FindByID(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if product == nil {
		RespondWithError(w, http.StatusNotFound, "Product not found")
		return
	}

	if IsPalindrome(strconv.Itoa(id)) {
		product.Price = product.Price / 2
	}
	RespondWithJSON(w, http.StatusOK, product)
}

func (a *App) GetProductsByToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := Vars(r)
	token, present := vars["token"]
	if !present {
		RespondWithError(w, http.StatusBadRequest, "Empty input")
		return
	}

	var productsByBrand model.Products
	var productsByDescription model.Products
	var errBrand, errDescription error

	productsByBrand, errBrand = a.d.FindByBrand(token)
	productsByDescription, errDescription = a.d.FindByDescription(token)
	if errBrand != nil || errDescription != nil {
		RespondWithError(w, http.StatusInternalServerError, errBrand.Error())
		return
	}
	if len(productsByBrand) == 0 && len(productsByDescription) == 0 {
		RespondWithError(w, http.StatusNotFound, "No product found for such token")
		return
	}

	products := append(productsByBrand, productsByDescription...)
	_ = products.ToJSON(w)
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}

// Needed in order to disable CORS for local development
func disableCors(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		h(w, r)
	}
}

type contextKey int

const (
	varsKey contextKey = iota
	routeKey
)

// Vars returns the route variables for the current request, if any.
func Vars(r *http.Request) map[string]string {
	if rv := r.Context().Value(varsKey); rv != nil {
		return rv.(map[string]string)
	}
	return nil
}
