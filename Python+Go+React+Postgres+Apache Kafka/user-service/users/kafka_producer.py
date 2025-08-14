import json
import logging
from kafka import KafkaProducer
from kafka.errors import KafkaError
from django.conf import settings

logger = logging.getLogger(__name__)

class UserEventProducer:
    def __init__(self):
        self.producer = None
        self.bootstrap_servers = getattr(settings, 'KAFKA_BOOTSTRAP_SERVERS', 'localhost:9092')
        self._connect()
    
    def _connect(self):
        try:
            self.producer = KafkaProducer(
                bootstrap_servers=[self.bootstrap_servers],
                value_serializer=lambda v: json.dumps(v).encode('utf-8'),
                key_serializer=lambda k: str(k).encode('utf-8') if k else None,
                retries=3,
                acks='all'
            )
            logger.info("Kafka producer connected successfully")
        except Exception as e:
            logger.error(f"Failed to connect to Kafka: {e}")
            self.producer = None
    
    def send_user_created_event(self, user_data):
        """Send user created event"""
        if not self.producer:
            logger.warning("Kafka producer not available")
            return False
        
        try:
            event = {
                'event_type': 'user_created',
                'user_id': user_data['id'],
                'user_name': user_data['name'],
                'user_email': user_data['email'],
                'timestamp': user_data['created_at']
            }
            
            future = self.producer.send(
                'user-events', 
                key=str(user_data['id']),
                value=event
            )
            
            # Wait for the message to be sent
            record_metadata = future.get(timeout=10)
            logger.info(f"User created event sent: {record_metadata}")
            return True
            
        except KafkaError as e:
            logger.error(f"Failed to send user created event: {e}")
            return False
    
    def send_user_deleted_event(self, user_id):
        """Send user deleted event"""
        if not self.producer:
            logger.warning("Kafka producer not available")
            return False
        
        try:
            event = {
                'event_type': 'user_deleted',
                'user_id': user_id,
                'timestamp': None  # Add current timestamp in production
            }
            
            future = self.producer.send(
                'user-events',
                key=str(user_id),
                value=event
            )
            
            record_metadata = future.get(timeout=10)
            logger.info(f"User deleted event sent: {record_metadata}")
            return True
            
        except KafkaError as e:
            logger.error(f"Failed to send user deleted event: {e}")
            return False
    
    def close(self):
        """Close the producer"""
        if self.producer:
            self.producer.close()

# Global producer instance
user_event_producer = UserEventProducer()