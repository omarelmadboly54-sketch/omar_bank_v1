/* 
  --------------------------------------------------
  Note: I made this database work and wrote the 
  migrations according to sections 4, 5. 
  I also implemented the transfer and entry SQL 
  queries entirely by myself as homework assigned 
  by the mentor on YouTube.
  --------------------------------------------------
*/

-- name: CreateAccount :one
INSERT INTO accounts(
    owner,
    balance,
    currency
) VALUES (
    $1,$2,$3
)RETURNING *;

-- name: GetAccount :one 
SELECT * FROM accounts
WHERE id =$1 LIMIT 1;


/* 
  --------------------------------------------------
  Note: I made this modification according to section 7
  the modification:get accont for update
  --------------------------------------------------
*/

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id=$1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many 
SELECT * FROM accounts
WHERE owner=$1
ORDER BY id
LIMIT $2
OFFSET $3;


-- name: UpdateAccount :one
UPDATE accounts
SET balance =$2
WHERE id=$1
RETURNING *;

/* 
  --------------------------------------------------
  Note: I made this modification according to section 7
  the modification:addaccountbalance
  --------------------------------------------------
*/

-- name: AddAccoountBalance :one
UPDATE accounts
SET balance=balance+sqlc.arg(amount)
WHERE id=sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;


