@internal/endpoints/post_login.go

- check login, etc, like the other ones
- two situations if cancel is pushed, go back to viewDomains

- if submit is pushed then
1. check both password and repeat are filled out
2. check they are identical
3. if they are identical, then update user by the id with this new password which shall be hashed with service.HashPassword
4. go back to viewDomain

