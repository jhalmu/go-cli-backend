# go-cli-backend

Little paractice app vol 2 with Go, Gorm, Postgresql and of cource Cobra with Pterm.

## What we have

Db is remote so in theory You can take this app to Go. Pun intended. You can add and list posts. Content area is multiline. We have Slug!

## What we don't have

- Things like update or delete post nor show one by id or slug
- We can not now add own db where to add posts or save possible other configurations.

## What I learned

- Cobra
- Pterm
- Gorm
- Postgresql
- Remote db

## What I want to learn next

- [ ] Tests
- [ ] Update and
- [ ] Delete posts
- [ ] Show one post by id or slug
- [ ] Add tags
- [ ] DRY w/ refactoring
- [ ] Config
- [ ] Add more fields like author etc...
- [ ] Add pagination
- [ ] Add search by tags
- [ ] Using flags or subcommands
- [ ] Add SQLC for working with database(s)
- [ ] Add logging
- [ ] Add CI/CD
- [ ] Add Dockerfile (not really needed or...)
- [ ] Add API
- [ ] Add frontend(s) <3

## End words

Yes. I tryed to move database functionalities to funcs in separate file but there was wierd errors with Gorm, so, I decided to leave it as is - for now. I will fix this in n project.

Using flags or subcommands should be better user experience when using commands where users are used to use them.I want to learn how to do it properly and make it more user friendly.

## Well

How about securing env? I have tryed to use embedding but there was some issues with it so I decided to leave it as is for now. But build seems to work in differend machines when I tryed. That's nice.
