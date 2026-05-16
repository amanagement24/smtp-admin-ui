@internal/service/service.go

- implement UpdateDomain 
- Convert domain (the name) to lowercase.
- Domain must be filled out
- At least one dot - domain is whatever.something at least, can be more than one dot, however at least one dot
- Check to see if it exists in the db already under an id different than current one, and, if so, return error. 
- if catchall is filled out, catchalllogin must be filled out
- then look for catchalllogin under that particular domain and it must exist, or else error
 