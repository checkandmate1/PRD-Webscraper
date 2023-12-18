"Little project I've worked on to learn web scraping, and what better way to do it than to check some routes! \n
Thank you to the Virtual New York ARTCC (nyartcc.org) for providing their preferred route database! 
You just input a departure airport, an arrival airport, and a route, and it will validate it. I am currently working on a better way to print the routes, so for the time being, this is an example print you will see:

{JFK MERIT ORW ORW# BOS 110-FL210 TURBOJET NON-RNAV ZNY ZBW}

The curly braces separate each preferred route.
"JFK" is the departure airport.
"MERIT ORW ORW#" is the route.
"BOS" is the arrival airport.
"110-FL210" is the altitude.
"TURBOJET NON-RNAV" is the aircraft required to fly this route.
"ZNY" and "ZBW" are the departure and arrival ARTCCs, respectively.

Note: When specifying a certain SID/ STAR, use a "#" instead of the number for that SID/ STAR. A fix for this will come in the future.
