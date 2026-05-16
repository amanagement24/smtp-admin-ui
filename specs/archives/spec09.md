- I created extract.go
- Create an entry for each of the templates we got, example for the domains

```
type PostDomains {
    Selected []string
}
```

The above will be used to collect whatever is in request after somebody posts the domains.html form. Create
something like that for everything, pay attention: only for the form fields! For example viewDomain.html has
a lot of users in a list, the list does not get triggered, however the checkboxes associated with them do; 

**Please** do as follows: 

- Create Post* struct for each of the templates
- In extract.go create for each template a function called Extract* for example ExtractLogin etc that parses the request and returns a Post* for example PostLogin with the form data at the ready
