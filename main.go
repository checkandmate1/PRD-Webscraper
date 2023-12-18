package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"bufio"
	"os"
)


type Flight struct {
	Departure, Route, Destination, Altitude, Type, DepartureARTCC, DestinationARTCC string
}
var DepartureAirport, ArrivalAirport, AircraftRoute string
var Number int

func main() {
	fmt.Println("Input Departure")
	fmt.Scanln(&DepartureAirport)
	fmt.Println("Input Arrival")
	fmt.Scanln(&ArrivalAirport)
	fmt.Println("Input Route")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		AircraftRoute = scanner.Text()
	}
	fmt.Println("Working...")
	scraper()
}

func scraper() {
	Link := fmt.Sprintf("https://nyartcc.org/prd?from=%s&to=%s", DepartureAirport, ArrivalAirport)

	var correctRoutes []Flight

	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()


	var flightNodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(Link), 
		chromedp.Nodes("tr.tier1", &flightNodes, chromedp.ByQueryAll),
	)
	if err != nil {
		log.Fatal("Error:", err)
	}

	
	for _, node := range flightNodes {
		var departure, route, destination, altitude, flightType, departureARTCC, destinationARTCC string

	
		err := chromedp.Run(ctx,
			chromedp.Text("td:nth-child(1)", &departure, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("td:nth-child(2)", &route, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("td:nth-child(3)", &destination, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("td:nth-child(6)", &altitude, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("td:nth-child(7)", &flightType, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("td:nth-child(8)", &departureARTCC, chromedp.ByQuery, chromedp.FromNode(node)),
			chromedp.Text("td:nth-child(9)", &destinationARTCC, chromedp.ByQuery, chromedp.FromNode(node)),
		)

		if err != nil {
			log.Fatal("Error:", err)
		}

		
		flight := Flight{
			Departure:        departure,
			Route:         route,
			Destination:      destination,
			Altitude:         altitude,
			Type:             flightType,
			DepartureARTCC:   departureARTCC,
			DestinationARTCC: destinationARTCC,
		}
		
		if flight.Route == AircraftRoute {
			Number = Number + 1000
			fmt.Println("Correct Route!")
			main()
		} else {
			Number = Number - 1
			correctRoutes = append(correctRoutes, flight)
		}
		

	}
	if Number < 0 {
		fmt.Println("Incorrect Route. These are the correct ones: \n", correctRoutes)
		main()
	}

	
}
