@service.go

- create function AddMailbox which adds new mailbox to the database, mailbox table
- table defined in docs/ddl.sql
- create function UpdateMailbox which updates mailbox
- if mailbox name is equals ignore case with inbox, force it to INBOX all capitals
- you cannot have two mailboxes with name INBOX under the same user

primary key MailboxID is provided from outside service.go, most likely will be a v7 guid
