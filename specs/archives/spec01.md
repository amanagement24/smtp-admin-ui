- ui/templates there are templates
- template is loaded by loading header.html and then each of the other templates so when you exec you header, then the other one and one piece of data
- load all those templates so that they are embedded in code
- create struct ViewHeader 

type ViewHeader struct {
    Login string
    Error string    
}

and attach to it a method IsLoggedIn that returns true if len(login) is bigger than zero

- for login.html template create a struct ViewLogin that embeds an instance of ViewHeader and then another attribute Login that will be used to fill out the field in the page

- create a function RenderLogin(w, viewLogin) that renders login on the screen, w is the http writer, viewLogin is the ViewLogin struct and if you have all the data, please populate the login.html template
- 