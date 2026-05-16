@internal/service/service.go

- implement AddDomain - basically add domain with provided id. 
- Convert domain (the name) to lowercase.
- Domain must be filled out
- At least one dot - domain is whatever.something at least, can be more than one dot, however at least one dot
- Check to see if it exists in the db already, and, if so, return error. 
- catchall_ind will be 'N' and  catchall_login will be empty string
 