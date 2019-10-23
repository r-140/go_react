# use rabbitmq official
FROM rabbitmq
#https://dzone.com/articles/go-microservices-part-8-centralized-configuration
# enable management plugin
RUN rabbitmq-plugins enable --offline rabbitmq_management

# enable mqtt plugin
RUN rabbitmq-plugins enable --offline rabbitmq_mqtt

# expose management port
EXPOSE 15672
EXPOSE 5672