package orderbook

import (
	"fmt"
	"time"

)

type Order struct {
	ID        int64
	UserID    int64
	Size      float64
	Bid       bool
	Limit     *Limit
	Timestamp int64
}

func NewOrder(bid bool, size float64, userID int64) *Order {
	return &Order{
		UserID:    userID,
		// ID:        int64(rand.Intn(10000000)),
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().UnixNano(),
	}
}

func (o *Order) String() string {
	return fmt.Sprintf("[size: %2f]", o.Size)
}

type Limit struct{
	Price float64
	Orders []*Order
	TotalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}


func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

func (l *Limit) DeleteOrder(o *Order) {
	for i := 0; i < len(l.Orders); i++ {
		if l.Orders[i] == o {
			l.Orders[i] = l.Orders[len(l.Orders)-1]
			l.Orders = l.Orders[:len(l.Orders)-1]
		}
	}

	o.Limit = nil
	l.TotalVolume -= o.Size

	// sort.Sort(l.Orders)
}

// func (l *Limit) Fill(o *Order) []Match {
// 	var (
// 		matches        []Match
// 		ordersToDelete []*Order
// 	)

// 	for _, order := range l.Orders {
// 		if o.IsFilled() {
// 			break
// 		}

// 		match := l.fillOrder(order, o)
// 		matches = append(matches, match)

// 		l.TotalVolume -= match.SizeFilled

// 		if order.IsFilled() {
// 			ordersToDelete = append(ordersToDelete, order)
// 		}
// 	}

// 	for _, order := range ordersToDelete {
// 		l.DeleteOrder(order)
// 	}

// 	return matches
// }

// func (l *Limit) fillOrder(a, b *Order) Match {
// 	var (
// 		bid        *Order
// 		ask        *Order
// 		sizeFilled float64
// 	)

// 	if a.Bid {
// 		bid = a
// 		ask = b
// 	} else {
// 		bid = b
// 		ask = a
// 	}

// 	if a.Size >= b.Size {
// 		a.Size -= b.Size
// 		sizeFilled = b.Size
// 		b.Size = 0.0
// 	} else {
// 		b.Size -= a.Size
// 		sizeFilled = a.Size
// 		a.Size = 0.0
// 	}

// 	return Match{
// 		Bid:        bid,
// 		Ask:        ask,
// 		SizeFilled: sizeFilled,
// 		Price:      l.Price,
// 	}
// }


type Orderbook struct {
	asks []*Limit
	bids []*Limit

	// Trades []*Trade

	// mu        sync.RWMutex
	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
	Orders    map[int64]*Order
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		asks:      []*Limit{},
		bids:      []*Limit{},
		// Trades:    []*Trade{},
		AskLimits: make(map[float64]*Limit),
		BidLimits: make(map[float64]*Limit),
		Orders:    make(map[int64]*Order),
	}
}

//24.29