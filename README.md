# DSTUHack2021back

Endpoints:

`/auth/reg` (POST): [email, first_name, last_name, password] => (id, token, name)

`/auth/login` (POST): [email, password]
=> (id, token, name)


`/user/info` (GET) 
=> (id, email, first_name, last_name)

`/user/operations` (GET) 
=> array of (id, name, price, quantity, type, user_id)


`/api/operation` (POST): [type, ticker, price, quantity]
=> (balance)

`/api/tickers` (GET)
=> array of (name, symbol)

`/api/tickers/stocks?symbol=...` (GET): [symbol (as query param)]
=> array of (close, date, high, low, open, symbol, volume)