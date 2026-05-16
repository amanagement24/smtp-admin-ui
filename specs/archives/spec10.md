@session.go

- SessionStore will hold store for a session for the application
- There will be a cookie on the client called smtpadmin
- The cookie is also the key in the Session map, used to retrieve the right store for the user calling. 

Please implement GetSession function in @session.go as follows:
```
get cookie smtpadmin from request
if present
  search it in the session map
  if found
    retrieve existent instance of the session store
  else 
    generate a new cookie value
    set the new cookie value under smtpadmin
    create key entry in the SessionMap with empty SessionStore
    return newly created SessionStore
  endif
else
    generate a new cookie value
    set the new cookie value under smtpadmin
    create key entry in the SessionMap with empty SessionStore
    return newly created SessionStore
   
endif 

```

- when you set cookie, always http only
