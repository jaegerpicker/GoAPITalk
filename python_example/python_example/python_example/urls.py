from django.conf.urls import patterns, include, url
from django.contrib import admin
from django.contrib.auth.models import User
from api.models import ToDoUser, Tasks
from rest_framework import routers, serializers, viewsets

# Serializers define the API representation.
class UserSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = User
        fields = ('username',)

class ToDoUserSerializer(serializers.HyperlinkedModelSerializer):
    todos = serializers.RelatedField(many=True)
    class Meta:
        model = ToDoUser
        fields = ('user', 'todos')

# ViewSets define the view behavior.
class UserViewSet(viewsets.ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer

class ToDoUserViewSet(viewsets.ModelViewSet):
    queryset = ToDoUser.objects.all()
    serializer_class = ToDoUserSerializer

class TasksSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Tasks
        fields = ('todo',)

class TasksViewSet(viewsets.ModelViewSet):
    queryset = Tasks.objects.all()
    serializer_class = TasksSerializer

# Routers provide an easy way of automatically determining the URL conf.
router = routers.DefaultRouter()
router.register(r'users', UserViewSet)
router.register(r'todousers', ToDoUserViewSet)
router.register(r'tasks', TasksViewSet)

urlpatterns = patterns('',
    # Examples:
    # url(r'^$', 'python_example.views.home', name='home'),
    # url(r'^blog/', include('blog.urls')),
    url(r'^', include(router.urls)),
    url(r'^admin/', include(admin.site.urls)),
    url(r'^api-auth/', include('rest_framework.urls', namespace='rest_framework')),
)
