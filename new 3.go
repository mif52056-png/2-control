type Order struct {
	ID          int
	Items       []int
	Total       float64
	Address     string
	IsCompleted bool
}

func addOrder(orders map[int]Order, order Order) {
	orders[order.ID] = order
}