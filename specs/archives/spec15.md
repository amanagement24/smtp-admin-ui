@service.go

- implement getuserbylogin
- login is user@domain like info@mltt.site where info is in User table and mltt.site is the name of the domain in domain table
- find the user and if so load it in user value object and return it, if not, return nil but do not throw error because you don't find the user
- of course, for any other error, you will return that error
