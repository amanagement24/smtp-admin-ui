I want to implement add and update mailbox. Please proceed to the following steps:

**step 1** @service.go
@sPlease alongside to the AddMailbox method, add UpdateMailbox. 
The same thing, you should not have two mailboxes with the same name
If somebody edits the mailbox with name inbox, the name changes shall be ignored

**step 2** page
Create a template similar with the existent ones, that would be used for adding / editing the mailbox
There are two buttons at the end, submit and cancel
both trigger POST /editmailbox - see below - but cancel merely goes back to edituser, whereas submit triggers add or edit operation
Title should be either Add or Edit Mailbox depending on the approach
The mailbox edit template shall show the user name just in case, keep user id in a hidden field and allow editing of the name and all the mailbox flags
Create the ViewEditMailbox similar with the the other View types that holds all the data needed to render the edit mailbox template
Create the RenderEditMailbox function

**step3** trigger buttons
in the edituser.html, when the user is in editing mode, you see the list of mailboxes
add a column called edit that for each mailbox renders a hyperlink called "edit" that triggers /editmailbox?id=value - whatever id the mailbox has

**step4** GET controller
Create controller GET /editmailbox and its class
The controller once triggered looks for id in the query parameters
if ID is found, shows the editmailbox.html in edit mode
if no id is found, shows the editmailbox.html in add mode
For adding, create the new mailbox ID as a v7 guid from this controller, so that the editmailbox.html is already populated with the mailbox id even when adding

**step5** POST controller
Create controller POST /editmailbox and its class
Gather the data and either add or edit, depending on operation
Add or edit
If error, show the error and stay on page editmailbox.html
if not error, go back to the edituser.html; show new mailbox in the list at the end, or update existing one in the list. 


