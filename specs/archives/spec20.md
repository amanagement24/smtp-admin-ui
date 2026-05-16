@internal/ui/templates/cddomains.html this show the domains that are about to be delete for confirmation
It triggers post in post_cddomains.go - please ensure this

If cancel is present in form fields, just render domains back and that's it

@service.go create a method that delete domains provided you got domain ids. Forget about referential integrity: let it fail, and then collect the error and say it failed because there are dependent items.

if submit is present in form fields, go ahead and call method in service.go that deletes domains
if failed, stay in cddomains.html page but show the error message
if succeeded, call GetDomains in service, update the domains in SessionStore, replace the SelectedDomains with nil and then render domains page

