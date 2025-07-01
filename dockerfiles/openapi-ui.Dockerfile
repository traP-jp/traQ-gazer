FROM swaggerapi/swagger-ui:v5.25.4

COPY docs/openapi.yaml /docs/openapi.yaml

ENV SWAGGER_JSON=/docs/openapi.yaml