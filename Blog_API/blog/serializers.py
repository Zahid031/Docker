from rest_framework import serializers
from django.contrib.auth.models import User
from .models import Post
from authentication.serializers import UserSerializer


class PostSerializer(serializers.ModelSerializer):
    author = UserSerializer(read_only=True)
    author_id = serializers.IntegerField(write_only=True, required=False)
    
    class Meta:
        model = Post
        fields = ['id', 'title', 'content', 'author', 'author_id', 'created_at', 'updated_at', 'image']
        read_only_fields = ['id','created_at', 'updated_at']
    
    def create(self, validated_data):
        validated_data['author'] = self.context['request'].user
        if 'author_id' in validated_data:
            validated_data.pop('author_id')
        return super().create(validated_data)
