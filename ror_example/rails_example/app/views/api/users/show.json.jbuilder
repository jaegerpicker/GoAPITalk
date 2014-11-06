json.user do
  json.id    @user.id
  json.username @user.username

  json.todo_id @user.todo ? @user.todo.id : nil
end
