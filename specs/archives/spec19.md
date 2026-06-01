@templates/editdomain.html 
domains.html edit link should point to /editdomain?id= and then id must be provided

@get_editdomain.go
- get session, check logged in like in other controllers
- get id from link
- get domain with this id from session
- if not found, render domains.html with error messge
- if found, prepare viewEditDomain wiht adding false (edit) and with this one render the edit domain
