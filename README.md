# SETUP

1. cd server
2. make air
3. cd web
4. yarn dev

# TODO

- Pagination
  fuzzy search title
  -created at DESC (default) ASC
  -followers DESC
  -checkpoints DESC ASC
- Create user profile page
- User to user follow
- Fuzzy search
- Create index for fields used in order by (https://leopard.in.ua/2014/10/11/postgresql-paginattion#.YN4ynehKjb0)
- Remove password from return value
- Implement testing
- Query complexity

# EXTRA

- Add docker
- Add docker hot reload
- Add testing
- Add continous integration
- Split into microservices

# PAGES

/
/explore
/login
/register
/create/map
/map/[id]
/user/[id]
/checkpoint/[id]?

# TECHNOLOGIES

- React
- Nextjs
- Typescript
- Tailwindcss
- Apollo
- GraphQL
- GraphQL Code Generator
- Go
- Gqlgen
- Dataloader
- Testify
- Testfixtures
- Cypress - replace with Jest?
