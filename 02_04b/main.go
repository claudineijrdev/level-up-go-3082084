package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// the amount of bidders we have at our auction
const bidderCount = 10

// initial wallet value for all bidders
const walletAmount = 250

// items is the map of auction items
var items = []string{
	"The \"Best Gopher\" trophy",
	"The \"Learn Go with Adelina\" experience",
	"Two tickets to a Go conference",
	"Signed copy of \"Beautiful Go code\"",
	"Vintage Gopher plushie",
}

// bid is a type that pairs the bidder id and the amount they want to bid
type bid struct {
	bidderID string
	amount   int
}

// auctioneer receives bids and announces winners
type auctioneer struct {
	bidders map[string]*bidder
}

// runAuction and manages the auction for all the items to be sold
// Change the signature of this function as required
func (a *auctioneer) runAuction(bids <-chan bid, open chan<- bool) {
	for _, item := range items {
		maxBid := 0
		var bidOwner bid
		a.openBinds(open)
		log.Printf("Opening bids for %s!\n", item)
		for range bidderCount {
			bid := <-bids
			if bid.amount > maxBid {
				bidOwner = bid
			}
		}
		log.Printf("The winner of the item %s is %s by the value of %d", item, bidOwner.bidderID, bidOwner.amount)
		a.bidders[bidOwner.bidderID].payBid(bidOwner.amount)
	}
}

func (a *auctioneer) openBinds(open chan<- bool) {
	for range bidderCount {
		open <- true
	}
}

// bidder is a type that holds the bidder id and wallet
type bidder struct {
	id     string
	wallet int
}

// placeBid generates a random amount and places it on the bids channels
// Change the signature of this function as required
func (b *bidder) placeBid(bids chan bid, open chan bool) {
	for range items {
		<-open
		amount := getRandomAmount(b.wallet)
		bid := bid{
			bidderID: b.id,
			amount:   amount,
		}
		bids <- bid
	}
}

// payBid subtracts the bid amount from the wallet of the auction winner
func (b *bidder) payBid(amount int) {
	b.wallet -= amount
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to the LinkedIn Learning auction.")
	bidders := make(map[string]*bidder, bidderCount)
	bids := make(chan bid, bidderCount)
	open := make(chan bool)
	for i := 0; i < bidderCount; i++ {
		id := fmt.Sprint("Bidder ", i)
		b := bidder{
			id:     id,
			wallet: walletAmount,
		}
		bidders[id] = &b
		go b.placeBid(bids, open)
	}
	a := auctioneer{
		bidders: bidders,
	}
	a.runAuction(bids, open)
	log.Println("The LinkedIn Learning auction has finished!")
}

// getRandomAmount generates a random integer amount up to max
func getRandomAmount(max int) int {
	return rand.Intn(int(max))
}
