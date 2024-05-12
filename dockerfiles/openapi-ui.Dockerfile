FROM swaggerapi/swagger-ui:v5.17.7

COPY docs/openapi.yaml /docs/openapi.yaml

ENV SWAGGER_JSON=/docs/openapi.yaml