json.users @users do |user|
  json.id    user.id
  json.username user.username

  json.todo_id user.todo ? user.todo.id : nil
end
