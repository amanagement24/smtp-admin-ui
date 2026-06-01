This might be a little bit tough

In the database, the users are admins or not. The checkLoggedIn was changed so that only admin users 
can do most of the things. 

Please do the following changes:
If the user is not admin:
- "Domains" should not be shown on the menu
- Already the admin bound endpoints cannot be accessed if user is not admin - took care of that
- The application, after successful login, should proceed directly to change password for the currently logged in user
- you cannot use get chpass controller for that because there will be no call in this direction
- post_login.go should be changed, checkLoggedIn must not be with "true" parameter, for admin, must be false
- however check, if the post_login is called by the non admin, the operation can only be done by the currently logged in user. 
- after error, success, no matter, if the user is not admin, the application stays in the chpass page
- if error, show error and stay in page
- if success, I placed the following in the page

<h2>The password has been successfully updated</h2>

This should be visible only upon successful change and it is request bound, not session, so next time should show only another successful password change

