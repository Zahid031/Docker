from rest_framework import viewsets, status
from rest_framework.response import Response
from .models import User
from .serializers import UserSerializer
from .kafka_producer import user_event_producer
import logging

logger = logging.getLogger(__name__)

class UserViewSet(viewsets.ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer
    
    def list(self, request):
        """Get all users"""
        users = User.objects.all()
        serializer = UserSerializer(users, many=True)
        return Response(serializer.data)
    
    def create(self, request):
        """Create a new user and send Kafka event"""
        serializer = UserSerializer(data=request.data)
        if serializer.is_valid():
            user = serializer.save()
            
            # Send Kafka event
            user_data = {
                'id': user.id,
                'name': user.name,
                'email': user.email,
                'created_at': user.created_at.isoformat()
            }
            
            event_sent = user_event_producer.send_user_created_event(user_data)
            if event_sent:
                logger.info(f"User created event sent for user {user.id}")
            else:
                logger.warning(f"Failed to send user created event for user {user.id}")
            
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
    
    def retrieve(self, request, pk=None):
        """Get a specific user"""
        try:
            user = User.objects.get(pk=pk)
            serializer = UserSerializer(user)
            return Response(serializer.data)
        except User.DoesNotExist:
            return Response({'error': 'User not found'}, status=status.HTTP_404_NOT_FOUND)
    
    def update(self, request, pk=None):
        """Update a user"""
        try:
            user = User.objects.get(pk=pk)
            serializer = UserSerializer(user, data=request.data)
            if serializer.is_valid():
                serializer.save()
                return Response(serializer.data)
            return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
        except User.DoesNotExist:
            return Response({'error': 'User not found'}, status=status.HTTP_404_NOT_FOUND)
    
    def destroy(self, request, pk=None):
        """Delete a user and send Kafka event"""
        try:
            user = User.objects.get(pk=pk)
            user_id = user.id
            user.delete()
            
            # Send Kafka event
            event_sent = user_event_producer.send_user_deleted_event(user_id)
            if event_sent:
                logger.info(f"User deleted event sent for user {user_id}")
            else:
                logger.warning(f"Failed to send user deleted event for user {user_id}")
            
            return Response(status=status.HTTP_204_NO_CONTENT)
        except User.DoesNotExist:
            return Response({'error': 'User not found'}, status=status.HTTP_404_NOT_FOUND)