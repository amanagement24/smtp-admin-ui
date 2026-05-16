- add value in config.json for session timeout in seconds, called sessionTimeout
- add it to ConfigData 
 
- put time.Time in SessionStore
- when it is created, will be time.Now
- every time you access the session you look to see if difference between current time and stored SessionStore time is bigger than sessionTimeout from config data
- if it is bigger, you clear the session - there is a function for that - and return to login page with "session expired" 
