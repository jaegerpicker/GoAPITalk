from django.db import models
from django.contrib.auth.models import User


# Create your models here.
class ToDoUser(models.Model):
    user = models.OneToOneField(User)

    def __unicode__(self):
        return '%s' % (self.user.username)


class Tasks(models.Model):
    todo = models.CharField(max_length=255)
    todouser = models.ForeignKey(ToDoUser, related_name="todos")

    def __unicode__(self):
        return '%s' % (self.todo)
