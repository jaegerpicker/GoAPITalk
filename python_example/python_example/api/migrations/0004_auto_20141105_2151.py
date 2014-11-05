# -*- coding: utf-8 -*-
from __future__ import unicode_literals

from django.db import models, migrations


class Migration(migrations.Migration):

    dependencies = [
        ('api', '0003_auto_20141105_2150'),
    ]

    operations = [
        migrations.AlterField(
            model_name='tasks',
            name='todouser',
            field=models.ForeignKey(related_name='todos', to='api.ToDoUser'),
            preserve_default=True,
        ),
    ]
