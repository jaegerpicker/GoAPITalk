from django.contrib import admin
from .models import ToDoUser, Tasks
# Register your models here.
admin.site.register(ToDoUser)
admin.site.register(Tasks)
