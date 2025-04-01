FROM swaggerapi/swagger-ui:v5.20.2

COPY docs/openapi.yaml /docs/openapi.yaml

ENV SWAGGER_JSON=/docs/openapi.yaml