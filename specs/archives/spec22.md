implement post /edituser

- check login is filled out properly
- no matter add / update, check before operating there is no user with same login but different id
- cancel goes back to viewdomain
- update will look to see if adding or updating and will proceed accordingly

- after adding, add the newly created user to the Users in viewdomain and then render view domaoin
- after updating, refresh the whole list of Users to be rendered in viewdomain and then render the view domain
