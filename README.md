# DSTUHack2021back

Endpoints:

`/auth/reg` (POST): [email, first_name, last_name, password] => (id, token, name)

`/auth/login` (POST): [email, password]
=> (id, token, name)


`/user/info` (GET) 
=> (id, email, first_name, last_name, balance)

`/user/operations` (GET) 
=> array of (id, name, price, quantity, type, user_id)

`/user/portfolio` (GET)
=> map with symbols as keys and quantity as values


`/api/operation` (POST): [type, ticker, price, quantity]
=> (balance)

`/api/stocks?page=...` (GET)
=> (balance, array of [name, symbol, close, diference])

`/api/tickers` (GET)
=> array of (name, symbol)

`/api/tickers/stocks?symbol=...` (GET): [symbol (as query param)]
=> array of (close, date, high, low, open, symbol, volume)
