@post_domains.go

- implement processDeleteDomains
- SelectedDomains in session must have at least one item, if not, show error (view header error field) and stay in domains page
- If there are selected domains, get a list from the session.Domains with the ids in SelectedDomains, assemble the ViewCdDomains and renderCdDomains page