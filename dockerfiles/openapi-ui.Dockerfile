FROM swaggerapi/swagger-ui:v5.0.0

COPY docs/openapi.yaml /openapi.yaml

ENV SWAGGER_JSON=/openapi.yaml