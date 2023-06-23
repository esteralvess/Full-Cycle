package entity

type OrderQueue struct {
	Orders []*Order
}

// Less - comparar os valores
func (oq *OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Swap - inverte os valores
func (oq *OrderQueue) Swap(i, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}

// Len - tamanho dos dados que tem
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// Push - adiciona novos dados
func (oq *OrderQueue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// Pop - remove o valor da posição
func (oq *OrderQueue) Pop() any {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}
