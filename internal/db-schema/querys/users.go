package querys

const UsersGet = `select * from users`
const UsersInsert = `insert into users (name, username, password, groupid) values (?,?,?,?,?)`
