@internal/endpoints/get_viewdomain.go do as follows:

- get session and checkLoggedIn like in the other ones
- expect id as a query parameter, this is the domain id
- get the id, load the domain and the users by domain with respective methods in service
- if there isn't domain for the id, then render domains.html with error message
- if domain exists, then assemble ViewDomain and render viewdomain.html
