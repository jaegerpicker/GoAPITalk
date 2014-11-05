# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import models, migrations


class Migration(migrations.Migration):

    dependencies = [
        ('api', '0002_auto_20141105_2133'),
    ]

    operations = [
        migrations.RemoveField(
            model_name='todouser',
            name='tasks',
        ),
        migrations.AddField(
            model_name='tasks',
            name='todouser',
            field=models.ForeignKey(default=1, to='api.ToDoUser'),
            preserve_default=False,
        ),
    ]
