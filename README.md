# supermarket-checkout
Simple Supermarket Checkout problem solution, this repository is to solve the following problem statement.

```
Implement the code for a checkout system that handles pricing schemes such as "pineapples cost 50, three pineapples cost 130."

Implement the code for a supermarket checkout that calculates the total price of a number of items. In a normal supermarket, things are identified using Stock Keeping Units, or SKUs. In our store, we’ll use individual letters of the alphabet (A, B, C, and so on). Our goods are priced individually. In addition, some items are multi-priced: buy n of them, and they’ll cost you y pence. For example, item A might cost 50 individually, but this week we have a special offer: buy three As and they’ll cost you 130. In fact the prices are:

SKU | Unit Price | Special Price

A | 50 | 3 for 130

B | 30 | 2 for 45

C | 20 | 

D | 15 | 

The checkout accepts items in any order, so that if we scan a B, an A, and another B, we’ll recognize the two Bs and price them at 45 (for a total price so far of 95). The implementation should consider if the pricing model may change frequently.
```


### How to Run:
Use the following command in a terminal from the root of this directory
`go run ./cmd/main.go testFile.txt`