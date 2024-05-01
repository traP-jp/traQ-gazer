FROM swaggerapi/swagger-ui:v5.17.3

COPY docs/openapi.yaml /docs/openapi.yaml

ENV SWAGGER_JSON=/docs/openapi.yaml