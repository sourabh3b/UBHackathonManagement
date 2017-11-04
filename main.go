package main

import (
	"fmt"
	"github.com/unrolled/render"
	"net/http"
)

//TestRoute - handler to get expense by ID
func TestRoute(w http.ResponseWriter, r *http.Request) {
	render := render.New()

	expenseID := r.URL.Query().Get("expenseID")

	expenseObject, err := model.GetExpense(expenseID)
	if err != nil {
		fmt.Println("Cannot Save expense ", err.Error())
		render.JSON(w, http.StatusBadGateway, "Saved Expense ")
		return
	}
	render.JSON(w, http.StatusOK, expenseObject)
	return
}

func main() {
	fmt.Println("Started UB Hackathon Management....")
	http.HandleFunc("/testRoute", TestRoute)
	http.ListenAndServe(":8889", nil)
}
