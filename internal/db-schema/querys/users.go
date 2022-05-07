package querys

const UsersGet = `select * from users`
const UserGet = `select * from users where id = $1`
const UsersInsert = `insert into users (name, username, password, groupid) values (?,?,?,?)`
